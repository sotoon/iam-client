package client

import (
	"fmt"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"net/http"
)


func (c *bepaClient) GetWorkspaces() ([]*types.Workspace, error) {
	apiURL := trimURLSlash(routes.RouteWorkspaceGetAll)

	workspaces := []*types.Workspace{}
	err := c.Do(http.MethodGet, apiURL, nil, &workspaces)
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (c *bepaClient) SetConfigDefaultWorkspace(uuid *uuid.UUID) error {
	context := viper.GetString("current-context")
	viper.Set(fmt.Sprintf("contexts.%s.workspace", context), uuid.String())
	c.defaultWorkspace = uuid.String()
	return persistClientConfigFile()
}

func (c *bepaClient) GetWorkspaceByName(name string) (*types.Workspace, error) {
	replaceDict := map[string]string{
		workspaceNamePlaceholder: name,
		userUUIDPlaceholder:      c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOneWorkspaceByName), replaceDict)

	workspace := &types.Workspace{}
	err := c.Do(http.MethodGet, apiURL, nil, &workspace)
	if err != nil {
		return nil, err
	}
	return workspace, nil
}

func (c *bepaClient) GetWorkspace(uuid *uuid.UUID) (*types.Workspace, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetOne), replaceDict)

	workspace := &types.Workspace{}
	err := c.Do(http.MethodGet, apiURL, nil, &workspace)
	if err != nil {
		return nil, err
	}
	return workspace, nil
}

func (c *bepaClient) GetMyWorkspaces() ([]*types.Workspace, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetAllWorkspaces), replaceDict)

	workspaces := []*types.Workspace{}
	err := c.Do(http.MethodGet, apiURL, nil, &workspaces)
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (c *bepaClient) GetWorkspaceUsers(uuid *uuid.UUID) ([]*types.User, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetUsers), replaceDict)

	users := []*types.User{}
	err := c.Do(http.MethodGet, apiURL, nil, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *bepaClient) CreateWorkspace(name string) (*types.Workspace, error) {
	workspaceRequest := &types.WorkspaceReq{
		Name: name,
	}

	createdWorkspace := &types.Workspace{}
	apiURL := trimURLSlash(routes.RouteWorkspaceCreate)
	err := c.Do(http.MethodPost, apiURL, workspaceRequest, &createdWorkspace)
	if err != nil {
		return nil, err
	}
	return createdWorkspace, nil
}

func (c *bepaClient) GetWorkspaceRules(uuid *uuid.UUID) ([]*types.Rule, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetAllRules), replaceDict)
	rules := []*types.Rule{}
	err := c.Do(http.MethodGet, apiURL, nil, &rules)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func (c *bepaClient) GetWorkspaceRoles(uuid *uuid.UUID) ([]*types.Role, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetAllRoles), replaceDict)
	roles := []*types.Role{}
	err := c.Do(http.MethodGet, apiURL, nil, &roles)
	if err != nil {
		return nil, err
	}
	return roles, nil

}

func (c *bepaClient) DeleteWorkspace(uuid *uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceDelete), replaceDict)

	return c.Do(http.MethodPost, apiURL, nil, nil)
}
