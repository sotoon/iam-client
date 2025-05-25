package client

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/routes"
	"github.com/sotoon/iam-client/pkg/types"
)

func (c *iamClient) GetGroup(workspaceUUID, groupUUID *uuid.UUID) (*types.Group, error) {
	replaceDict := map[string]string{
		groupUUIDPlaceholder:     groupUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupGetOne), replaceDict)

	group := &types.Group{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, group); err != nil {
		return nil, err
	}
	return group, nil
}

func (c *iamClient) GetAllGroups(workspaceUUID *uuid.UUID) ([]*types.Group, error) {

	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	groups := []*types.Group{}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupGetALL), replaceDict)
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}

func (c *iamClient) DeleteGroup(workspaceUUID, groupUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		groupUUIDPlaceholder:     groupUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}

func (c *iamClient) GetGroupByName(workspaceName string, groupName string) (*types.Group, error) {
	replaceDict := map[string]string{
		groupNamePlaceholder:     groupName,
		workspaceNamePlaceholder: workspaceName,
		userUUIDPlaceholder:      c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupGetByName), replaceDict)

	group := &types.Group{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, group); err != nil {
		return nil, err
	}
	return group, nil
}
func (c *iamClient) CreateGroup(groupName string, workspace *uuid.UUID) (*types.GroupRes, error) {
	userRequest := &types.GroupReq{
		Name:      groupName,
		Workspace: workspace.String(),
	}
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspace.String(),
	}
	createdGroup := &types.GroupRes{}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupCreate), replaceDict)
	if err := c.Do(http.MethodPost, apiURL, 0, userRequest, createdGroup); err != nil {
		return nil, err
	}
	return createdGroup, nil
}
func (c *iamClient) GetGroupUser(workspaceUUID, groupUUID, userUUID *uuid.UUID) (*types.User, error) {
	replaceDict := map[string]string{
		groupUUIDPlaceholder:     groupUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		userUUIDPlaceholder:      userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupUserGetOne), replaceDict)

	user := &types.User{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *iamClient) GetAllGroupUsers(workspaceUUID, groupUUID *uuid.UUID) ([]*types.User, error) {

	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		groupUUIDPlaceholder:     groupUUID.String(),
	}
	users := []*types.User{}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupUserGetALL), replaceDict)
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *iamClient) GetAllGroupServiceUsers(workspaceUUID, groupUUID *uuid.UUID) ([]*types.ServiceUser, error) {

	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		groupUUIDPlaceholder:     groupUUID.String(),
	}
	serviceUsers := []*types.ServiceUser{}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupServiceUserGetALL), replaceDict)
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &serviceUsers); err != nil {
		return nil, err
	}
	return serviceUsers, nil
}

func (c *iamClient) UnbindUserFromGroup(workspaceUUID, groupUUID, userUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		groupUUIDPlaceholder:     groupUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		userUUIDPlaceholder:      userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupUnbindUser), replaceDict)
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}

func (c *iamClient) BindGroup(groupName string, workspace, groupUUID, userUUID *uuid.UUID) error {
	userRequest := &types.GroupReq{
		Name:      groupName,
		Workspace: workspace.String(),
	}
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspace.String(),
		groupUUIDPlaceholder:     groupUUID.String(),
		userUUIDPlaceholder:      userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupBindUser), replaceDict)
	if err := c.Do(http.MethodPost, apiURL, 0, userRequest, nil); err != nil {
		return err
	}
	return nil
}

func (c *iamClient) BindServiceUserToGroup(workspaceUUID, groupUUID, serviceUserUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
		groupUUIDPlaceholder:       groupUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupBindServiceUser), replaceDict)
	if err := c.Do(http.MethodPost, apiURL, 0, nil, nil); err != nil {
		return err
	}
	return nil
}

func (c *iamClient) UnbindServiceUserFromGroup(workspaceUUID, groupUUID, serviceUserUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
		groupUUIDPlaceholder:       groupUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupUnbindServiceUser), replaceDict)
	if err := c.Do(http.MethodDelete, apiURL, 0, nil, nil); err != nil {
		return err
	}
	return nil
}

func (c *iamClient) GetGroupServiceUser(workspaceUUID, groupUUID, serviceUserUUID *uuid.UUID) (*types.ServiceUser, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
		groupUUIDPlaceholder:       groupUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
	}
	serviceUser := &types.ServiceUser{}
	apiURL := substringReplace(trimURLSlash(routes.RouteGroupServiceUserGetOne), replaceDict)
	if err := c.Do(http.MethodGet, apiURL, 0, nil, serviceUser); err != nil {
		return nil, err
	}
	return serviceUser, nil
}
