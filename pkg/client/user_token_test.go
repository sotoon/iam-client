package client

import (
	"regexp"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/internal/pkg/testutils"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
)

func TestCreateMyUserTokenWithToken(t *testing.T) {
	var object types.UserToken
	var secret string
	faker.FakeData(&secret)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{secret},
		URLregexp:        regexp.MustCompile(`/user/(.+)/user-token/`),
		ClientMethodName: "CreateMyUserTokenWithToken",
	}
	testutils.DoTestCreateAPI(t, config)
}

func TestGetMyUserToken(t *testing.T) {
	var object types.UserToken
	tokenUUID := uuid.NewV4()
	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&tokenUUID},
		ParamNames:       []string{"UUID"},
		URLregexp:        regexp.MustCompile(`/user/.+/user-token/(.+)/`),
		ClientMethodName: "GetMyUserToken",
	}

	testutils.DoTestReadAPI(t, config)
}

func TestGetAllMyUserTokens(t *testing.T) {
	var objects []types.UserToken
	conf := testutils.TestConfig{
		ClientMethodName: "GetAllMyUserTokens",
		Object:           &objects,
		URLregexp:        regexp.MustCompile(`/user/.+/user-token/`),
	}
	testutils.DoTestListingAPI(t, conf)
}

func TestDeleteMyUserToken(t *testing.T) {
	tokenUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		URLregexp:        regexp.MustCompile(`/user/.+/user-token/(.*)/`),
		ClientMethodName: "DeleteMyUserToken",
		Params:           []interface{}{&tokenUUID},
		ParamsInURL:      []interface{}{&tokenUUID},
	}
	testutils.DoTestDeleteAPI(t, conf)
}
