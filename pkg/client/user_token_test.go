package client

import (
	"regexp"
	"testing"

	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/types"
)

func TestCreateMyUserToken(t *testing.T) {
	var object types.UserToken
	var secret string
	faker.FakeData(&secret)

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{secret},
		URLregexp:        regexp.MustCompile(`/user/(.+)/user-token/`),
		ClientMethodName: "CreateMyUserToken",
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

func TestGetAllMyUserTokenList(t *testing.T) {
	var objects []types.UserToken
	conf := TestConfig{
		ClientMethodName: "GetAllMyUserTokenList",
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
