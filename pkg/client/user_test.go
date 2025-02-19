package client

import (
	"net/http"
	"regexp"
	"testing"

	"git.platform.sotoon.ir/iam/golang-bepa-client/pkg/routes"
	"git.platform.sotoon.ir/iam/golang-bepa-client/pkg/types"
	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
)

func TestCreateUser(t *testing.T) {
	var object types.User
	var userName, email, password string
	faker.FakeData(&userName)
	faker.FakeData(&email)
	faker.FakeData(&password)

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{userName, email, password},
		ParamNames:       []string{"Name", "Email", ""},
		URLregexp:        regexp.MustCompile(`/api/v1/user/`),
		ClientMethodName: "CreateUser",
	}
	DoTestCreateAPI(t, config)
}

func TestGetSecret(t *testing.T) {
	var object types.UserSecret
	user := uuid.NewV4()
	config := TestConfig{
		Object:      &object,
		Params:      []interface{}{&user},
		ParamsInURL: []interface{}{&user},

		URLregexp:        regexp.MustCompile(`/user/(.+)/secret/`),
		ClientMethodName: "GetSecret",
	}

	DoTestReadAPI(t, config)
}

func TestRevokeSecret(t *testing.T) {
	var object types.UserSecret
	user := uuid.NewV4()
	config := TestConfig{
		Object:      &object,
		Params:      []interface{}{&user},
		ParamsInURL: []interface{}{&user},

		URLregexp:        regexp.MustCompile(`/user/(.+)/secret/`),
		ClientMethodName: "RevokeSecret",
	}

	DoTestUpdateAPI(t, config, http.MethodPost)
}

func TestCreateUserTokenByCreds(t *testing.T) {
	var object types.UserToken
	var email, password string
	faker.FakeData(&email)
	faker.FakeData(&password)

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{email, password},
		URLregexp:        regexp.MustCompile(`/api/v1/authn/`),
		ClientMethodName: "CreateUserTokenByCreds",
	}

	DoTestCreateAPI(t, config)
}

func TestUpdateUser(t *testing.T) {
	var object types.User
	user := uuid.NewV4()
	var userName, email, password string
	faker.FakeData(&userName)
	faker.FakeData(&email)
	faker.FakeData(&password)

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{&user, userName, email, password},
		ParamNames:       []string{"", "Name", "Email", ""},
		ParamsInURL:      []interface{}{&user},
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/`),
		ClientMethodName: "UpdateUser",
	}
	DoTestUpdateAPI(t, config, http.MethodPatch)
}

func TestGetUserByName(t *testing.T) {
	var object types.User
	var userName string
	faker.FakeData(&userName)
	var workspaceId uuid.UUID = uuid.NewV4()

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{userName, &workspaceId},
		URLregexp:        routeToRegex(routes.RouteWorkspaceUserGetByEmail),
		ClientMethodName: "GetUserByName",
	}
	DoTestReadAPI(t, config)
}

func TestSuspendUserInWorkspace(t *testing.T) {
	var workspaceId uuid.UUID = uuid.NewV4()
	var userId uuid.UUID = uuid.NewV4()

	config := TestConfig{
		Params:           []interface{}{&workspaceId, &userId},
		URLregexp:        routeToRegex(routes.RouteSuspendUserInWorkspace),
		ClientMethodName: "SuspendUserInWorkspace",
	}
	DoTestUpdateAPI(t, config, "PUT")
}

func TestActivateUserInWorkspace(t *testing.T) {
	var workspaceId uuid.UUID = uuid.NewV4()
	var userId uuid.UUID = uuid.NewV4()

	config := TestConfig{
		Params:           []interface{}{&workspaceId, &userId},
		URLregexp:        routeToRegex(routes.RouteActivateUserInWorkspace),
		ClientMethodName: "ActivateUserInWorkspace",
	}
	DoTestUpdateAPI(t, config, "PUT")
}

func TestGetUserByEmail(t *testing.T) {
	var object types.User
	var userEmail string
	faker.FakeData(&userEmail)
	var workspaceId uuid.UUID = uuid.NewV4()

	config := TestConfig{
		Object:           []types.User{object},
		Params:           []interface{}{userEmail, &workspaceId},
		URLregexp:        routeToRegex(routes.RouteWorkspaceGetUsers),
		ClientMethodName: "GetUserByEmail",
	}
	DoTestReadAPI(t, config)
}

func TestGetMySelf(t *testing.T) {
	var object types.User

	config := TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/`),
		ClientMethodName: "GetMySelf",
	}
	DoTestReadAPI(t, config)
}

