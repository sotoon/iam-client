package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

// APIURI represents api addr to be appended to server url
const APIURI = "/api/v1/"

type LogLevel int

const (
	DEBUG int = 0
	INFO  int = 1
	ERROR int = 2
)

type bepaClient struct {
	accessToken      string
	baseURL          url.URL
	defaultWorkspace string
	userUUID         string
	logLevel         LogLevel
	apiUrlsList      []*url.URL
	isReliable       bool
	bepaTimeout      time.Duration
}

var _ Client = &bepaClient{}

func NewMinimalClient(baseURL string) (Client, error) {
	return NewClient("", baseURL, "", "")
}

// NewClient creates a new client to interact with bepa server
func NewClient(accessToken string, baseURL string, defaultWorkspace, userUUID string) (Client, error) {
	client := &bepaClient{}
	client.logLevel = LogLevel(DEBUG)
	client.accessToken = accessToken
	client.defaultWorkspace = defaultWorkspace
	client.userUUID = userUUID
	url, err := url.Parse(baseURL + APIURI)
	if err != nil {
		fmt.Printf("Base URL `%s` is not valid\r\n", baseURL)
		panic(err)
	}
	client.baseURL = *url
	return client, nil
}

const MIN_TIMEOUT_SECUNDS time.Duration = 1 * time.Second
const MAX_TIMEOUT_SECUNDS time.Duration = 3 * time.Second

// returns a reasonable timeout if user has set a bad value
func tuneTimeout(userTimeout time.Duration) time.Duration {
	if MIN_TIMEOUT_SECUNDS < userTimeout && userTimeout < MAX_TIMEOUT_SECUNDS {
		return userTimeout
	}
	return MIN_TIMEOUT_SECUNDS
}

// returns a reasonable URL if user has set a bad value
func organizeUrl(userUrl string) string {
	return strings.TrimSpace(userUrl)
}

// NewReliableClient creates a new reliable client to interact with bepa server
// ReliableClient is a client that implements clientside fail-over using a list of bepa servers
func NewReliableClient(accessToken string, serverUrlsList []string, defaultWorkspace, userUUID string, bepaTimeout time.Duration) (Client, error) {
	client := &bepaClient{}
	client.logLevel = LogLevel(DEBUG)
	client.accessToken = accessToken
	client.defaultWorkspace = defaultWorkspace
	client.userUUID = userUUID
	client.isReliable = true
	client.bepaTimeout = tuneTimeout(bepaTimeout)
	for _, serverUrl := range serverUrlsList {
		serverUrl = organizeUrl(serverUrl)
		fullUrl, err := url.Parse(serverUrl + APIURI)
		if err != nil {
			client.log("URL `%s` is not valid\r\n", fullUrl)
			return nil, err
		}
		client.apiUrlsList = append(client.apiUrlsList, fullUrl)
	}
	return client, nil
}

func NewMinimalReliableClient(serverUrlsList []string) (Client, error) {
	return NewReliableClient("", serverUrlsList, "", "", MAX_TIMEOUT_SECUNDS)
}

func (c *bepaClient) SetAccessToken(token string) {
	c.accessToken = token
}

func (c *bepaClient) SetDefaultWorkspace(workspace string) {
	c.defaultWorkspace = workspace
}

func (c *bepaClient) SetUser(userUUID string) {
	c.userUUID = userUUID
}

func (c *bepaClient) Do(method, path string, successCode int, req interface{}, resp interface{}) error {
	return c.DoWithParams(method, path, nil, successCode, req, resp)
}

func (c *bepaClient) DoMinimal(method, path string, resp interface{}) error {
	USUAL_SUCCESS_CODE_2XX := 0
	return c.DoWithParams(method, path, nil, USUAL_SUCCESS_CODE_2XX, nil, resp)
}

func (c *bepaClient) DoSimple(method, path string, parameters map[string]string, req interface{}, resp interface{}) error {
	USUAL_SUCCESS_CODE_2XX := 0
	return c.DoWithParams(method, path, parameters, USUAL_SUCCESS_CODE_2XX, req, resp)
}

