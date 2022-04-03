package client

import (
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	"regexp"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestDeleteUserKiseSecret(t *testing.T) {
	kiseSecretUUID := uuid.NewV4()
	conf := TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/.+/user/.+/kise/key/(.+)/`),
		ClientMethodName: "DeleteUserKiseSecret",
		Params:           []interface{}{&kiseSecretUUID},
		ParamsInURL:      []interface{}{&kiseSecretUUID},
	}
	DoTestDeleteAPI(t, conf)
}

func TestGetAllUserKiseSecret(t *testing.T) {
	var object []types.KiseSecret
	config := TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/.+/user/.+/kise/key/`),
		ClientMethodName: "GetAllUserKiseSecret",
	}

	DoTestListingAPI(t, config)
}

//func TestCreateKiseSecretForDefaultUser(t *testing.T) {
//	var object types.KiseSecret
//	config := TestConfig{
//		Object:           &object,
//		URLregexp:        regexp.MustCompile(`/api/v1/workspace/.+/user/.+/kise/key/`),
//		ClientMethodName: "CreateKiseSecretForDefaultUser",
//	}
//	DoTestCreateAPI(t, config)
//}
//
//