func TestGetUser(t *testing.T) {
	var object types.User
	user := uuid.NewV4()

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{&user},
		ParamsInURL:      []interface{}{&user},
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/`),
		ClientMethodName: "GetUser",
	}
	DoTestReadAPI(t, config)
}

func TestGetUsers(t *testing.T) {
	var object []types.User

	config := TestConfig{
		Object: &object,

		URLregexp:        regexp.MustCompile(`/api/v1/user/`),
		ClientMethodName: "GetUsers",
	}
	DoTestListingAPI(t, config)
}

func TestDeleteUser(t *testing.T) {
	user := uuid.NewV4()

	config := TestConfig{
		Params:           []interface{}{&user},
		ParamsInURL:      []interface{}{&user},
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/`),
		ClientMethodName: "DeleteUser",
	}
	DoTestDeleteAPI(t, config)
}

func TestDeleteMySelf(t *testing.T) {

	config := TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/`),
		ClientMethodName: "DeleteMySelf",
	}
	DoTestDeleteAPI(t, config)
}

func TestAddUserToWorkspace(t *testing.T) {
	user := uuid.NewV4()
	workspace := uuid.NewV4()

	config := TestConfig{
		Params:           []interface{}{&workspace, &user},
		ParamsInURL:      []interface{}{&user, &workspace},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/user/(.+)/`),
		ClientMethodName: "AddUserToWorkspace",
	}
	DoTestUpdateAPI(t, config, http.MethodPost)
}

func TestRemoveUserFromWorkspace(t *testing.T) {
	user := uuid.NewV4()
	workspace := uuid.NewV4()
	config := TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/user/(.+)/`),
		Params:           []interface{}{&workspace, &user},
		ParamsInURL:      []interface{}{&user, &workspace},
		ClientMethodName: "RemoveUserFromWorkspace",
	}
	DoTestDeleteAPI(t, config)
}

func TestSetMyPassword(t *testing.T) {
	var password string
	var object types.User
	faker.FakeData(&password)

	conf := TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/user/(.+)/`),
		ClientMethodName: "SetMyPassword",
		Params:           []interface{}{password},
	}
	DoTestUpdateAPI(t, conf, http.MethodPatch)
}

func TestSetMyName(t *testing.T) {
	var name string
	var object types.User
	faker.FakeData(&name)

	conf := TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/user/(.+)/`),
		ClientMethodName: "SetMyName",
		Params:           []interface{}{name},
	}
	DoTestUpdateAPI(t, conf, http.MethodPatch)
}

func TestSetMyEmail(t *testing.T) {
	var email string
	var object types.User
	faker.FakeData(&email)

	conf := TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/user/(.+)/`),
		ClientMethodName: "SetMyEmail",
		Params:           []interface{}{email},
	}
	DoTestUpdateAPI(t, conf, http.MethodPatch)
}

func TestInviteUser(t *testing.T) {
	var object types.InvitationInfo
	var email string
	workspace := uuid.NewV4()
	faker.FakeData(&email)

	conf := TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/invite/`),
		Params:           []interface{}{&workspace, email},
		ClientMethodName: "InviteUser",
	}
	DoTestUpdateAPI(t, conf, http.MethodPost)
}

func TestJoinByInvitationToken(t *testing.T) {
	var object types.User
	var name, password, invitationToken string
	faker.FakeData(&name)
	faker.FakeData(&password)
	faker.FakeData(&invitationToken)

	conf := TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/accept-invitation/(.+)/`),
		ClientMethodName: "JoinByInvitationToken",
		Params:           []interface{}{name, password, invitationToken},
	}
	DoTestCreateAPI(t, conf)
}

func TestSuspendUser(t *testing.T) {
	userUUID := uuid.NewV4()

	conf := TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/suspend/`),
		ClientMethodName: "SuspendUser",
		Params:           []interface{}{&userUUID},
	}
	DoTestUpdateAPI(t, conf, http.MethodPut)
}

func TestActivateUser(t *testing.T) {
	userUUID := uuid.NewV4()

	conf := TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/activate`),
		ClientMethodName: "ActivateUser",
		Params:           []interface{}{&userUUID},
	}
	DoTestUpdateAPI(t, conf, http.MethodPut)
}
