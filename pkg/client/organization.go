package client

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/routes"
	"github.com/sotoon/iam-client/pkg/types"
)

func (c *iamClient) GetOrganizations() ([]*types.Organization, error) {
	apiURL := trimURLSlash(routes.RouteOrganizationGetAll)
	organizations := []*types.Organization{}
	err := c.Do(http.MethodGet, apiURL, 0, nil, &organizations)
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func (c *iamClient) GetOrganization(organizationUUID *uuid.UUID) (*types.Organization, error) {
	replaceDict := map[string]string{
		organizationUUIDPlaceholder: organizationUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteOrganizationGetOne), replaceDict)
	var organization *types.Organization
	err := c.Do(http.MethodGet, apiURL, 0, nil, &organization)
	if err != nil {
		return nil, err
	}
	return organization, nil
}

func (c *iamClient) GetOrganizationWorkspaces(organizationUUID *uuid.UUID) ([]*types.Workspace, error) {
	replaceDict := map[string]string{
		organizationUUIDPlaceholder: organizationUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteOrganizationWorkspacesGetAll), replaceDict)
	workspaces := []*types.Workspace{}
	err := c.Do(http.MethodGet, apiURL, 0, nil, &workspaces)
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (c *iamClient) GetOrganizationWorkspace(organizationUUID, workspaceUUID *uuid.UUID) (*types.Workspace, error) {
	replaceDict := map[string]string{
		organizationUUIDPlaceholder: organizationUUID.String(),
		workspaceUUIDPlaceholder:    workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteOrganizationWorkspacesGetOne), replaceDict)
	var workspace *types.Workspace
	err := c.Do(http.MethodGet, apiURL, 0, nil, &workspace)
	if err != nil {
		return nil, err
	}
	return workspace, nil
}
