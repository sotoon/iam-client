package client

import (
	"net/http"
	"regexp"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/internal/pkg/testutils"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
)

func TestGetGroup(t *testing.T) {
	var object types.Group
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/`),
		ClientMethodName: "GetGroup",
		Params:           []interface{}{&workspaceUUID, &groupUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &groupUUID},
	}
	testutils.DoTestReadAPI(t, conf)
}

func TestGetAllGroups(t *testing.T) {
	var object []types.Group
	workspaceUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/`),
		ClientMethodName: "GetAllGroups",
		Params:           []interface{}{&workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID},
	}
	testutils.DoTestListingAPI(t, conf)
}

func TestDeleteGroup(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()

	conf := testutils.TestConfig{
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/`),
		ClientMethodName: "DeleteGroup",
		Params:           []interface{}{&workspaceUUID, &groupUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &groupUUID},
	}
	testutils.DoTestDeleteAPI(t, conf)
}

func TestGetGroupByName(t *testing.T) {
	var object types.Group
	var groupName, workspaceName string
	faker.FakeData(&groupName)
	faker.FakeData(&workspaceName)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{workspaceName, groupName},
		ParamsInURL:      []interface{}{workspaceName, groupName},
		ParamNames:       []string{"Workspace", "Name"},
		URLregexp:        regexp.MustCompile(`^/api/v1/user/.+/workspace/name=(.+)/group/name=(.+)/$`),
		ClientMethodName: "GetGroupByName",
	}

	testutils.DoTestReadAPI(t, config)
}

func TestCreateGroup(t *testing.T) {
	var object types.Group
	var name string
	workspaceUUID := uuid.NewV4()
	faker.FakeData(&name)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{name, &workspaceUUID},
		ParamNames:       []string{"Name"},
		URLregexp:        regexp.MustCompile(`/workspace/`),
		ClientMethodName: "CreateGroup",
	}
	testutils.DoTestCreateAPI(t, config)
}

func TestGetGroupUser(t *testing.T) {
	var object types.User
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	userUUID := uuid.NewV4()

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&workspaceUUID, &groupUUID, &userUUID},
		URLregexp:        regexp.MustCompile(`/workspace/`),
		ClientMethodName: "GetGroupUser",
	}
	testutils.DoTestReadAPI(t, config)
}

func TestGetAllGroupUsers(t *testing.T) {
	var object []types.User
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&workspaceUUID, &groupUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &groupUUID},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/user/`),
		ClientMethodName: "GetAllGroupUsers",
	}
	testutils.DoTestListingAPI(t, config)
}

func TestUnbindUserFromGroup(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	userUUID := uuid.NewV4()

	config := testutils.TestConfig{
		Params:      []interface{}{&workspaceUUID, &groupUUID, &userUUID},
		ParamsInURL: []interface{}{&workspaceUUID, &groupUUID, &userUUID},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/user/(.+)/`),
		ClientMethodName: "UnbindUserFromGroup",
	}
	testutils.DoTestDeleteAPI(t, config)

}

func TestBindGroup(t *testing.T) {
	var object types.GroupUserRes
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	userUUID := uuid.NewV4()
	var groupName string
	faker.FakeData(&groupName)

	config := testutils.TestConfig{
		Object:      &object,
		Params:      []interface{}{groupName, &workspaceUUID, &groupUUID, &userUUID},
		ParamsInURL: []interface{}{&workspaceUUID, &groupUUID, &userUUID},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/user/(.+)/`),
		ClientMethodName: "BindGroup",
	}
	testutils.DoTestUpdateAPI(t, config, http.MethodPost)
}
