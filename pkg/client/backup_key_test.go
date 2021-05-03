package client

import (
	"fmt"
	"regexp"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/internal/pkg/testutils"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"

	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
)

func TestDeleteDefaultBackupKey(t *testing.T) {
	backupKeyUUID := uuid.NewV4()
	conf := testutils.TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/.+/backup-key/(.+)/`),
		ClientMethodName: "DeleteDefaultWorkspaceBackupKey",
		Params:           []interface{}{&backupKeyUUID},
		ParamsInURL:      []interface{}{&backupKeyUUID},
	}
	testutils.DoTestDeleteAPI(t, conf)
}

func TestGetOneDefaultBackupKey(t *testing.T) {
	var object types.BackupKey
	backupKeyUUID := uuid.NewV4()
	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{&backupKeyUUID},
		ParamNames:       []string{"UUID"},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/.+/backup-key/(.+)/`),
		ClientMethodName: "GetOneDefaultBackupKey",
	}

	testutils.DoTestReadAPI(t, config)
}

func TestGetAllDefaultBackupKeys(t *testing.T) {
	var object []types.BackupKey
	config := testutils.TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/.+/backup-key/`),
		ClientMethodName: "GetAllDefaultBackupKeys",
	}

	testutils.DoTestListingAPI(t, config)
}

func TestCreateBackupKeyForDefaultWorkspace(t *testing.T) {
	var object types.BackupKey
	var title, keyType, key string
	faker.FakeData(&title)
	faker.FakeData(&keyType)
	faker.FakeData(&key)

	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{title, keyType, key},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/.+/backup-key/`),
		ClientMethodName: "CreateBackupKeyForDefaultWorkspace",
	}
	testutils.DoTestCreateAPI(t, config)
}

func TestCreateBackupKeyFromFileForDefaultWorkspace(t *testing.T) {
	fmt.Println("TestCreateBackupKeyFromFileForDefaultWorkspace not implemented.")
}

