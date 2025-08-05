package client

import (
	"net/http"

	"github.com/sotoon/iam-client/pkg/routes"
	"github.com/sotoon/iam-client/pkg/types"

	uuid "github.com/satori/go.uuid"
)

func (c *iamClient) CreateRole(roleName string, workspaceUUID *uuid.UUID) (*types.Role, error) {
	roleRequest := &types.RoleReq{
		Name:      roleName,
		Workspace: workspaceUUID.String(),
	}

	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleCreate), replaceDict)

	createdRole := &types.Role{}
	if err := c.Do(http.MethodPost, apiURL, 0, roleRequest, createdRole); err != nil {
		return nil, err
	}
	return createdRole, nil
}

func (c *iamClient) UpdateRole(roleUUID *uuid.UUID, roleName string, workspaceUUID *uuid.UUID) (*types.Role, error) {
	roleRequest := &types.RoleReq{
		Name:      roleName,
		Workspace: workspaceUUID.String(),
	}

	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleUpdate), replaceDict)

	updatedRole := &types.Role{}
	if err := c.Do(http.MethodPatch, apiURL, 0, roleRequest, updatedRole); err != nil {
		return nil, err
	}
	return updatedRole, nil
}

func (c *iamClient) GetRoleByName(roleName, workspaceName string) (*types.RoleRes, error) {
	replaceDict := map[string]string{
		workspaceNamePlaceholder: workspaceName,
		roleNamePlaceholder:      roleName,
		userUUIDPlaceholder:      c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOneRoleByName), replaceDict)
	roleResponse := &types.RoleRes{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, roleResponse); err != nil {
		return nil, err
	}
	return roleResponse, nil
}

func (c *iamClient) GetRole(roleUUID, workspaceUUID *uuid.UUID) (*types.RoleRes, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleGetOne), replaceDict)

	role := &types.RoleRes{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, role); err != nil {
		return nil, err
	}
	return role, nil
}

func (c *iamClient) GetRoleUsers(roleUUID, workspaceUUID *uuid.UUID) ([]*types.User, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleGetAllUsers), replaceDict)

	users := []*types.User{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *iamClient) GetRoleRules(roleUUID, workspaceUUID *uuid.UUID) ([]*types.Rule, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleGetAllRules), replaceDict)

	rules := []*types.Rule{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &rules); err != nil {
		return nil, err
	}
	return rules, nil
}

func (c *iamClient) GetUserRoles(userUUID *uuid.UUID) ([]*types.RoleBinding, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetAllRoles), replaceDict)

	roles := []*types.RoleBinding{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (c *iamClient) DeleteRole(roleUUID, workspaceUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}

func (c *iamClient) GetAllRoles() ([]*types.Role, error) {
	replaceDict := map[string]string{}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleGetAll), replaceDict)

	roles := []*types.Role{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (c *iamClient) GetBindedRoleToUserItems(workspaceUUID, roleUUID, userUUID *uuid.UUID) (map[string]string, error) {
	replaceDict := map[string]string{
		roleUUIDPlaceholder:      roleUUID.String(),
		userUUIDPlaceholder:      userUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetBindedRole), replaceDict)

	roleBindingRes := &types.RoleBindingRes{}
	if err := c.Do(http.MethodGet, apiURL, 200, nil, roleBindingRes); err != nil {
		return nil, err
	}
	if len(roleBindingRes.Items) == 0 {
		return map[string]string{}, nil
	}
	return roleBindingRes.Items[0], nil
}

func (c *iamClient) GetBindedRoleToGroupItems(workspaceUUID, roleUUID, groupUUID *uuid.UUID) (map[string]string, error) {
	replaceDict := map[string]string{
		roleUUIDPlaceholder:      roleUUID.String(),
		groupUUIDPlaceholder:     groupUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupGetBindedRole), replaceDict)

	roleBindingRes := &types.RoleBindingRes{}
	if err := c.Do(http.MethodGet, apiURL, 200, nil, roleBindingRes); err != nil {
		return nil, err
	}
	if len(roleBindingRes.Items) == 0 {
		return map[string]string{}, nil
	}
	return roleBindingRes.Items[0], nil
}

func (c *iamClient) GetBindedRoleToServiceUserItems(workspaceUUID, roleUUID, serviceUserUUID *uuid.UUID) (map[string]string, error) {
	replaceDict := map[string]string{
		roleUUIDPlaceholder:        roleUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserGetBindedRole), replaceDict)

	roleBindingRes := &types.RoleBindingRes{}
	if err := c.Do(http.MethodGet, apiURL, 200, nil, roleBindingRes); err != nil {
		return nil, err
	}
	if len(roleBindingRes.Items) == 0 {
		return map[string]string{}, nil
	}
	return roleBindingRes.Items[0], nil
}

