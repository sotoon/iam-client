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

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
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
func NewClient(accessToken string, baseURL string, defaultWorkspace, userUUID string) (Client, error) {
	client := &bepaClient{}

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

	if c.accessToken != "" {
		httpRequest.Header.Add("Content-Type", "application/json")
	}

	data, err := proccessRequest(httpRequest)

	if err == nil {
		if resp != nil {
			return json.Unmarshal(data, resp)
		}
		return nil

	}
	return err
}

func proccessRequest(httpRequest *http.Request) ([]byte, error) {
	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)

	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	err = ensureStatusOK(httpResponse)
	if err == nil {
		data, innerErr := ioutil.ReadAll(httpResponse.Body)
		if innerErr != nil {
			return nil, innerErr
		}
		return data, nil
	}

	return nil, err
}

func (c *bepaClient) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	pathURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	fullPath := c.baseURL.ResolveReference(pathURL)

	req, err := http.NewRequest(method, fullPath.String(), body)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	return req, nil

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
