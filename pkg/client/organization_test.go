package client

import (
	"regexp"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	uuid "github.com/satori/go.uuid"
)

func TestGetAllOrganizations(t *testing.T) {
	var objects []types.Organization
	conf := TestConfig{
		ClientMethodName: "GetOrganizations",
		Object:           &objects,
		URLregexp:        regexp.MustCompile(`^/api/v1/organization/$`),
	}
	DoTestListingAPI(t, conf)
}

func TestGetOneOrganization(t *testing.T) {
	var object types.Organization
	organization := uuid.NewV4()
	conf := TestConfig{
		ClientMethodName: "GetOrganization",
		Params:           []interface{}{&organization},
		ParamNames:       []string{"organizationUUID"},
		Object:           &object,
		URLregexp:        regexp.MustCompile(`^/api/v1/organization/[^/]+/$`),
	}
	DoTestReadAPI(t, conf)
}

func TestGetOrganizationAllWorkspaces(t *testing.T) {
	var object []types.Workspace
	organization := uuid.NewV4()
	conf := TestConfig{
		ClientMethodName: "GetOrganizationWorkspaces",
		Params:           []interface{}{&organization},
		ParamNames:       []string{"organizationUUID"},
		Object:           &object,
		URLregexp:        regexp.MustCompile(`^/api/v1/organization/[^/]+/workspace/$`),
	}
	DoTestListingAPI(t, conf)
}

func TestGetOrganizationOneWorkspace(t *testing.T) {
	var object types.Workspace
	organization := uuid.NewV4()
	workspace := uuid.NewV4()
	conf := TestConfig{
		ClientMethodName: "GetOrganizationWorkspace",
		Params:           []interface{}{&organization, &workspace},
		ParamNames:       []string{"organizationUUID", "workspaceUUID"},
		Object:           &object,
		URLregexp:        regexp.MustCompile(`^/api/v1/organization/[^/]+/workspace/[^/]+/$`),
	}
	DoTestReadAPI(t, conf)
}
