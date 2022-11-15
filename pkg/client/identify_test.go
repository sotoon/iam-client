package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.cafebazaar.ir/infrastructure/integration/sib/bepa-client/pkg/types"

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
