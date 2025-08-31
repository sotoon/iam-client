package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	cache "github.com/patrickmn/go-cache"
	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/types"
	"github.com/spf13/viper"
)

// APIURI represents api addr to be appended to server url
const APIURI = "/api/v1/"

const (
	BepaURL    = "https://bepa.sotoon.ir"
	GatewayURL = "https://api.sotoon.ir"
)

type LogLevel int

const (
	DEBUG LogLevel = 0
	INFO  LogLevel = 1
	ERROR LogLevel = 2
)

const HealthyIamURLCachedKey = "healthy_iam_url"
const CacheExpirationDuration = 10 * time.Minute
const CacheCleanupInterval = 10 * time.Minute

type iamClient struct {
	accessToken      string
	baseURL          url.URL
	defaultWorkspace string
	userUUID         string
	logLevel         LogLevel
	apiUrlsList      []*url.URL
	isReliable       bool
	timeout          time.Duration
	cache            Cache
	logger           *log.Logger
}

var _ Client = &iamClient{}

func NewMinimalClient(baseURL string) (Client, error) {
	return NewClient("", baseURL, "", "", DEBUG)
}

// NewClient creates a new client to interact with iam server
func NewClient(accessToken string, baseURL string, defaultWorkspace, userUUID string, logLevel LogLevel) (Client, error) {
	client := &iamClient{}
	client.logLevel = logLevel
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

const DEFAULT_TIMEOUT time.Duration = 2 * time.Second
const MIN_TIMEOUT time.Duration = 1 * time.Second
const MAX_TIMEOUT time.Duration = 5 * time.Second

// returns a reasonable timeout if user has set a bad value
func tuneTimeout(userTimeout time.Duration) time.Duration {
	if userTimeout < MIN_TIMEOUT {
		return MIN_TIMEOUT
	}
	if userTimeout > MAX_TIMEOUT {
		return MAX_TIMEOUT
	}
	return userTimeout
}

// returns a reasonable URL if user has set a bad value
func organizeUrl(userUrl string) string {
	return strings.TrimSpace(userUrl)
}

func (c *iamClient) SetLogger(logger *log.Logger) {
	c.logger = logger
}

func (c *iamClient) initializeServerUrls(serverUrls []string) error {
	if serverUrls == nil || len(serverUrls) == 0 {
		return errors.New("at least one iam server is required")
	}
	for _, serverUrl := range serverUrls {
		serverUrl = organizeUrl(serverUrl)
		fullUrl, err := url.Parse(serverUrl + APIURI)
		if err != nil {
			c.log("URL `%s` is not valid\r\n", fullUrl)
			return err
		}
		c.apiUrlsList = append(c.apiUrlsList, fullUrl)
	}
	return nil
}

// NewReliableClient creates a new reliable client to interact with iam server
// ReliableClient is a client that implements clientside fail-over using a list of iam servers
func NewReliableClient(accessToken string, serverUrls []string, defaultWorkspace, userUUID string, iamTimeout time.Duration) (Client, error) {
	client := &iamClient{}
	client.logLevel = LogLevel(DEBUG)
	client.accessToken = accessToken
	client.defaultWorkspace = defaultWorkspace
	client.userUUID = userUUID
	client.isReliable = true
	client.timeout = tuneTimeout(iamTimeout)
	err := client.initializeServerUrls(serverUrls)
	if err != nil {
		return nil, err
	}
	client.cache = cache.New(CacheExpirationDuration, CacheCleanupInterval)
	return client, nil
}

func NewMinimalReliableClient(serverUrls []string) (Client, error) {
	return NewReliableClient("", serverUrls, "", "", DEFAULT_TIMEOUT)
}

func (c *iamClient) SetAccessToken(token string) {
	c.accessToken = token
}

func (c *iamClient) SetDefaultWorkspace(workspace string) {
	c.defaultWorkspace = workspace
}

func (c *iamClient) SetUser(userUUID string) {
	c.userUUID = userUUID
}

func (c *iamClient) Do(method, path string, successCode int, req interface{}, resp interface{}) error {
	return c.DoWithParams(method, path, nil, successCode, req, resp)
}

func (c *iamClient) DoMinimal(method, path string, resp interface{}) error {
	UsualSuccessCode2xx := 0
	return c.DoWithParams(method, path, nil, UsualSuccessCode2xx, nil, resp)
}

func (c *iamClient) DoSimple(method, path string, parameters map[string]string, req interface{}, resp interface{}) error {
	UsualSuccessCode2xx := 0
	return c.DoWithParams(method, path, parameters, UsualSuccessCode2xx, req, resp)
}

func (c *iamClient) DoWithParams(method, path string, parameters map[string]string, successCode int, req interface{}, resp interface{}) error {

	var body io.Reader
	if req != nil {
		data, err := json.Marshal(req)
		c.log("iam-client marshaling request body:%v", string(data))
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(data)
	}

	httpRequest, err := c.NewRequestWithParameters(method, path, parameters, body)

	if err != nil {
		return err
	}

	// do not log whole request containing authorization secret
	c.log("iam-client performing request method:%v", httpRequest.Method)
	c.log("iam-client performing request url:%v", httpRequest.URL)

	if c.accessToken != "" {
		httpRequest.Header.Add("Content-Type", "application/json")
	}

	data, statusCode, err := processRequest(httpRequest, successCode)

	c.log("iam-client received response code:%d", statusCode)
	c.log("iam-client received response body:%s", data)

	if err == nil {
		if resp != nil {
			return json.Unmarshal(data, resp)
		}
		return nil

	}
	c.log("iam-client faced error:%v", err)
	return &types.RequestExecutionError{
		Err:        err,
		StatusCode: statusCode,
		Data:       data,
	}
}

func processRequest(httpRequest *http.Request, successCode int) ([]byte, int, error) {
	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)

	if err != nil {
		return nil, 0, err
	}

	defer httpResponse.Body.Close()

	err = ensureStatusOK(httpResponse, successCode)
	_, ok := err.(*HTTPResponseError)

	if err == nil || ok {
		data, innerErr := io.ReadAll(httpResponse.Body)
		if innerErr != nil {
			return nil, httpResponse.StatusCode, innerErr
		}
		return data, httpResponse.StatusCode, err
	}

	return nil, httpResponse.StatusCode, err
}

