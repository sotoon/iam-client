package client

import (
	"fmt"
	"regexp"
	"testing"

	"git.platform.sotoon.ir/iam/golang-bepa-client/pkg/types"

	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
)

func TestDeleteDefaultBackupKey(t *testing.T) {
	backupKeyUUID := uuid.NewV4()
	conf := TestConfig{
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/.+/backup-key/(.+)/`),
		ClientMethodName: "DeleteDefaultWorkspaceBackupKey",
		Params:           []interface{}{&backupKeyUUID},
		ParamsInURL:      []interface{}{&backupKeyUUID},
	}
	DoTestDeleteAPI(t, conf)
}

func TestGetOneDefaultBackupKey(t *testing.T) {
	var object types.BackupKey
	backupKeyUUID := uuid.NewV4()
	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{&backupKeyUUID},
		ParamNames:       []string{"UUID"},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/.+/backup-key/(.+)/`),
		ClientMethodName: "GetOneDefaultBackupKey",
	}

	DoTestReadAPI(t, config)
}

func TestGetAllDefaultBackupKeys(t *testing.T) {
	var object []types.BackupKey
	config := TestConfig{
		Object:           &object,
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/.+/backup-key/`),
		ClientMethodName: "GetAllDefaultBackupKeys",
	}

	DoTestListingAPI(t, config)
}

func TestCreateBackupKeyForDefaultWorkspace(t *testing.T) {
	var object types.BackupKey
	var title, keyType, key string
	faker.FakeData(&title)
	faker.FakeData(&keyType)
	faker.FakeData(&key)

	config := TestConfig{
		Object:           &object,
		Params:           []interface{}{title, keyType, key},
		URLregexp:        regexp.MustCompile(`/api/v1/workspace/.+/backup-key/`),
		ClientMethodName: "CreateBackupKeyForDefaultWorkspace",
	}
	DoTestCreateAPI(t, config)
}

func TestCreateBackupKeyFromFileForDefaultWorkspace(t *testing.T) {
	fmt.Println("TestCreateBackupKeyFromFileForDefaultWorkspace not implemented.")
}