func (c *iamClient) BindRoleToUser(workspaceUUID, roleUUID, userUUID *uuid.UUID, items map[string]string) error {
	replaceDict := map[string]string{
		roleUUIDPlaceholder:      roleUUID.String(),
		userUUIDPlaceholder:      userUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	values := &types.RoleBindingReq{Items: items}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserAppendRole), replaceDict)
	return c.Do(http.MethodPost, apiURL, 0, values, nil)
}

func (c *iamClient) UnbindRoleFromUser(workspaceUUID, roleUUID, userUUID *uuid.UUID, items map[string]string) error {
	replaceDict := map[string]string{
		roleUUIDPlaceholder:      roleUUID.String(),
		userUUIDPlaceholder:      userUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserDropRole), replaceDict)
	if items != nil {
		apiURL = AddItemsAsQueryParams(apiURL, items)
	}
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}

func (c *iamClient) BindRoleToServiceUser(workspaceUUID, roleUUID, serviceUserUUID *uuid.UUID, items map[string]string) error {
	replaceDict := map[string]string{
		roleUUIDPlaceholder:        roleUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
	}
	values := &types.RoleBindingReq{Items: items}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserAppendRole), replaceDict)
	return c.Do(http.MethodPost, apiURL, 0, values, nil)
}

func (c *iamClient) UnbindRoleFromServiceUser(workspaceUUID, roleUUID, serviceUserUUID *uuid.UUID, items map[string]string) error {
	replaceDict := map[string]string{
		roleUUIDPlaceholder:        roleUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
	}

	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserDropRole), replaceDict)
	if items != nil {
		apiURL = AddItemsAsQueryParams(apiURL, items)
	}
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}

func (c *iamClient) GetRoleServiceUsers(roleUUID, workspaceUUID *uuid.UUID) ([]*types.ServiceUser, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleGetAllServiceUsers), replaceDict)

	serviceUsers := []*types.ServiceUser{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &serviceUsers); err != nil {
		return nil, err
	}
	return serviceUsers, nil
}

func (c *iamClient) BulkAddServiceUsersToRole(workspaceUUID, roleUUID uuid.UUID, serviceUserUUIDs []uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	
	serviceUserUUIDStrings := make([]string, len(serviceUserUUIDs))
	for i, uuid := range serviceUserUUIDs {
		serviceUserUUIDStrings[i] = uuid.String()
	}
	
	req := map[string][]string{
		"service_users": serviceUserUUIDStrings,
	}
	
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleBulkAddServiceUsers), replaceDict)
	return c.Do(http.MethodPost, apiURL, 0, req, nil)
}

func (c *iamClient) BulkAddUsersToRole(workspaceUUID, roleUUID uuid.UUID, userUUIDs []uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	
	userUUIDStrings := make([]string, len(userUUIDs))
	for i, uuid := range userUUIDs {
		userUUIDStrings[i] = uuid.String()
	}
	
	req := map[string][]string{
		"users": userUUIDStrings,
	}
	
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleBulkAddUsers), replaceDict)
	return c.Do(http.MethodPost, apiURL, 0, req, nil)
}

func (c *iamClient) BulkAddRulesToRole(workspaceUUID, roleUUID uuid.UUID, ruleUUIDs []uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	
	ruleUUIDStrings := make([]string, len(ruleUUIDs))
	for i, uuid := range ruleUUIDs {
		ruleUUIDStrings[i] = uuid.String()
	}
	
	req := map[string][]string{
		"rules": ruleUUIDStrings,
	}
	
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleBulkAddRules), replaceDict)
	return c.Do(http.MethodPost, apiURL, 0, req, nil)
}

func (c *iamClient) GetRoleGroups(roleUUID, workspaceUUID *uuid.UUID) ([]*types.Group, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		roleUUIDPlaceholder:      roleUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteRoleGetAllGroups), replaceDict)

	groups := []*types.Group{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}

func (c *iamClient) BindRoleToGroup(workspaceUUID, roleUUID, groupUUID *uuid.UUID, items map[string]string) error {
	replaceDict := map[string]string{
		roleUUIDPlaceholder:      roleUUID.String(),
		groupUUIDPlaceholder:     groupUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	values := &types.RoleBindingReq{Items: items}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupAppendRole), replaceDict)
	return c.Do(http.MethodPost, apiURL, 0, values, nil)
}

// Is it right?
func (c *iamClient) UnbindRoleFromGroup(workspaceUUID, roleUUID, groupUUID *uuid.UUID, items map[string]string) error {

	replaceDict := map[string]string{
		roleUUIDPlaceholder:      roleUUID.String(),
		groupUUIDPlaceholder:     groupUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}

	apiURL := substringReplace(trimURLSlash(routes.RouteGroupDropRole), replaceDict)
	if items != nil {
		apiURL = AddItemsAsQueryParams(apiURL, items)
	}
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}
