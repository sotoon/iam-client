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

func TestCreateUser(t *testing.T) {
	var object types.User
	var userName, email, password string
	faker.FakeData(&userName)
	faker.FakeData(&email)
	faker.FakeData(&password)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{userName, email, password},
		ParamNames:       []string{"Name", "Email", ""},
		URLregexp:        regexp.MustCompile(`/api/v1/user/`),
		ClientMethodName: "CreateUser",
	}
	testutils.DoTestCreateAPI(t, config)
}

func TestGetSecret(t *testing.T) {
	var object types.UserSecret
	user := uuid.NewV4()
	config := testutils.TestConfig{
		Object:      &object,
		Params:      []interface{}{&user},
		ParamsInURL: []interface{}{&user},

		URLregexp:        regexp.MustCompile(`/user/(.+)/secret/`),
		ClientMethodName: "GetSecret",
	}

	testutils.DoTestReadAPI(t, config)
}

func TestRevokeSecret(t *testing.T) {
	var object types.UserSecret
	user := uuid.NewV4()
	config := testutils.TestConfig{
		Object:      &object,
		Params:      []interface{}{&user},
		ParamsInURL: []interface{}{&user},

		URLregexp:        regexp.MustCompile(`/user/(.+)/secret/`),
		ClientMethodName: "RevokeSecret",
	}

	testutils.DoTestUpdateAPI(t, config, http.MethodPost)
}

func TestCreateUserTokenByCreds(t *testing.T) {
	var object types.UserToken
	var email, password string
	faker.FakeData(&email)
	faker.FakeData(&password)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{email, password},
		URLregexp:        regexp.MustCompile(`/api/v1/authn/`),
		ClientMethodName: "CreateUserTokenByCreds",
	}

	testutils.DoTestCreateAPI(t, config)
}

func TestUpdateUser(t *testing.T) {
	var object types.User
	user := uuid.NewV4()
	var userName, email, password string
	faker.FakeData(&userName)
	faker.FakeData(&email)
	faker.FakeData(&password)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&user, userName, email, password},
		ParamNames:       []string{"", "Name", "Email", ""},
		ParamsInURL:      []interface{}{&user},
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/`),
		ClientMethodName: "UpdateUser",
	}
	testutils.DoTestUpdateAPI(t, config, http.MethodPatch)
}

func TestGetUserByName(t *testing.T) {

}

func TestGetMySelf(t *testing.T) {
	var object types.User

	config := testutils.TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/`),
		ClientMethodName: "GetMySelf",
	}
	testutils.DoTestReadAPI(t, config)
}

func TestGetUser(t *testing.T) {
	var object types.User
	user := uuid.NewV4()

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&user},
		ParamsInURL:      []interface{}{&user},
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/`),
		ClientMethodName: "GetUser",
	}
	testutils.DoTestReadAPI(t, config)
}

func TestGetUsers(t *testing.T) {
	var object []types.User

	config := testutils.TestConfig{
		Object: &object,

		URLregexp:        regexp.MustCompile(`/api/v1/user/`),
		ClientMethodName: "GetUsers",
	}
	testutils.DoTestListingAPI(t, config)
}

func TestDeleteUser(t *testing.T) {
	user := uuid.NewV4()

	config := testutils.TestConfig{
		Params:           []interface{}{&user},
		ParamsInURL:      []interface{}{&user},
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/`),
		ClientMethodName: "DeleteUser",
	}
	testutils.DoTestDeleteAPI(t, config)
}

func TestDeleteMySelf(t *testing.T) {

	config := testutils.TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/`),
		ClientMethodName: "DeleteMySelf",
	}
	testutils.DoTestDeleteAPI(t, config)
}

func TestAddUserToWorkspace(t *testing.T) {
	user := uuid.NewV4()
	workspace := uuid.NewV4()

	config := testutils.TestConfig{
		Params:           []interface{}{&workspace, &user},
		ParamsInURL:      []interface{}{&user, &workspace},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/user/(.+)/`),
		ClientMethodName: "AddUserToWorkspace",
	}
	testutils.DoTestUpdateAPI(t, config, http.MethodPost)
}

func TestRemoveUserFromWorkspace(t *testing.T) {
	user := uuid.NewV4()
	workspace := uuid.NewV4()
	config := testutils.TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/user/(.+)/`),
		Params:           []interface{}{&workspace, &user},
		ParamsInURL:      []interface{}{&user, &workspace},
		ClientMethodName: "RemoveUserFromWorkspace",
	}
	testutils.DoTestDeleteAPI(t, config)
}

func TestSetMyPassword(t *testing.T) {
	var password string
	var object types.User
	faker.FakeData(&password)

	conf := testutils.TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/user/(.+)/`),
		ClientMethodName: "SetMyPassword",
		Params:           []interface{}{password},
	}
	testutils.DoTestUpdateAPI(t, conf, http.MethodPatch)
}

func TestSetMyName(t *testing.T) {
	var name string
	var object types.User
	faker.FakeData(&name)

	conf := testutils.TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/user/(.+)/`),
		ClientMethodName: "SetMyName",
		Params:           []interface{}{name},
	}
	testutils.DoTestUpdateAPI(t, conf, http.MethodPatch)
}

func TestSetMyEmail(t *testing.T) {
	var email string
	var object types.User
	faker.FakeData(&email)

	conf := testutils.TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/user/(.+)/`),
		ClientMethodName: "SetMyEmail",
		Params:           []interface{}{email},
	}
	testutils.DoTestUpdateAPI(t, conf, http.MethodPatch)
}

func TestInviteUser(t *testing.T) {
	var object types.InvitationInfo
	var email string
	workspace := uuid.NewV4()
	faker.FakeData(&email)

	conf := testutils.TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/(.+)/invite/`),
		Params:           []interface{}{&workspace, email},
		ClientMethodName: "InviteUser",
	}
	testutils.DoTestUpdateAPI(t, conf, http.MethodPost)
}

func TestJoinByInvitationToken(t *testing.T) {
	var object types.User
	var name, password, invitationToken string
	faker.FakeData(&name)
	faker.FakeData(&password)
	faker.FakeData(&invitationToken)

	conf := testutils.TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/accept-invitation/(.+)/`),
		ClientMethodName: "JoinByInvitationToken",
		Params:           []interface{}{name, password, invitationToken},
	}
	testutils.DoTestCreateAPI(t, conf)
}

func TestSuspendUser(t *testing.T) {
	userUUID := uuid.NewV4()

	conf := testutils.TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/suspend/`),
		ClientMethodName: "SuspendUser",
		Params:           []interface{}{&userUUID},
	}
	testutils.DoTestUpdateAPI(t, conf, http.MethodPut)
}

func TestActivateUser(t *testing.T) {
	userUUID := uuid.NewV4()

	conf := testutils.TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/user/(.+)/activate`),
		ClientMethodName: "ActivateUser",
		Params:           []interface{}{&userUUID},
	}
	testutils.DoTestUpdateAPI(t, conf, http.MethodPut)
}
