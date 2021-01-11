package client

import (
	"net/http"
	"regexp"
	"testing"

	"github.com/bxcodec/faker"

	uuid "github.com/satori/go.uuid"

	"git.cafebazaar.ir/infrastructure/bepa-client/internal/pkg/testutils"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
)

func TestCreateRole(t *testing.T) {
	var object types.Role
	var roleName string
	workspaceUUID := uuid.NewV4()
	faker.FakeData(&roleName)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{roleName, &workspaceUUID},
		ParamNames:       []string{"Name"},
		ParamsInURL:      []interface{}{&workspaceUUID},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/role/`),
		ClientMethodName: "CreateRole",
	}
	testutils.DoTestCreateAPI(t, config)
}

func TestGetRoleByName(t *testing.T) {
	var object types.Role
	var roleName, workspaceName string
	faker.FakeData(&roleName)
	faker.FakeData(&workspaceName)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{roleName, workspaceName},
		ParamNames:       []string{"Name"},
		ParamsInURL:      []interface{}{workspaceName, roleName},
		URLregexp:        regexp.MustCompile(`/api/v1/user/.+/workspace/name=(.+)/role/name=(.+)/`),
		ClientMethodName: "GetRoleByName",
	}
	testutils.DoTestReadAPI(t, config)
}

func TestGetRole(t *testing.T) {
	var object types.Role
	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&roleUUID, &workspaceUUID},
		ParamNames:       []string{"Name"},
		ParamsInURL:      []interface{}{&workspaceUUID, &roleUUID},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/role/(.+)/`),
		ClientMethodName: "GetRole",
	}
	testutils.DoTestReadAPI(t, config)
}

func TestGetRoleUsers(t *testing.T) {
	roles := []types.User{}
	workspace := uuid.NewV4()
	role := uuid.NewV4()
	config := testutils.TestConfig{
		Object:           &roles,
		URLregexp:        regexp.MustCompile(`/workspace/(.*)/role/(.*)/user/`),
		ClientMethodName: "GetRoleUsers",
		Params:           []interface{}{&role, &workspace},
		ParamsInURL:      []interface{}{&workspace, &role},
	}
	testutils.DoTestListingAPI(t, config)
}

func TestGetRoleRules(t *testing.T) {
	rules := []types.Rule{}
	workspace := uuid.NewV4()
	role := uuid.NewV4()
	config := testutils.TestConfig{
		Object:    &rules,
		URLregexp: regexp.MustCompile(`/workspace/(.*)/role/(.*)/rule/`),

		ClientMethodName: "GetRoleRules",
		Params:           []interface{}{&role, &workspace},
		ParamsInURL:      []interface{}{&workspace, &role},
	}
	testutils.DoTestListingAPI(t, config)
}

func TestGetUserRoles(t *testing.T) {
	services := []types.RoleBinding{}
	user := uuid.NewV4()
	config := testutils.TestConfig{
		Object:    &services,
		URLregexp: regexp.MustCompile(`/user/(.*)/role/`),

		ClientMethodName: "GetUserRoles",
		Params:           []interface{}{&user},
		ParamsInURL:      []interface{}{&user},
	}
	testutils.DoTestListingAPI(t, config)
}

func TestDeleteRole(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()

	conf := testutils.TestConfig{
		URLregexp:        regexp.MustCompile(`/workspace/(.+)/role/(.+)/`),
		ClientMethodName: "DeleteRole",
		Params:           []interface{}{&workspaceUUID, &roleUUID},
		ParamsInURL:      []interface{}{&roleUUID, &workspaceUUID},
	}
	testutils.DoTestDeleteAPI(t, conf)
}

func TestGetAllRoles(t *testing.T) {
	services := []types.Role{}
	config := testutils.TestConfig{
		Object:           &services,
		URLregexp:        regexp.MustCompile(`^/api/v1/role/$`),
		ClientMethodName: "GetAllRoles",
	}
	testutils.DoTestListingAPI(t, config)
}

