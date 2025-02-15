package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"git.platform.sotoon.ir/iam/golang-bepa-client/pkg/types"

	"github.com/stretchr/testify/require"
)

type rule struct {
	uuid     string
	userType string
	action   string
	rri      string
}

var rriRegex = regexp.MustCompile(`^rri:v1:cafebazaar\.cloud:(.+):(.+):(\/.+)*\/?$`)

func (r1 rule) equal(r2 rule) bool {
	return r1.uuid == r2.uuid &&
		r1.rri == r2.rri &&
		r1.action == r2.action
}

func (r rule) isValid() bool {
	return r.uuid != "" &&
		r.action != "" &&
		r.userType != "" &&
		rriRegex.MatchString(r.rri)
}

func TestAuthorization(t *testing.T) {

	testCases := []struct {
		rule
		valid bool
	}{
		{rule{"user1uuid", "user", "get", newRRI("workspace", "ns", "pod")}, true},
		{rule{"user2uuid", "service-user", "write", newRRI("workspace", "ns", "pod")}, true},
		{rule{"user2uuid", "user", "get", newRRI("workspace", "ns", "pod")}, true},
		{rule{"user1uuid", "service-user", "list", newRRI("workspace", "ns", "pod")}, false},
		{rule{"user3uuid", "user", "get", newRRI("workspace", "ns", "pod")}, false},
	}

	for _, tc := range testCases {
		s := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				require.True(t, strings.HasPrefix(r.URL.Path, "/api/v1/authz"))
				require.True(t, r.Method == "GET")

				query := r.URL.Query()
				rl := rule{
					uuid:     query.Get("identity"),
					action:   query.Get("action"),
					userType: query.Get("user_type"),
					rri:      query.Get("object"),
				}
				require.True(t, rl.isValid())
				if tc.valid {
					w.WriteHeader(http.StatusOK)
					return
				}
				resp := types.ResponseError{
					Error: "forbidden",
				}
				w.WriteHeader(http.StatusForbidden)
				require.NoError(t, json.NewEncoder(w).Encode(resp))
			}))
		defer s.Close()
		c := NewTestClient(s)

		err := c.Authorize(tc.uuid, "UserType", tc.action, tc.rri)
		require.True(t, (err == nil) == tc.valid, fmt.Sprintln(tc, err))
	}
}

func newRRI(workspace, ns, resource string) string {
	return fmt.Sprintf("rri:v1:cafebazaar.cloud:%s:godel:/%s/%s", workspace, ns, resource)
}

var concurrentAuthzRequests int = 100
var authzBepaEndpoint string = os.Getenv("BENCHMARK_BEPA_ENDPOINT")
var authzBepaBenchmarkToken string = os.Getenv("BENCHMARK_TOKEN")
var authzBepaBenchmarkAuthCaseUUID string = os.Getenv("BENCHMARK_AUTH_CASE_UUID")
var authzBepaBenchmarkAuthCaseUserType string = os.Getenv("BENCHMARK_AUTH_CASE_USER_TYPE")
var authzBepaBenchmarkAuthCaseAction string = os.Getenv("BENCHMARK_AUTH_CASE_ACTION")
var authzBepaBenchmarkAuthCaseObjectBegining string = os.Getenv("BENCHMARK_AUTH_CASE_OBJECT_BEGINING")
var authzBepaBenchmarkAuthCaseObjectPath string = os.Getenv("BENCHMARK_AUTH_CASE_OBJECT_PATH")
var authzTimeoutDuration time.Duration = 10 * time.Second

func DoSingleBenchmarkAuthz(token string, testCase rule, wg *sync.WaitGroup) {
	serverList := []string{authzBepaEndpoint, authzBepaEndpoint, authzBepaEndpoint}
	c, _ := NewReliableClient(authzBepaBenchmarkToken, serverList, "", "", authzTimeoutDuration)
	c.Authorize(testCase.uuid, testCase.userType, testCase.action, testCase.rri)
	wg.Done()
}

func BenchmarkMultipleValidAuthz(b *testing.B) {
	testCase := rule{
		authzBepaBenchmarkAuthCaseUUID,
		authzBepaBenchmarkAuthCaseUserType,
		authzBepaBenchmarkAuthCaseAction,
		authzBepaBenchmarkAuthCaseObjectBegining + authzBepaBenchmarkAuthCaseObjectPath,
	}
	var wg sync.WaitGroup
	b.Run(fmt.Sprintf("concurrent_iters_%d", concurrentAuthzRequests), func(b *testing.B) {
		var iters int = concurrentAuthzRequests
		for j := 0; j < b.N; j++ {
			wg.Add(iters)
			for i := 0; i < iters; i++ {
				go DoSingleBenchmarkAuthz(authzBepaBenchmarkToken, testCase, &wg)
			}
			wg.Wait()
		}
	})
}

func BenchmarkMultipleInvalidAuthz(b *testing.B) {
	var wg sync.WaitGroup
	b.Run(fmt.Sprintf("concurrent_iters_%d", concurrentAuthzRequests), func(b *testing.B) {
		var iters int = concurrentAuthzRequests
		for j := 0; j < b.N; j++ {
			wg.Add(iters)
			for i := 0; i < iters; i++ {
				randomRRI := authzBepaBenchmarkAuthCaseObjectBegining + randomString(10)
				testCase := rule{
					authzBepaBenchmarkAuthCaseUUID,
					authzBepaBenchmarkAuthCaseUserType,
					authzBepaBenchmarkAuthCaseAction,
					randomRRI,
				}
				go DoSingleBenchmarkAuthz(authzBepaBenchmarkToken, testCase, &wg)
			}
			wg.Wait()
		}
	})
}

func TestIdentifyAndAuthorize(t *testing.T) {
	example := struct {
		token  string
		object string
		action string
	}{"sample-token", "sample-rri-object", "get"}

	s := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			require.True(t, strings.HasPrefix(r.URL.Path, "/api/v1/identify-and-authz"))
			var req types.IdentifyAndAuthorizeReq
			require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
			require.Equal(t, example.token, req.Token)
			w.WriteHeader(http.StatusOK)
		}))

	c := NewTestClient(s)
	err := c.IdentifyAndAuthorize(example.token, example.action, example.object)
	require.NoError(t, err)
	s.Close()

}
