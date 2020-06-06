package client

import (
	"fmt"
	"net/http"
	"time"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

type Workspace struct {
	UUID      *uuid.UUID `json:"uuid"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (c *bepaClient) GetWorkspaces() ([]*Workspace, error) {
	apiURL := trimURLSlash(routes.RouteWorkspaceGetAll)

	workspaces := []*Workspace{}
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

func (c *bepaClient) GetWorkspaceByName(name string) (*Workspace, error) {
	replaceDict := map[string]string{
		workspaceNamePlaceholder: name,
		userUUIDPlaceholder:      c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOneWorkspaceByName), replaceDict)

	workspace := &Workspace{}
	err := c.Do(http.MethodGet, apiURL, nil, &workspace)
	if err != nil {
		return nil, err
	}
	return workspace, nil
}

func (c *bepaClient) GetWorkspace(uuid *uuid.UUID) (*Workspace, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetOne), replaceDict)

	workspace := &Workspace{}
	err := c.Do(http.MethodGet, apiURL, nil, &workspace)
	if err != nil {
		return nil, err
	}
	return workspace, nil
}

func (c *bepaClient) GetMyWorkspaces() ([]*Workspace, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetAllWorkspaces), replaceDict)

	workspaces := []*Workspace{}
	err := c.Do(http.MethodGet, apiURL, nil, &workspaces)
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (c *bepaClient) GetWorkspaceUsers(uuid *uuid.UUID) ([]*User, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetUsers), replaceDict)

	users := []*User{}
	err := c.Do(http.MethodGet, apiURL, nil, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *bepaClient) CreateWorkspace(name string) (*Workspace, error) {
	workspaceRequest := &types.WorkspaceReq{
		Name: name,
	}

	createdWorkspace := &Workspace{}
	apiURL := trimURLSlash(routes.RouteWorkspaceCreate)
	err := c.Do(http.MethodPost, apiURL, workspaceRequest, &createdWorkspace)
	if err != nil {
		return nil, err
	}
	return createdWorkspace, nil
}

func (c *bepaClient) GetWorkspaceRules(uuid *uuid.UUID) ([]*Rule, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetAllRules), replaceDict)
	rules := []*Rule{}
	err := c.Do(http.MethodGet, apiURL, nil, &rules)
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func (c *bepaClient) GetWorkspaceRoles(uuid *uuid.UUID) ([]*Role, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: uuid.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetAllRoles), replaceDict)
	roles := []*Role{}
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