func TestBindRoleToUser(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()
	userUUID := uuid.NewV4()
	var params map[string]string
	faker.FakeData(params)

	config := testutils.TestConfig{
		Params:      []interface{}{&workspaceUUID, &roleUUID, &userUUID, params},
		ParamsInURL: []interface{}{&workspaceUUID, &roleUUID, &userUUID},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/role/(.+)/user/(.+)/`),
		ClientMethodName: "BindRoleToUser",
	}
	testutils.DoTestUpdateAPI(t, config, http.MethodPost)
}

func TestUnbindRoleFromUser(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()
	userUUID := uuid.NewV4()
	var params map[string]string
	faker.FakeData(params)

	config := testutils.TestConfig{
		Params:      []interface{}{&workspaceUUID, &roleUUID, &userUUID, params},
		ParamsInURL: []interface{}{&workspaceUUID, &roleUUID, &userUUID},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/role/(.+)/user/(.+)/`),
		ClientMethodName: "UnbindRoleFromUser",
	}
	testutils.DoTestDeleteAPI(t, config)
}

func TestBindRoleToServiceUser(t *testing.T) {

	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()
	userUUID := uuid.NewV4()
	var params map[string]string
	faker.FakeData(params)

	config := testutils.TestConfig{
		Params:      []interface{}{&workspaceUUID, &roleUUID, &userUUID, params},
		ParamsInURL: []interface{}{&workspaceUUID, &roleUUID, &userUUID, params},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/role/(.+)/service-user/(.+)/`),
		ClientMethodName: "BindRoleToServiceUser",
	}
	testutils.DoTestUpdateAPI(t, config, http.MethodPost)
}

func TestUnbindRoleFromServiceUser(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()
	userUUID := uuid.NewV4()
	var params map[string]string
	faker.FakeData(params)

	config := testutils.TestConfig{
		Params:      []interface{}{&workspaceUUID, &roleUUID, &userUUID, params},
		ParamsInURL: []interface{}{&workspaceUUID, &roleUUID, &userUUID},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/role/(.+)/service-user/(.+)/`),
		ClientMethodName: "UnbindRoleFromServiceUser",
	}
	testutils.DoTestDeleteAPI(t, config)
}

func TestGetRoleServiceUsers(t *testing.T) {
	servicesUsers := []types.ServiceUser{}
	roleUUID := uuid.NewV4()
	workspaceUUID := uuid.NewV4()
	config := testutils.TestConfig{
		Object:           &servicesUsers,
		URLregexp:        regexp.MustCompile(`^/api/v1/workspace/(.*)/role/(.*)/service-user/$`),
		ClientMethodName: "GetRoleServiceUsers",
		Params:           []interface{}{&roleUUID, &workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &roleUUID},
	}
	testutils.DoTestListingAPI(t, config)
}

func TestGetRoleGroups(t *testing.T) {
	groups := []types.Group{}
	roleUUID := uuid.NewV4()
	workspaceUUID := uuid.NewV4()
	config := testutils.TestConfig{
		Object:           &groups,
		URLregexp:        regexp.MustCompile(`^/api/v1/workspace/(.+)/role/(.+)/group/$`),
		ClientMethodName: "GetRoleGroups",
		Params:           []interface{}{&roleUUID, &workspaceUUID},
		ParamsInURL:      []interface{}{&workspaceUUID, &roleUUID},
	}
	testutils.DoTestListingAPI(t, config)
}

func TestBindRoleToGroup(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	var params map[string]string
	faker.FakeData(params)

	config := testutils.TestConfig{
		Params:      []interface{}{&workspaceUUID, &roleUUID, &groupUUID, params},
		ParamsInURL: []interface{}{&workspaceUUID, &roleUUID, &groupUUID, params},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/role/(.+)/group/(.+)/`),
		ClientMethodName: "BindRoleToGroup",
	}
	testutils.DoTestUpdateAPI(t, config, http.MethodPost)
}

func TestUnbindRoleFromGroup(t *testing.T) {
	workspaceUUID := uuid.NewV4()
	roleUUID := uuid.NewV4()
	groupUUID := uuid.NewV4()
	var params map[string]string
	faker.FakeData(params)

	config := testutils.TestConfig{
		Params:      []interface{}{&workspaceUUID, &roleUUID, &groupUUID, params},
		ParamsInURL: []interface{}{&workspaceUUID, &roleUUID, &groupUUID},

		URLregexp:        regexp.MustCompile(`/workspace/(.+)/role/(.+)/group/(.+)/`),
		ClientMethodName: "UnbindRoleFromGroup",
	}
	testutils.DoTestDeleteAPI(t, config)
}
