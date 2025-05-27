package client

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker/support/slice"
	"github.com/golang/mock/gomock"
	"github.com/sotoon/iam-client/mocks"
	"github.com/stretchr/testify/require"
)

func TestHealthCheckAPI(t *testing.T) {
	// todo: add more fields to test cases when health-check fields are updated
	testCases := []struct {
		token  string
		status int
	}{
		{"sampleusertoken", http.StatusOK},
		{"sampleusertoken", http.StatusInternalServerError},
		{"sampleusertoken", http.StatusGatewayTimeout},
		{"sampleusertoken", http.StatusBadGateway},
	}

	for _, tc := range testCases {
		s := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				require.True(t, strings.HasPrefix(r.URL.Path, "/api/v1/healthz"))

				if tc.status != http.StatusOK {
					w.WriteHeader(tc.status)
					return
				}
				w.WriteHeader(http.StatusOK)
			}))
		serverUrl, err := url.Parse(s.URL)
		pathURL, err := url.Parse("/api/v1/healthz/")
		fullPath := serverUrl.ResolveReference(pathURL)
		client := &iamClient{
			accessToken: tc.token,
			timeout:     1 * time.Second,
		}
		err = healthCheck(client, fullPath)

		if tc.status == http.StatusOK {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
		s.Close()
	}

}

type IamHandlerFunc func(w http.ResponseWriter, r *http.Request)

type IamServer struct {
	handler   IamHandlerFunc
	isHealthy bool
}

func timeoutIamHandler(w http.ResponseWriter, r *http.Request) {
	// todo: investigate good practice for removing constant time.
	timeoutDuration := MAX_TIMEOUT + 1
	time.Sleep(timeoutDuration)
	w.WriteHeader(http.StatusOK)
}

func unhealthyIamHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func rateLimitingIamHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTooManyRequests)
}

func healthyIamHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var timeoutIamServer = IamServer{timeoutIamHandler, false}
var unhealthyIamServer = IamServer{unhealthyIamHandler, false}
var healthyIamServer = IamServer{healthyIamHandler, true}
var rateLimitingIamServer = IamServer{rateLimitingIamHandler, false}

func TestGetHealthyIamURL(t *testing.T) {
	testCases := []struct {
		servers      []IamServer
		serverUrls   []string
		correctUrls  []string
		iamAvailable bool
	}{
		{servers: []IamServer{timeoutIamServer, timeoutIamServer, timeoutIamServer}, iamAvailable: false},
		{servers: []IamServer{timeoutIamServer, unhealthyIamServer, timeoutIamServer}, iamAvailable: false},
		{servers: []IamServer{unhealthyIamServer, unhealthyIamServer, unhealthyIamServer}, iamAvailable: false},
		{servers: []IamServer{healthyIamServer, unhealthyIamServer, unhealthyIamServer}, iamAvailable: true},
		{servers: []IamServer{healthyIamServer, unhealthyIamServer, timeoutIamServer}, iamAvailable: true},
		{servers: []IamServer{healthyIamServer, healthyIamServer, healthyIamServer}, iamAvailable: true},
		{servers: []IamServer{rateLimitingIamServer, rateLimitingIamServer, rateLimitingIamServer}, iamAvailable: false},
		{servers: []IamServer{rateLimitingIamServer, timeoutIamServer, unhealthyIamServer}, iamAvailable: false},
	}
	for _, tc := range testCases {
		for _, testIamServer := range tc.servers {
			s := httptest.NewServer(http.HandlerFunc(testIamServer.handler))
			tc.serverUrls = append(tc.serverUrls, s.URL)
			if testIamServer.isHealthy {
				fullURL, err := url.Parse(s.URL + APIURI)
				require.NoError(t, err)
				tc.correctUrls = append(tc.correctUrls, fullURL.String())
			}
			defer s.Close()
		}
		c := NewTestReliableClient(tc.serverUrls, nil)
		serverUrl, err := c.GetBaseURL()
		isHealthy, healthError := c.IsHealthy()
		if tc.iamAvailable {
			require.True(t, isHealthy)
			require.NoError(t, healthError)
			require.NoError(t, err)
			require.True(t, slice.Contains(tc.correctUrls, serverUrl.String()))
		} else {
			require.False(t, isHealthy)
			require.Error(t, healthError)
			require.Error(t, err)
			require.Nil(t, serverUrl)
		}

	}
	// to see the request logs
	time.Sleep(MAX_TIMEOUT)
}

func TestSetCacheOnHealthyIamURL(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var server = healthyIamServer

	s := httptest.NewServer(http.HandlerFunc(server.handler))
	serverUrl := s.URL
	fullUrl, err := url.Parse(serverUrl + APIURI)
	require.NoError(t, err)

	cache := mocks.NewMockCache(mockCtrl)
	cache.EXPECT().
		Get(HealthyIamURLCachedKey).
		Return(nil, false).
		Times(1)
	cache.EXPECT().
		Set(HealthyIamURLCachedKey, fullUrl, gomock.Any()).
		Times(1)
	cache.EXPECT().
		Get(HealthyIamURLCachedKey).
		Return(fullUrl, true).
		Times(1)

	c := NewTestReliableClient([]string{serverUrl}, cache)

	healthyUrl, err := c.GetBaseURL()
	require.NoError(t, err)
	require.Equal(t, fullUrl, healthyUrl)

	cachedUrl, err := c.GetBaseURL()
	require.NoError(t, err)
	require.Equal(t, fullUrl, cachedUrl)
}

func TestDontSetCacheOnUnhealthyIamURL(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var server = unhealthyIamServer

	s := httptest.NewServer(http.HandlerFunc(server.handler))
	serverUrl := s.URL

	cache := mocks.NewMockCache(mockCtrl)
	cache.EXPECT().
		Get(HealthyIamURLCachedKey).
		Return(nil, false).
		Times(1)
	cache.EXPECT().
		Set(gomock.Any(), gomock.Any(), gomock.Any()).
		MaxTimes(0)
	cache.EXPECT().
		Get(HealthyIamURLCachedKey).
		Return(nil, false).
		Times(1)
	cache.EXPECT().
		Set(gomock.Any(), gomock.Any(), gomock.Any()).
		MaxTimes(0)

	c := NewTestReliableClient([]string{serverUrl}, cache)

	_, err := c.GetBaseURL()
	require.Error(t, err)

	_, err = c.GetBaseURL()
	require.Error(t, err)
}
