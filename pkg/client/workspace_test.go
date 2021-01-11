package client

import (
	"net/http"
	"regexp"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/internal/pkg/testutils"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"

	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestGetAllWorkspaces(t *testing.T) {
	var objects []types.Workspace
	conf := testutils.TestConfig{
		ClientMethodName: "GetWorkspaces",
		Object:           &objects,
		URLregexp:        regexp.MustCompile(`/workspace/`),
	}
	testutils.DoTestListingAPI(t, conf)
}

func TestGetWorkspace(t *testing.T) {
	var object types.Workspace
	workspace := uuid.NewV4()
	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&workspace},
		ParamNames:       []string{"UUID"},
		URLregexp:        regexp.MustCompile(`^/api/v1/workspace/(.+)/$`),
		ClientMethodName: "GetWorkspace",
	}

	testutils.DoTestReadAPI(t, config)
}

func TestGetWorkspaceByName(t *testing.T) {
	var object types.Workspace
	var workspace string
	faker.FakeData(&workspace)
	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{workspace},
		ParamsInURL:      []interface{}{workspace},
		URLregexp:        regexp.MustCompile(`^/api/v1/user/.+/workspace/name=(.+)/`),
		ClientMethodName: "GetWorkspaceByName",
	}

	testutils.DoTestReadAPI(t, config)
}

func TestGetMyWorkspaces(t *testing.T) {
	var objects []types.Workspace
	conf := testutils.TestConfig{
		ClientMethodName: "GetMyWorkspaces",
		Object:           &objects,
		URLregexp:        regexp.MustCompile(`/user/(.*)/workspace/`),
		CustomHandlerTest: testutils.TestHandlerFunc(func(w *http.ResponseWriter, r *http.Request, regex *regexp.Regexp) bool {
			userUUID := regex.FindStringSubmatch(r.URL.Path)[1]
			require.Equal(t, userUUID, testutils.TestUserUUID)
			return false
		}),
	}
	testutils.DoTestListingAPI(t, conf)

}

func TestCreateWorkspace(t *testing.T) {
	var object types.Workspace
	var name string
	faker.FakeData(&name)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{name},
		ParamNames:       []string{"Name"},
		URLregexp:        regexp.MustCompile(`/workspace/`),
		ClientMethodName: "CreateWorkspace",
	}
	testutils.DoTestCreateAPI(t, config)
}

func TestGetWorkspaceRules(t *testing.T) {
	var objects []types.Rule
	workspaceUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		ClientMethodName: "GetWorkspaceRules",
		Object:           &objects,
		Params:           []interface{}{&workspaceUUID},
	}
	testutils.DoTestListingAPI(t, conf)
}

func TestGetWorkspaceRoles(t *testing.T) {
	var objects []types.Role
	workspaceUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		ClientMethodName: "GetWorkspaceRoles",
		Object:           &objects,
		Params:           []interface{}{&workspaceUUID},
	}
	testutils.DoTestListingAPI(t, conf)
}

func TestGetWorkspaceUsers(t *testing.T) {
	var objects []types.User
	workspaceUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		ClientMethodName: "GetWorkspaceUsers",
		Object:           &objects,
		Params:           []interface{}{&workspaceUUID},
	}
	testutils.DoTestListingAPI(t, conf)
}

func TestDeleteWorkspace(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		URLregexp:        regexp.MustCompile(`/workspace/(.*)/`),
		ClientMethodName: "DeleteWorkspace",
		Params:           []interface{}{&workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID},
	}
	testutils.DoTestDeleteAPI(t, conf)
}
