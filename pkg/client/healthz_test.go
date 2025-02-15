package client

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"git.platform.sotoon.ir/iam/golang-bepa-client/mocks"
	"github.com/bxcodec/faker/support/slice"
	"github.com/golang/mock/gomock"
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
		client := &bepaClient{
			accessToken: tc.token,
			bepaTimeout: 1 * time.Second,
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

type BepaHandlerFunc func(w http.ResponseWriter, r *http.Request)

type BepaServer struct {
	handler   BepaHandlerFunc
	isHealthy bool
}

func timeoutBepaHandler(w http.ResponseWriter, r *http.Request) {
	// todo: investigate good practice for removing constant time.
	timeoutDuration := MAX_TIMEOUT + 1
	time.Sleep(timeoutDuration)
	w.WriteHeader(http.StatusOK)
}

func unhealthyBepaHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func rateLimitingBepaHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTooManyRequests)
}

func healthyBepaHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var timeoutBepaServer = BepaServer{timeoutBepaHandler, false}
var unhealthyBepaServer = BepaServer{unhealthyBepaHandler, false}
var healthyBepaServer = BepaServer{healthyBepaHandler, true}
var rateLimitingBepaServer = BepaServer{rateLimitingBepaHandler, false}

func TestGetHealthyBepaURL(t *testing.T) {
	testCases := []struct {
		servers       []BepaServer
		serverUrls    []string
		correctUrls   []string
		bepaAvailable bool
	}{
		{servers: []BepaServer{timeoutBepaServer, timeoutBepaServer, timeoutBepaServer}, bepaAvailable: false},
		{servers: []BepaServer{timeoutBepaServer, unhealthyBepaServer, timeoutBepaServer}, bepaAvailable: false},
		{servers: []BepaServer{unhealthyBepaServer, unhealthyBepaServer, unhealthyBepaServer}, bepaAvailable: false},
		{servers: []BepaServer{healthyBepaServer, unhealthyBepaServer, unhealthyBepaServer}, bepaAvailable: true},
		{servers: []BepaServer{healthyBepaServer, unhealthyBepaServer, timeoutBepaServer}, bepaAvailable: true},
		{servers: []BepaServer{healthyBepaServer, healthyBepaServer, healthyBepaServer}, bepaAvailable: true},
		{servers: []BepaServer{rateLimitingBepaServer, rateLimitingBepaServer, rateLimitingBepaServer}, bepaAvailable: false},
		{servers: []BepaServer{rateLimitingBepaServer, timeoutBepaServer, unhealthyBepaServer}, bepaAvailable: false},
	}
	for _, tc := range testCases {
		for _, testBepaServer := range tc.servers {
			s := httptest.NewServer(http.HandlerFunc(testBepaServer.handler))
			tc.serverUrls = append(tc.serverUrls, s.URL)
			if testBepaServer.isHealthy {
				fullURL, err := url.Parse(s.URL + APIURI)
				require.NoError(t, err)
				tc.correctUrls = append(tc.correctUrls, fullURL.String())
			}
			defer s.Close()
		}
		c := NewTestReliableClient(tc.serverUrls, nil)
		serverUrl, err := c.GetBepaURL()
		isHealthy, healthError := c.IsHealthy()
		if tc.bepaAvailable {
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

func TestSetCacheOnhealthyBepaURL(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var server = healthyBepaServer

	s := httptest.NewServer(http.HandlerFunc(server.handler))
	serverUrl := s.URL
	fullUrl, err := url.Parse(serverUrl + APIURI)
	require.NoError(t, err)

	cache := mocks.NewMockCache(mockCtrl)
	cache.EXPECT().
		Get(HealthyBepaURLCachedKey).
		Return(nil, false).
		Times(1)
	cache.EXPECT().
		Set(HealthyBepaURLCachedKey, fullUrl, gomock.Any()).
		Times(1)
	cache.EXPECT().
		Get(HealthyBepaURLCachedKey).
		Return(fullUrl, true).
		Times(1)

	c := NewTestReliableClient([]string{serverUrl}, cache)

	healthyUrl, err := c.GetBepaURL()
	require.NoError(t, err)
	require.Equal(t, fullUrl, healthyUrl)

	cachedUrl, err := c.GetBepaURL()
	require.NoError(t, err)
	require.Equal(t, fullUrl, cachedUrl)
}

func TestDontSetCacheOnUnhealthyBepaURL(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var server = unhealthyBepaServer

	s := httptest.NewServer(http.HandlerFunc(server.handler))
	serverUrl := s.URL

	cache := mocks.NewMockCache(mockCtrl)
	cache.EXPECT().
		Get(HealthyBepaURLCachedKey).
		Return(nil, false).
		Times(1)
	cache.EXPECT().
		Set(gomock.Any(), gomock.Any(), gomock.Any()).
		MaxTimes(0)
	cache.EXPECT().
		Get(HealthyBepaURLCachedKey).
		Return(nil, false).
		Times(1)
	cache.EXPECT().
		Set(gomock.Any(), gomock.Any(), gomock.Any()).
		MaxTimes(0)

	c := NewTestReliableClient([]string{serverUrl}, cache)

	_, err := c.GetBepaURL()
	require.Error(t, err)

	_, err = c.GetBepaURL()
	require.Error(t, err)
}
