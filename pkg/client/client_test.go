package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"

	"github.com/stretchr/testify/require"
)

type rule struct {
	uuid   string
	action string
	rri    string
}

var rriRegex = regexp.MustCompile(`^rri:v1:cafebazaar\.cloud:(.+):godel:(\/.+)*\/?$`)

func (r1 rule) equal(r2 rule) bool {
	return r1.uuid == r2.uuid &&
		r1.rri == r2.rri &&
		r1.action == r2.action
}

func (r rule) isValid() bool {
	return r.uuid != "" &&
		r.action != "" &&
		rriRegex.MatchString(r.rri)
}

func TestAuthorization(t *testing.T) {
	verified := []rule{
		{"user1uuid", "get", newRRI("workspace", "ns", "pod")},
		{"user2uuid", "get", newRRI("workspace", "ns", "pod")},
		{"user2uuid", "write", newRRI("workspace", "ns", "pod")},
	}

	s := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			require.True(t, strings.HasPrefix(r.URL.Path, "/api/v1/authz"))

			query := r.URL.Query()
			rl := rule{
				uuid:   query.Get("identity"),
				action: query.Get("action"),
				rri:    query.Get("object"),
			}
			require.True(t, rl.isValid())

			for _, v := range verified {
				if v.equal(rl) {
					resp := types.VerifRes{
						Message: http.StatusText(http.StatusOK),
					}
					require.NoError(t, json.NewEncoder(w).Encode(resp))
					w.WriteHeader(http.StatusOK)
					return
				}
			}
			resp := types.ResponseError{
				Error: "forbidden",
			}
			w.WriteHeader(http.StatusForbidden)
			require.NoError(t, json.NewEncoder(w).Encode(resp))
		}))
	defer s.Close()

	c, err := NewClient("sampleaccesstoken", s.URL, "default_workspace", "user_uuid")
	require.NoError(t, err)

	testCases := []struct {
		rule
		valid bool
	}{
		{rule{"user1uuid", "get", newRRI("workspace", "ns", "pod")}, true},
		{rule{"user2uuid", "write", newRRI("workspace", "ns", "pod")}, true},
		{rule{"user2uuid", "get", newRRI("workspace", "ns", "pod")}, true},
		{rule{"user1uuid", "list", newRRI("workspace", "ns", "pod")}, false},
		{rule{"user3uuid", "get", newRRI("workspace", "ns", "pod")}, false},
	}

	for _, tc := range testCases {
		err := c.Authorize(tc.uuid, tc.action, tc.rri)
		require.True(t, (err == nil) == tc.valid)
	}
}

func TestIdentification(t *testing.T) {
	testCases := []struct {
		uuid  string
		token string
		found bool
	}{
		{"user1uuid", "sampleusertoken", true},
		{"user2uuid", "sampleusertoken", true},
		{"user3uuid", "sampleusertoken", true},
		{"user4uuid", "sampleusertoken", false},
		{"user5uuid", "sampleusertoken", false},
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

		c, err := NewClient("sampleaccesstoken", s.URL, "default_workspace", "user_uuid")
		require.NoError(t, err)

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

func newRRI(workspace, ns, resource string) string {
	return fmt.Sprintf("rri:v1:cafebazaar.cloud:%s:godel:/%s/%s", workspace, ns, resource)
}
