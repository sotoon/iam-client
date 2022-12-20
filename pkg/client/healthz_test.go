package client

import (
	"github.com/bxcodec/faker/support/slice"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
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
	timeoutDuration := 100 * time.Millisecond
	time.Sleep(timeoutDuration)
	w.WriteHeader(http.StatusOK)
}

func unhealthyBepaHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func healthyBepaHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var timeoutBepaServer = BepaServer{timeoutBepaHandler, false}
var unhealthyBepaServer = BepaServer{unhealthyBepaHandler, false}
var healthyBepaServer = BepaServer{healthyBepaHandler, true}

func TestGetBepaURL(t *testing.T) {
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
		c := NewTestReliableClient(tc.serverUrls)
		serverUrl, err := c.GetBepaURL()
		if tc.bepaAvailable {
			require.NoError(t, err)
			require.True(t, slice.Contains(tc.correctUrls, serverUrl.String()))
		} else {
			require.Error(t, err)
			require.Nil(t, serverUrl)
		}

	}

}