func (c *bepaClient) DoWithParams(method, path string, parameters map[string]string, successCode int, req interface{}, resp interface{}) error {

	var body io.Reader
	if req != nil {
		data, err := json.Marshal(req)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(data)
	}

	httpRequest, err := c.NewRequestWithParameters(method, path, parameters, body)
	c.log("bepa-client performing request:%v", httpRequest)

	if err != nil {
		return err
	}

	if c.accessToken != "" {
		httpRequest.Header.Add("Content-Type", "application/json")
	}

	data, statusCode, err := proccessRequest(httpRequest, successCode)

	c.log("bepa-client received response code:%d", statusCode)
	c.log("bepa-client received response body:%v", data)
	c.log("bepa-client faced error:%v", err)

	if err == nil {
		if resp != nil {
			return json.Unmarshal(data, resp)
		}
		return nil

	}

	return &types.RequestExecutionError{
		Err:        err,
		StatusCode: statusCode,
		Data:       data,
	}
}

func proccessRequest(httpRequest *http.Request, successCode int) ([]byte, int, error) {
	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)

	if err != nil {
		return nil, 0, err
	}

	defer httpResponse.Body.Close()

	err = ensureStatusOK(httpResponse, successCode)
	_, ok := err.(*HTTPResponseError)

	if err == nil || ok {
		data, innerErr := ioutil.ReadAll(httpResponse.Body)
		if innerErr != nil {
			return nil, httpResponse.StatusCode, innerErr
		}
		return data, httpResponse.StatusCode, err
	}

	return nil, httpResponse.StatusCode, err
}

func (c *bepaClient) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	return c.NewRequestWithParameters(method, path, nil, body)
}

func (c *bepaClient) NewRequestWithParameters(method, path string, parameters map[string]string, body io.Reader) (*http.Request, error) {
	pathURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	if parameters != nil {
		params := url.Values{}
		for key, val := range parameters {
			params.Add(key, val)
		}
		pathURL.RawQuery = params.Encode()
	}

	serverAddress, err := c.GetBepaURL()
	if err != nil {
		return nil, err
	}
	fullPath := serverAddress.ResolveReference(pathURL)

	req, err := http.NewRequest(method, fullPath.String(), body)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	return req, nil

}

func getHealthCheckValue(c *bepaClient, serverUrl *url.URL, resultChannel chan *url.URL) error {
	err := healthCheck(c, serverUrl)
	resp := types.HealthCheckResponse{serverUrl.String(), err}
	if err != nil {
		c.log("healthCheck failed. error: %v\n", err)
		return err
	} else {
		c.log("healthCheck successful. %v\n", resp)
		resultChannel <- serverUrl
		return nil
	}
}

func (c *bepaClient) GetBepaURL() (*url.URL, error) {
	if !c.isReliable {
		return &c.baseURL, nil
	}
	//todo: add cache for healthy url
	//todo: stop go routines after first healthcheck ack arrives
	serverUrlChannel := make(chan *url.URL, 1)
	for _, serverUrl := range c.apiUrlsList {
		newServerUrl := serverUrl
		go getHealthCheckValue(c, newServerUrl, serverUrlChannel)
	}

	select {
	case res := <-serverUrlChannel:
		return res, nil
	case <-time.After(c.bepaTimeout):
		return nil, errors.New("no available BEPA servers found")
	}
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

func (c *bepaClient) SetConfigDefaultUserData(context, token, userUUID, email string) error {
	if context == "" {
		context = "default"
	}
	viper.Set(fmt.Sprintf("contexts.%s.token", context), token)
	viper.Set(fmt.Sprintf("contexts.%s.user-uuid", context), userUUID)
	viper.Set(fmt.Sprintf("contexts.%s.user", context), email)
	viper.Set(fmt.Sprintf("contexts.%s.addr", context), c.GetServerURL())
	c.accessToken = token
	c.userUUID = userUUID
	return persistClientConfigFile()
}

func (c *bepaClient) SetCurrentContext(context string) error {
	contexts := viper.GetStringMap("contexts")
	if _, ok := contexts[context]; ok {
		viper.Set("current-context", context)
		if err := persistClientConfigFile(); err == nil {
			fmt.Printf("set default context to %s\n", context)
			return nil
		}
	}
	return fmt.Errorf("could not find context %s", context)
}

func (c *bepaClient) SetConfigDefaultWorkspace(uuid *uuid.UUID) error {
	context := viper.GetString("current-context")
	viper.Set(fmt.Sprintf("contexts.%s.workspace", context), uuid.String())
	c.defaultWorkspace = uuid.String()
	return persistClientConfigFile()
}

func (c *bepaClient) log(messageFmt string, object interface{}) {
	if c.logLevel <= LogLevel(DEBUG) {
		log.Printf(messageFmt, object)
	}
}
