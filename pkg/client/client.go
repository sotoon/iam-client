package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
)

// APIURI represents api addr to be appended to server url
const APIURI = "/api/v1/"

var errRetriesExceeded = errors.New("maximum retries exceeded on targets")

type targets []*target

func (s targets) Len() int {
	return len(s)
}
func (s targets) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s targets) Less(i, j int) bool {
	return (s[i].penaltyTime.Unix() > s[j].penaltyTime.Unix())
}

type loadBalancer struct {
	Targets targets
}

type target struct {
	BaseURL     *url.URL
	penaltyTime time.Time
}

func (t *target) setPenalty() {
	t.penaltyTime = time.Unix(0, 0)
}

func (t *target) setReward() {
	t.penaltyTime = time.Now()
}

func (lb *loadBalancer) getTarget() *target {
	sort.Sort(lb.Targets)
	var noiseRation float32 = 0.1
	if targetsCount := len(lb.Targets); targetsCount > 1 && rand.Float32() < noiseRation {
		return lb.Targets[rand.Intn(targetsCount-1)+1]
	}

	return lb.Targets[0]
}

func (lb *loadBalancer) setTargets(urls []*url.URL) {
	targets := make([]target, len(urls))
	for i, url := range urls {
		targets[i] = target{
			BaseURL:     url,
			penaltyTime: time.Now(),
		}
	}
}

func (lb *loadBalancer) addTarget(url *url.URL) {
	lb.Targets = append(lb.Targets, &target{
		BaseURL:     url,
		penaltyTime: time.Now(),
	})
}

func (lb *loadBalancer) TargetsLen() int {
	return len(lb.Targets)
}

type bepaClient struct {
	accessToken      string
	baseURL          url.URL
	defaultWorkspace string
	userUUID         string
	loadBalancer     *loadBalancer
}

var _ Client = &bepaClient{}

// NewClient creates a new client to interact with bepa server
func NewClient(accessToken string, targets []string, defaultWorkspace, userUUID string) (Client, error) {
	client := &bepaClient{}
	lb := &loadBalancer{Targets: make([]*target, 0)}
	for _, urlStr := range targets {
		url, err := createServerURL(urlStr)
		if err != nil {
			return nil, err
		}

		lb.addTarget(url)

	}
	client.accessToken = accessToken
	client.defaultWorkspace = defaultWorkspace
	client.userUUID = userUUID
	client.loadBalancer = lb
	return client, nil
}

func (c *bepaClient) SetAccessToken(token string) {
	c.accessToken = token
}

func (c *bepaClient) SetUser(userUUID string) {
	c.userUUID = userUUID
}
func checkErrorsAndPenaltyReward(err error, target *target) (bool, error) {
	httpErr, ok := err.(*HTTPResponseError)
	if ok && !httpErr.IsFaulty || err == nil {
		target.setReward()
		return err != nil, err
	}
	target.setPenalty()
	return false, err
}
func (c *bepaClient) Do(method, path string, req interface{}, resp interface{}) error {
	var body io.Reader
	if req != nil {
		data, err := json.Marshal(req)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(data)
	}

	for i := 0; i < c.loadBalancer.TargetsLen()*2; i++ {
		httpRequest, target, err := c.NewRequest(method, path, body)

		if err != nil {
			return err
		}

		if c.accessToken != "" {
			httpRequest.Header.Add("Content-Type", "application/json")
		}

		data, err := proccessRequest(httpRequest, target)

		if loopBreaker, err := checkErrorsAndPenaltyReward(err, target); err != nil && loopBreaker {
			return err
		}

		if err == nil && resp != nil {
			return json.Unmarshal(data, resp)
		}

	}
	return errRetriesExceeded
}

func proccessRequest(httpRequest *http.Request, target *target) ([]byte, error) {
	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)

	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()
	if err := ensureStatusOK(httpResponse); err == nil {
		data, innerErr := ioutil.ReadAll(httpResponse.Body)
		if innerErr != nil {
			return nil, innerErr
		}

		return data, nil
	}
	return nil, err
}

func (c *bepaClient) NewRequest(method, path string, body io.Reader) (*http.Request, *target, error) {
	pathURL, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	target := c.loadBalancer.getTarget()
	fullPath := target.BaseURL.ResolveReference(pathURL)

	req, err := http.NewRequest(method, fullPath.String(), body)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	return req, target, nil

}

func (c *bepaClient) Authorize(identity, action, object string) error {
	for i := 0; i < c.loadBalancer.TargetsLen()*2; i++ {
		req, target, err := c.NewRequest(http.MethodGet, trimURLSlash(routes.RouteAuthz), nil)

		if err != nil {
			return err
		}

		query := req.URL.Query()
		query.Set("identity", identity)
		query.Set("object", object)
		query.Set("action", action)

		req.URL.RawQuery = query.Encode()
		_, err = proccessRequest(req, target)

		if loopBreaker, err := checkErrorsAndPenaltyReward(err, target); err != nil && loopBreaker {
			return err
		}
		return nil
	}
	return errRetriesExceeded
}

func (c *bepaClient) Identify(token string) (*types.UserRes, error) {
	idenReq := &types.UserTokenReq{
		Secret: token,
	}

	userRes := &types.UserRes{}
	err := c.Do(http.MethodPost, trimURLSlash(routes.RouteUserTokenIdentify), idenReq, userRes)
	if err != nil {
		return nil, err
	}

	return userRes, nil
}

func createServerURL(serverURL string) (*url.URL, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}

	apiURL, err := url.Parse(APIURI)

	if err != nil {
		return nil, err
	}
	u = u.ResolveReference(apiURL)
	return u, nil
}

func (c *bepaClient) GetServerURL() string {
	url := c.baseURL.String()
	return strings.Replace(url, APIURI, "", -1)
}
