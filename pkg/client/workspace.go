package client

import (
	"errors"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/routes"
	"github.com/sotoon/iam-client/pkg/types"
)

func (c *iamClient) GetWorkspaces() ([]*types.Workspace, error) {
	apiURL := trimURLSlash(routes.RouteWorkspaceGetAll)

	workspaces := []*types.Workspace{}
	err := c.Do(http.MethodGet, apiURL, 0, nil, &workspaces)
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}

// todo deprecate and remove this functoin
func (c *iamClient) GetWorkspaceByName(name string) (*types.Workspace, error) {
	replaceDict := map[string]string{
		workspaceNamePlaceholder: name,
		userUUIDPlaceholder:      c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOneWorkspaceByName), replaceDict)

	workspace := &types.Workspace{}
	err := c.Do(http.MethodGet, apiURL, 0, nil, &workspace)
	if err != nil {
		return nil, err
	}
	return workspace, nil
}

func (c *iamClient) GetWorkspaceByNameAndOrgName(name string, organizationName string) (*types.WorkspaceWithOrganization, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOneWorkspace), replaceDict)

	parameters := map[string]string{
		"name":     name,
		"org_name": organizationName,
	}

	var workspacesSingleArray []types.WorkspaceWithOrganization

	err := c.DoWithParams(http.MethodGet, apiURL, parameters, 0, nil, &workspacesSingleArray)
	if err != nil {
		return nil, err
	}
	if len(workspacesSingleArray) == 1 {
		return &workspacesSingleArray[0], nil
	} else {
		return nil, errors.New("No workspace found")
	}
}

func (c *iamClient) GetWorkspace(uuid *uuid.UUID) (*types.Workspace, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetOne), replaceDict)

	workspace := &types.Workspace{}
	err := c.Do(http.MethodGet, apiURL, 0, nil, &workspace)
	if err != nil {
		return nil, err
	}
	return workspace, nil
}

func (c *iamClient) GetMyWorkspaces() ([]*types.WorkspaceWithOrganization, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetAllWorkspaces), replaceDict)

	workspaces := []*types.WorkspaceWithOrganization{}
	err := c.Do(http.MethodGet, apiURL, 0, nil, &workspaces)
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (c *iamClient) GetWorkspaceUsers(uuid *uuid.UUID) ([]*types.User, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetUsers), replaceDict)

	users := []*types.User{}
	err := c.Do(http.MethodGet, apiURL, 0, nil, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *iamClient) CreateWorkspace(name string) (*types.Workspace, error) {
	workspaceRequest := &types.WorkspaceReq{
		Name: name,
	}

	createdWorkspace := &types.Workspace{}
	apiURL := trimURLSlash(routes.RouteWorkspaceCreate)
	err := c.Do(http.MethodPost, apiURL, 0, workspaceRequest, &createdWorkspace)
	if err != nil {
		return nil, err
	}
	return createdWorkspace, nil
}

func (c *iamClient) GetWorkspaceRules(uuid *uuid.UUID) ([]*types.Rule, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetAllRules), replaceDict)
	rules := []*types.Rule{}
	err := c.Do(http.MethodGet, apiURL, 0, nil, &rules)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func (c *iamClient) GetWorkspaceRoles(uuid *uuid.UUID) ([]*types.Role, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetAllRoles), replaceDict)
	roles := []*types.Role{}
	err := c.Do(http.MethodGet, apiURL, 0, nil, &roles)
	if err != nil {
		return nil, err
	}
	return roles, nil

}

func (c *iamClient) DeleteWorkspace(uuid *uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceDelete), replaceDict)

	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}