func (c *iamClient) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	return c.NewRequestWithParameters(method, path, nil, body)
}

func (c *iamClient) NewRequestWithParameters(method, path string, parameters map[string]string, body io.Reader) (*http.Request, error) {
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

	serverAddress, err := c.GetBaseURL()
	if err != nil {
		return nil, err
	}
	fullPath := serverAddress.ResolveReference(pathURL)

	req, err := http.NewRequest(method, fullPath.String(), body)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	return req, nil

}

func getHealthCheckValue(c *iamClient, serverUrl *url.URL, resultChannel chan *url.URL) error {
	err := healthCheck(c, serverUrl)
	resp := types.HealthCheckResponse{ServerUrl: serverUrl.String(), Err: err}
	if err != nil {
		c.log("healthCheck failed. error: %v\n", err)
		return err
	} else {
		select {
		case resultChannel <- serverUrl:
			c.log("healthCheck successful: %v\n", resp)
			return nil
		case <-time.After(c.timeout):
			c.log("healthCheck not used: %v", resp)
			return nil
		}
	}
}

func (c *iamClient) GetHealthyIamURL() (*url.URL, error) {
	/*
		A Note about the channelSize := 0
		The channel size for a healthy server is set to zero (0) because we want to block the execution until the first response is received.
		Since we don't require buffering in this case, the channel doesn't need a buffer.
		Additionally, the Go garbage collector (GC) will automatically remove the channel when it becomes unreachable.
		Therefore, there is no need to explicitly close the serverChannel, which could potentially cause a panic if other goroutines attempt to write to it.
	*/
	channelSize := 0
	serverUrlChannel := make(chan *url.URL, channelSize)
	for _, serverUrl := range c.apiUrlsList {
		newServerUrl := serverUrl
		go getHealthCheckValue(c, newServerUrl, serverUrlChannel)
	}

	select {
	case res := <-serverUrlChannel:
		return res, nil
	case <-time.After(c.timeout):
		return nil, errors.New("no available iam servers found")
	}
}

func (c *iamClient) GetBaseURL() (*url.URL, error) {
	if !c.isReliable {
		return &c.baseURL, nil
	}
	if cached, found := c.cache.Get(HealthyIamURLCachedKey); found {
		iamURL := cached.(*url.URL)
		return iamURL, nil
	}
	iamURL, err := c.GetHealthyIamURL()
	if err == nil {
		c.cache.Set(HealthyIamURLCachedKey, iamURL, cache.DefaultExpiration)
	}
	return iamURL, err
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

func (c *iamClient) GetServerURL() string {
	url := c.baseURL.String()
	return strings.Replace(url, APIURI, "", -1)
}

func (c *iamClient) SetConfigDefaultUserData(context, token, userUUID, email string) error {
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

func (c *iamClient) SetCurrentContext(context string) error {
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

func (c *iamClient) IsHealthy() (bool, error) {
	serverUrl, err := c.GetBaseURL()
	if c.isReliable {
		// there is no need to check health of iam in reliable client, the GetIamUrl has already did it
		if err == nil {
			return true, nil
		}
		return false, err
	}
	// we should check the health of iam endpoint, in simple client (when c.isReliable is false)
	err = healthCheck(c, serverUrl)
	return err == nil, err
}

func (c *iamClient) SetConfigDefaultWorkspace(uuid *uuid.UUID) error {
	context := viper.GetString("current-context")
	viper.Set(fmt.Sprintf("contexts.%s.workspace", context), uuid.String())
	c.defaultWorkspace = uuid.String()
	return persistClientConfigFile()
}

func getDefaultLogger() *log.Logger {
	return log.Default()
}

func (c *iamClient) log(messageFmt string, objects ...interface{}) {
	if c.logger == nil {
		c.logger = getDefaultLogger()
	}

	if c.logLevel <= LogLevel(DEBUG) {
		c.logger.Printf(messageFmt, objects...)
	}
}
