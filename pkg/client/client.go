package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	"github.com/pkg/errors"
)

// APIURI represents api addr to be appended to server url
const APIURI = "/api/v1/"

type bepaClient struct {
	accessToken      string
	baseURL          url.URL
	defaultWorkspace string
	userUUID         string
}

var _ Client = &bepaClient{}

// NewClient creates a new client to interact with bepa server
func NewClient(accessToken, baseURL, defaultWorkspace, userUUID string) (*bepaClient, error) {
	client := &bepaClient{}
	if err := client.SetServerURL(baseURL); err != nil {
		return nil, err
	}
	client.accessToken = accessToken
	client.defaultWorkspace = defaultWorkspace
	client.userUUID = userUUID
	return client, nil
}

func (c *bepaClient) SetAccessToken(token string) {
	c.accessToken = token
}

func (c *bepaClient) SetUser(userUUID string) {
	c.userUUID = userUUID
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

	httpRequest, err := c.NewRequest(method, path, body)
	if err != nil {
		return err
	}

	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return errors.Wrap(err, "error while requesting")
	}
	defer httpResponse.Body.Close()

	if err := ensureStatusOK(httpResponse); err != nil {
		return err
	}

	if resp != nil {
		data, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(data, resp)
	}

	return nil
}

func (c *bepaClient) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	pathURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	fullPath := c.baseURL.ResolveReference(pathURL)

	req, err := http.NewRequest(method, fullPath.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	return req, nil
}

func (c *bepaClient) Authorize(identity, action, object string) error {
	req, err := c.NewRequest(http.MethodGet, trimURLSlash(routes.RouteAuthz), nil)
	if err != nil {
		return err
	}

	query := req.URL.Query()
	query.Set("identity", identity)
	query.Set("object", object)
	query.Set("action", action)
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return ensureStatusOK(resp)
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

func (c *bepaClient) SetServerURL(serverURL string) error {
	u, err := url.Parse(serverURL)
	if err != nil {
		return err
	}

	apiURL, err := url.Parse(APIURI)
	if err != nil {
		return err
	}
	u = u.ResolveReference(apiURL)
	c.baseURL = *u
	return nil
}

func (c *bepaClient) GetServerURL() string {
	url := c.baseURL.String()
	return strings.Replace(url, APIURI, "", -1)
}
