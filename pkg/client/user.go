package client

import (
	"net/http"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	uuid "github.com/satori/go.uuid"
)

func (c *bepaClient) CreateUser(userName, email, password string) (*types.User, error) {
	userRequest := &types.UserReq{
		Name:     userName,
		Email:    email,
		Password: password,
	}

	createdUser := &types.User{}
	apiURL := trimURLSlash(routes.RouteUserCreate)
	if err := c.Do(http.MethodPost, apiURL, userRequest, createdUser); err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (c *bepaClient) GetSecret(userUUID *uuid.UUID) (*types.UserSecret, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserSecretGet), replaceDict)

	var secret types.UserSecret
	if err := c.Do(http.MethodGet, apiURL, nil, &secret); err != nil {
		return nil, err
	}
	return &secret, nil
}

func (c *bepaClient) RevokeSecret(userUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserSecretPost), replaceDict)

	return c.Do(http.MethodPost, apiURL, nil, nil)
}

func (c *bepaClient) CreateUserTokenByCreds(email, password string) (*types.UserToken, error) {
	tokenRequest := &types.UserTokenByCredsReq{
		Email:    email,
		Password: password,
	}

	createdToken := &types.UserToken{}
	apiURL := trimURLSlash(routes.RouteUserTokenCreateByCreds)
	if err := c.Do(http.MethodPost, apiURL, tokenRequest, createdToken); err != nil {
		return nil, err
	}
	return createdToken, nil
}

func (c *bepaClient) UpdateUser(userUUID *uuid.UUID, name, email, password string) error {
	userUpdateReq := &types.UserUpdateReq{
		Name:     name,
		Email:    email,
		Password: password,
	}

	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserUpdate), replaceDict)

	return c.Do(http.MethodPatch, apiURL, userUpdateReq, nil)
}

func (c *bepaClient) GetUserByName(userName string, workspaceUUID *uuid.UUID) (*types.User, error) {
	replaceDict := map[string]string{
		userEmailPlaceholder:     userName,
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceUserGetByEmail), replaceDict)

	user := &types.User{}
	if err := c.Do(http.MethodGet, apiURL, nil, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *bepaClient) GetMySelf() (*types.User, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOne), replaceDict)

	user := &types.User{}
	if err := c.Do(http.MethodGet, apiURL, nil, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *bepaClient) GetUser(userUUID *uuid.UUID) (*types.User, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOne), replaceDict)

	user := &types.User{}
	if err := c.Do(http.MethodGet, apiURL, nil, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *bepaClient) GetUsers() ([]*types.User, error) {
	users := []*types.User{}
	apiURL := trimURLSlash(routes.RouteUserGetAll)
	if err := c.Do(http.MethodGet, apiURL, nil, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *bepaClient) DeleteUser(userUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, nil, nil)
}

func (c *bepaClient) DeleteMySelf() error {
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserDelete), replaceDict)

	if err := c.Do(http.MethodDelete, apiURL, nil, nil); err != nil {
		return err
	}
	return nil
}

func (c *bepaClient) AddUserToWorkspace(userUUID, workspaceUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		userUUIDPlaceholder:      userUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserAppendWorkspace), replaceDict)
	return c.Do(http.MethodPost, apiURL, nil, nil)
}

func (c *bepaClient) RemoveUserFromWorkspace(userUUID, workspaceUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		userUUIDPlaceholder:      userUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserDropWorkspace), replaceDict)
	return c.Do(http.MethodDelete, apiURL, nil, nil)
}

func (c *bepaClient) SetMyPassword(password string) error {
	userUpdateReq := &types.UserUpdateReq{
		Password: password,
	}
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserUpdate), replaceDict)
	return c.Do(http.MethodPatch, apiURL, userUpdateReq, nil)
}

func (c *bepaClient) SetMyName(name string) error {
	userUpdateReq := &types.UserUpdateReq{
		Name: name,
	}
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserUpdate), replaceDict)
	return c.Do(http.MethodPatch, apiURL, userUpdateReq, nil)
}

func (c *bepaClient) SetMyEmail(email string) error {
	userUpdateReq := &types.UserUpdateReq{
		Email: email,
	}
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserUpdate), replaceDict)
	return c.Do(http.MethodPatch, apiURL, userUpdateReq, nil)
}

func (c *bepaClient) InviteUser(workspaceUUID *uuid.UUID, email string) (*types.InvitationInfo, error) {
	inviteReq := &types.InviteUserReq{
		Email: email,
	}
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	invitationInfo := &types.InvitationInfo{}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceInvite), replaceDict)
	err := c.Do(http.MethodPost, apiURL, inviteReq, invitationInfo)
	return invitationInfo, err
}

func (c *bepaClient) JoinByInvitationToken(name, password, invitationToken string) (*types.User, error) {
	joinReq := &types.UserAcceptInvitationReq{
		Name:     name,
		Password: password,
	}
	replaceDict := map[string]string{
		userInvitationTokenPlaceholder: invitationToken,
	}

	joinedUser := &types.User{}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserSetPassword), replaceDict)
	err := c.Do(http.MethodPost, apiURL, joinReq, joinedUser)
	return joinedUser, err
}

func (c *bepaClient) SuspendUser(userUUID *uuid.UUID) error {

	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserSuspend), replaceDict)

	return c.Do(http.MethodPut, apiURL, nil, nil)
}

func (c *bepaClient) ActivateUser(userUUID *uuid.UUID) error {

	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserActivate), replaceDict)

	return c.Do(http.MethodPut, apiURL, nil, nil)
}
