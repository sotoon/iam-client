package client

import (
	"regexp"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
)

func TestCreateMyUserTokenWithToken(t *testing.T) {
	var object types.UserToken
	var secret string
	faker.FakeData(&secret)

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{secret},
		URLregexp:        regexp.MustCompile(`/user/(.+)/user-token/`),
		ClientMethodName: "CreateMyUserTokenWithToken",
	}
	DoTestCreateAPI(t, config)
}

func TestGetMyUserToken(t *testing.T) {
	var object types.UserToken
	tokenUUID := uuid.NewV4()
	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{&tokenUUID},
		ParamNames:       []string{"UUID"},
		URLregexp:        regexp.MustCompile(`/user/.+/user-token/(.+)/`),
		ClientMethodName: "GetMyUserToken",
	}

	DoTestReadAPI(t, config)
}

func TestGetAllMyUserTokens(t *testing.T) {
	var objects []types.UserToken
	conf := TestConfig{
		ClientMethodName: "GetAllMyUserTokens",
		Object:           &objects,
		URLregexp:        regexp.MustCompile(`/user/.+/user-token/`),
	}
	DoTestListingAPI(t, conf)
}

func TestDeleteMyUserToken(t *testing.T) {
	tokenUUID := uuid.NewV4()
	conf := TestConfig{
		URLregexp:        regexp.MustCompile(`/user/.+/user-token/(.*)/`),
		ClientMethodName: "DeleteMyUserToken",
		Params:           []interface{}{&tokenUUID},
		ParamsInURL:      []interface{}{&tokenUUID},
	}
	DoTestDeleteAPI(t, conf)
}
