package client

import (
	"regexp"
	"testing"

	"github.com/bxcodec/faker"

	"git.cafebazaar.ir/infrastructure/bepa-client/internal/pkg/testutils"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	uuid "github.com/satori/go.uuid"
)

func TestGetServiceUser(t *testing.T) {
	var object types.ServiceUser

	workspace := uuid.NewV4()
	ServiceUser := uuid.NewV4()
	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&workspace, &ServiceUser},
		ParamNames:       []string{"UUID"},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/(.+)/`),
		ClientMethodName: "GetServiceUser",
	}

	testutils.DoTestReadAPI(t, config)
}

func TestGetServiceUsers(t *testing.T) {
	var objects []types.ServiceUser
	workspaceUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		ClientMethodName: "GetServiceUsers",
		Object:           &objects,
		Params:           []interface{}{&workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/`),
	}
	testutils.DoTestListingAPI(t, conf)
}

func TestDeleteServiceUser(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	serviceUserUUID := uuid.NewV4()

	conf := testutils.TestConfig{
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/(.+)/`),
		ClientMethodName: "DeleteServiceUser",
		Params:           []interface{}{&workspaceUUID, &serviceUserUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &serviceUserUUID},
	}
	testutils.DoTestDeleteAPI(t, conf)
}

func TestGetServiceUserByName(t *testing.T) {
	var object types.ServiceUser

	var workspaceName string
	var name string
	faker.FakeData(&name)
	faker.FakeData(&workspaceName)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{workspaceName, name},
		ParamsInURL:      []interface{}{workspaceName, name},
		URLregexp:        regexp.MustCompile(`/api/v1/user/.+/workspace/name=(.+)/service-user/name=(.+)/`),
		ClientMethodName: "GetServiceUserByName",
	}

	testutils.DoTestReadAPI(t, config)
}

func TestCreateServiceUser(t *testing.T) {
	var object types.ServiceUser
	var name string
	workspaceUUID := uuid.NewV4()
	faker.FakeData(&name)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{name, &workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/`),
		ClientMethodName: "CreateServiceUser",
	}
	testutils.DoTestCreateAPI(t, config)
}

func TestCreateServiceUserToken(t *testing.T) {
	var object types.ServiceUserToken
	workspaceUUID := uuid.NewV4()
	serviceUserUUID := uuid.NewV4()

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&serviceUserUUID, &workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &serviceUserUUID},
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/(.+)/token`),
		ClientMethodName: "CreateServiceUserToken",
	}
	testutils.DoTestCreateAPI(t, config)
}

func TestGetAllServiceUserToken(t *testing.T) {
	var objects []types.ServiceUserToken
	workspaceUUID := uuid.NewV4()
	serviceUserUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		ClientMethodName: "GetAllServiceUserToken",
		Params:           []interface{}{&serviceUserUUID, &workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &serviceUserUUID},
		Object:           &objects,
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/service-user/(.+)/token`),
	}
	testutils.DoTestListingAPI(t, conf)
}

func TestDeleteServiceUserToken(t *testing.T) {
	tokenUUID := uuid.NewV4()
	workspaceUUID := uuid.NewV4()
	serviceUserUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		URLregexp: regexp.MustCompile(`/workspace/(.+)/service-user/(.+)/token/(.+)/`),

		ClientMethodName: "DeleteServiceUserToken",
		Params:           []interface{}{&serviceUserUUID, &workspaceUUID, &tokenUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &serviceUserUUID, &tokenUUID},
	}
	testutils.DoTestDeleteAPI(t, conf)
}
