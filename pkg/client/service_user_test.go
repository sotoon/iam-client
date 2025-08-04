package client

import (
	"regexp"
	"testing"

	"github.com/bxcodec/faker"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/types"
)

func TestGetServiceUser(t *testing.T) {
	var object types.ServiceUser

	workspace := uuid.NewV4()
	ServiceUser := uuid.NewV4()
	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{&workspace, &ServiceUser},
		ParamNames:       []string{"UUID"},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/(.+)/`),
		ClientMethodName: "GetServiceUser",
	}

	DoTestReadAPI(t, config)
}

func TestGetServiceUsers(t *testing.T) {
	var objects []types.ServiceUser
	workspaceUUID := uuid.NewV4()
	conf := TestConfig{
		ClientMethodName: "GetServiceUsers",
		Object:           &objects,
		Params:           []interface{}{&workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/`),
	}
	DoTestListingAPI(t, conf)
}

func TestDeleteServiceUser(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	serviceUserUUID := uuid.NewV4()

	conf := TestConfig{
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/(.+)/`),
		ClientMethodName: "DeleteServiceUser",
		Params:           []interface{}{&workspaceUUID, &serviceUserUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &serviceUserUUID},
	}
	DoTestDeleteAPI(t, conf)
}

func TestGetServiceUserByName(t *testing.T) {
	var object types.ServiceUser

	var workspaceName string
	var name string
	faker.FakeData(&name)
	faker.FakeData(&workspaceName)

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{workspaceName, name},
		ParamsInURL:      []interface{}{workspaceName, name},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/workspace=(.+)/service-user/name=(.+)/`),
		ClientMethodName: "GetServiceUserByName",
	}

	DoTestReadAPI(t, config)
}

func TestCreateServiceUser(t *testing.T) {
	var object types.ServiceUser
	var name string
	workspaceUUID := uuid.NewV4()
	faker.FakeData(&name)

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{name, &workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/`),
		ClientMethodName: "CreateServiceUser",
	}
	DoTestCreateAPI(t, config)
}

func TestCreateServiceUserToken(t *testing.T) {
	var object types.ServiceUserToken
	workspaceUUID := uuid.NewV4()
	serviceUserUUID := uuid.NewV4()

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{&serviceUserUUID, &workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &serviceUserUUID},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/(.+)/token`),
		ClientMethodName: "CreateServiceUserToken",
	}
	DoTestCreateAPI(t, config)
}

func TestGetWorkspaceServiceUserTokenList(t *testing.T) {
	var objects []types.ServiceUserToken
	workspaceUUID := uuid.NewV4()
	serviceUserUUID := uuid.NewV4()
	conf := TestConfig{
		ClientMethodName: "GetWorkspaceServiceUserTokenList",
		Params:           []interface{}{&serviceUserUUID, &workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &serviceUserUUID},
		Object:           &objects,
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/(.+)/token`),
	}
	DoTestListingAPI(t, conf)
}

func TestDeleteServiceUserToken(t *testing.T) {
	tokenUUID := uuid.NewV4()
	workspaceUUID := uuid.NewV4()
	serviceUserUUID := uuid.NewV4()
	conf := TestConfig{
		URLregexp: regexp.MustCompile(`/workspace/(.+)/service-user/(.+)/token/(.+)/`),

		ClientMethodName: "DeleteServiceUserToken",
		Params:           []interface{}{&serviceUserUUID, &workspaceUUID, &tokenUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &serviceUserUUID, &tokenUUID},
	}
	DoTestDeleteAPI(t, conf)
}
