package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"

	"github.com/stretchr/testify/require"
)

func TestIdentification(t *testing.T) {
	testCases := []struct {
		uuid  string
		token string
		found bool
	}{
		{"user1uuid", "sampleusertoken", true},
		{"user5uuid", "sampleusertoken", false},
	}

	for _, tc := range testCases {
		s := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				require.True(t, strings.HasPrefix(r.URL.Path, "/api/v1/identify"))

				var idenReq types.UserTokenReq
				require.NoError(t, json.NewDecoder(r.Body).Decode(&idenReq))

				require.Equal(t, tc.token, idenReq.Secret)

				if !tc.found {
					w.WriteHeader(http.StatusNotFound)
					respErr := types.ResponseError{
						Error: http.StatusText(http.StatusNotFound),
					}
					require.NoError(t, json.NewEncoder(w).Encode(respErr))
					return
				}
				w.WriteHeader(http.StatusOK)
				userRes := types.UserRes{
					UUID: tc.uuid,
				}
				require.NoError(t, json.NewEncoder(w).Encode(userRes))
			}))

		c := NewTestClient(s)
		user, err := c.Identify(tc.token)

		if tc.found {
			require.NoError(t, err)
			require.Equal(t, tc.uuid, user.UUID)
		} else {
			require.Error(t, err)
		}

		s.Close()
	}

}

var benchmarkIterNumber int = 100
var bepaEndpoint string = "https://bepa.sotoon.ir"
var bepaBenchmarkToken string = os.Getenv("BENCHMARK_TOKEN")
var timeoutDuration time.Duration = 10 * time.Second

func DoSingleBenchmarkIdentify(token string, wg *sync.WaitGroup) {
	serverList := []string{bepaEndpoint, bepaEndpoint, bepaEndpoint}
	c, _ := NewReliableClient(bepaBenchmarkToken, serverList, "", "", timeoutDuration)
	c.Identify(token)
	wg.Done()
}

func BenchmarkMultipleValidIdentify(b *testing.B) {
	var wg sync.WaitGroup
	b.Run(fmt.Sprintf("concurrent_iters_%d", benchmarkIterNumber), func(b *testing.B) {
		var iters int = benchmarkIterNumber
		for j := 0; j < b.N; j++ {
			wg.Add(iters)
			for i := 0; i < iters; i++ {
				go DoSingleBenchmarkIdentify(bepaBenchmarkToken, &wg)
			}
			wg.Wait()
		}
	})
}

func BenchmarkMultipleInvalidIdentify(b *testing.B) {
	var wg sync.WaitGroup
	b.Run(fmt.Sprintf("concurrent_iters_%d", benchmarkIterNumber), func(b *testing.B) {
		var iters int = benchmarkIterNumber
		for j := 0; j < b.N; j++ {
			wg.Add(iters)
			for i := 0; i < iters; i++ {
				go DoSingleBenchmarkIdentify(randomString(10), &wg)
			}
			wg.Wait()
		}
	})
}
