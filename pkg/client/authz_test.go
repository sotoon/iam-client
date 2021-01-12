package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/internal/pkg/testutils"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"

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
		c := testutils.NewTestClient(s)

		err := c.Authorize(tc.uuid, "UserType", tc.action, tc.rri)
		require.True(t, (err == nil) == tc.valid)
	}
}

func newRRI(workspace, ns, resource string) string {
	return fmt.Sprintf("rri:v1:cafebazaar.cloud:%s:godel:/%s/%s", workspace, ns, resource)
}
