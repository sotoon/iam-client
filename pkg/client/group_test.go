package client

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/types"
)

func TestGetGroup(t *testing.T) {
	var object types.Group
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	conf := TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/`),
		ClientMethodName: "GetGroup",
		Params:           []interface{}{&workspaceUUID, &groupUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &groupUUID},
	}
	DoTestReadAPI(t, conf)
}

func TestGetAllGroups(t *testing.T) {
	var object []types.Group
	workspaceUUID := uuid.NewV4()
	conf := TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/`),
		ClientMethodName: "GetAllGroups",
		Params:           []interface{}{&workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID},
	}
	DoTestListingAPI(t, conf)
}

func TestDeleteGroup(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()

	conf := TestConfig{
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/`),
		ClientMethodName: "DeleteGroup",
		Params:           []interface{}{&workspaceUUID, &groupUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &groupUUID},
	}
	DoTestDeleteAPI(t, conf)
}

func TestGetGroupByName(t *testing.T) {
	var object types.Group
	var groupName, workspaceName string
	faker.FakeData(&groupName)
	faker.FakeData(&workspaceName)

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{workspaceName, groupName},
		ParamsInURL:      []interface{}{workspaceName, groupName},
		ParamNames:       []string{"Workspace", "Name"},
		URLregexp:        regexp.MustCompile(`^/api/v1/user/.+/workspace/name=(.+)/group/name=(.+)/$`),
		ClientMethodName: "GetGroupByName",
	}

	DoTestReadAPI(t, config)
}

func TestCreateGroup(t *testing.T) {
	var object types.GroupRes
	var name string
	workspaceUUID := uuid.NewV4()
	faker.FakeData(&name)

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{name, &workspaceUUID},
		ParamNames:       []string{"Name"},
		URLregexp:        regexp.MustCompile(`/workspace/`),
		ClientMethodName: "CreateGroup",
	}
	DoTestCreateAPI(t, config)
}

func TestGetGroupUser(t *testing.T) {
	var object types.User
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	userUUID := uuid.NewV4()

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{&workspaceUUID, &groupUUID, &userUUID},
		URLregexp:        regexp.MustCompile(`/workspace/`),
		ClientMethodName: "GetGroupUser",
	}
	DoTestReadAPI(t, config)
}

func TestGetAllGroupUserList(t *testing.T) {
	var object []types.User
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{&workspaceUUID, &groupUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &groupUUID},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/user/`),
		ClientMethodName: "GetAllGroupUserList",
	}
	DoTestListingAPI(t, config)
}

func TestGetAllGroupServiceUserList(t *testing.T) {
	var object []types.ServiceUser
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{&workspaceUUID, &groupUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &groupUUID},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/service-user/`),
		ClientMethodName: "GetAllGroupServiceUserList",
	}
	DoTestListingAPI(t, config)
}

func TestUnbindUserFromGroup(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	userUUID := uuid.NewV4()

	config := TestConfig{
		Params:      []interface{}{&workspaceUUID, &groupUUID, &userUUID},
		ParamsInURL: []interface{}{&workspaceUUID, &groupUUID, &userUUID},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/user/(.+)/`),
		ClientMethodName: "UnbindUserFromGroup",
	}
	DoTestDeleteAPI(t, config)

}

func TestBindGroup(t *testing.T) {
	var object types.GroupUserRes
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	userUUID := uuid.NewV4()
	var groupName string
	faker.FakeData(&groupName)

	config := TestConfig{
		Object:      &object,
		Params:      []interface{}{groupName, &workspaceUUID, &groupUUID, &userUUID},
		ParamsInURL: []interface{}{&workspaceUUID, &groupUUID, &userUUID},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/user/(.+)/`),
		ClientMethodName: "BindGroup",
	}
	DoTestUpdateAPI(t, config, http.MethodPost)
}

func TestBindServiceUserToGroup(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	serviceUserUUID := uuid.NewV4()

	config := TestConfig{
		Params:      []interface{}{&workspaceUUID, &groupUUID, &serviceUserUUID},
		ParamsInURL: []interface{}{&workspaceUUID, &groupUUID, &serviceUserUUID},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/service-user/(.+)/`),
		ClientMethodName: "BindServiceUserToGroup",
	}
	DoTestUpdateAPI(t, config, http.MethodPost)
}

func TestUnbindServiceUserFromGroup(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	serviceUserUUID := uuid.NewV4()

	config := TestConfig{
		Params:      []interface{}{&workspaceUUID, &groupUUID, &serviceUserUUID},
		ParamsInURL: []interface{}{&workspaceUUID, &groupUUID, &serviceUserUUID},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/service-user/(.+)/`),
		ClientMethodName: "UnbindServiceUserFromGroup",
	}
	DoTestDeleteAPI(t, config)
}

func TestGetGroupServiceUser(t *testing.T) {
	var object types.ServiceUser
	workspaceUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	serviceUserUUID := uuid.NewV4()

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{&workspaceUUID, &groupUUID, &serviceUserUUID},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/group/(.+)/service-user/(.+)/`),
		ClientMethodName: "GetGroupServiceUser",
	}
	DoTestReadAPI(t, config)
}
