package client

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/internal/pkg/testutils"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"

	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
)

func TestDeleteDefaultUserPublicKey(t *testing.T) {
	publicKeyUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/user/.+/public-key/(.+)/`),
		ClientMethodName: "DeleteDefaultUserPublicKey",
		Params:           []interface{}{&publicKeyUUID},
		ParamsInURL:      []interface{}{&publicKeyUUID},
	}
	testutils.DoTestDeleteAPI(t, conf)
}

func TestGetOneDefaultUserPublicKey(t *testing.T) {
	var object types.PublicKey
	publicKeyUUID := uuid.NewV4()
	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&publicKeyUUID},
		ParamNames:       []string{"UUID"},
		URLregexp:        regexp.MustCompile(`/api/v1/user/.+/public-key/(.+)/`),
		ClientMethodName: "GetOneDefaultUserPublicKey",
	}

	testutils.DoTestReadAPI(t, config)
}

func TestGetAllDefaultUserPublicKeys(t *testing.T) {
	var object []types.PublicKey
	config := testutils.TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/api/v1/user/.+/public-key/`),
		ClientMethodName: "GetAllDefaultUserPublicKeys",
	}

	testutils.DoTestListingAPI(t, config)
}

func TestCreatePublicKeyForDefaultUser(t *testing.T) {
	var object types.PublicKey
	var title, keyType, key string
	faker.FakeData(&title)
	faker.FakeData(&keyType)
	faker.FakeData(&key)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{title, keyType, key},
		URLregexp:        regexp.MustCompile(`/api/v1/user/.+/public-key/`),
		ClientMethodName: "CreatePublicKeyForDefaultUser",
	}
	testutils.DoTestCreateAPI(t, config)
}

func TestCreatePublicKeyFromFileForDefaultUser(t *testing.T) {
	fmt.Println("TestCreatePublicKeyFromFileForDefaultUser not implemented.")
}

func TestVerifyPublicKey(t *testing.T) {
	var keyType, key, workspaceUUID, username, hostname string
	faker.FakeData(&keyType)
	faker.FakeData(&key)
	faker.FakeData(&workspaceUUID)
	faker.FakeData(&username)
	faker.FakeData(&hostname)
	config := testutils.TestConfig{
		Params:           []interface{}{keyType, key, workspaceUUID, username, hostname},
		URLregexp:        regexp.MustCompile(`/api/v1/public-key/verify/`),
		ClientMethodName: "VerifyPublicKey",
	}
	testutils.DoTestUpdateAPI(t, config, http.MethodPost)

}
