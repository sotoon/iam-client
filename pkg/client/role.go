package client

import (
	"net/http"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"

	uuid "github.com/satori/go.uuid"
)

type Role struct {
	UUID      *uuid.UUID `json:"uuid"`
	Workspace *Workspace `json:"workspace"`
	Name      string     `json:"name"`
}

type RoleBinding struct {
	RoleName  string            `json:"name"`
	UserUUID  *uuid.UUID        `json:"user_uuid"`
	Workspace *Workspace        `json:"workspace"`
	Items     map[string]string `json:"items,omitempty"`
}

func (c *bepaClient) CreateRole(roleName string, workspaceUUID *uuid.UUID) (*Role, error) {
	roleRequest := &types.RoleReq{
		Name:      roleName,
		Workspace: workspaceUUID.String(),
	}

	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleCreate), replaceDict)

	createdRole := &Role{}
	if err := c.Do(http.MethodPost, apiURL, roleRequest, createdRole); err != nil {
		return nil, err
	}
	return createdRole, nil
}

func (c *bepaClient) GetRoleByName(roleName, workspaceName string) (*Role, error) {
	replaceDict := map[string]string{
		workspaceNamePlaceholder: workspaceName,
		roleNamePlaceholder:      roleName,
		userUUIDPlaceholder:      c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOneRoleByName), replaceDict)
	role := &Role{}
	if err := c.Do(http.MethodGet, apiURL, nil, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (c *bepaClient) GetRole(roleUUID, workspaceUUID *uuid.UUID) (*Role, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleGetOne), replaceDict)

	role := &Role{}
	if err := c.Do(http.MethodGet, apiURL, nil, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (c *bepaClient) GetRoleUsers(roleUUID, workspaceUUID *uuid.UUID) ([]*User, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleGetAllUsers), replaceDict)

	users := []*User{}
	if err := c.Do(http.MethodGet, apiURL, nil, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *bepaClient) GetRoleRules(roleUUID, workspaceUUID *uuid.UUID) ([]*Rule, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleGetAllRules), replaceDict)

	rules := []*Rule{}
	if err := c.Do(http.MethodGet, apiURL, nil, &rules); err != nil {
		return nil, err
	}
	return rules, nil
}

func (c *bepaClient) GetUserRoles(userUUID *uuid.UUID) ([]*RoleBinding, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetAllRoles), replaceDict)

	roles := []*RoleBinding{}
	if err := c.Do(http.MethodGet, apiURL, nil, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (c *bepaClient) DeleteRole(roleUUID, workspaceUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, nil, nil)
}

func (c *bepaClient) GetAllRoles() ([]*Role, error) {
	replaceDict := map[string]string{}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleGetAll), replaceDict)

	roles := []*Role{}
	if err := c.Do(http.MethodGet, apiURL, nil, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (c *bepaClient) BindRoleToUser(workspaceUUID, roleUUID, userUUID *uuid.UUID, items map[string]string) error {
	replaceDict := map[string]string{
		roleUUIDPlaceholder:      roleUUID.String(),
		userUUIDPlaceholder:      userUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	values := &types.RoleBindingReq{Items: items}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserAppendRole), replaceDict)
	return c.Do(http.MethodPost, apiURL, values, nil)
}

func (c *bepaClient) UnbindRoleFromUser(workspaceUUID, roleUUID, userUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		roleUUIDPlaceholder:      roleUUID.String(),
		userUUIDPlaceholder:      userUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserDropRole), replaceDict)
	return c.Do(http.MethodDelete, apiURL, nil, nil)
}
