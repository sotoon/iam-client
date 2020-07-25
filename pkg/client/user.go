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

type UserToken struct {
	UUID         string     `json:"uuid"`
	User         string     `json:"user"`
	Secret       string     `json:"secret"`
	Active       bool       `json:"active"`
	LastAccessAt *time.Time `json:"last_access_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type User struct {
	UUID            *uuid.UUID `json:"uuid"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	InvitationToken string     `json:"invitation_token,omitempty"`
}

type InvitationInfo struct {
	Token string `json:"invitation_token"`
}

type UserSecret struct {
	Secret string `json:"secret"`
}
type Service struct {
	Name string `json:"name"`
	Actions []string `json:"actions"`
}
func (c *bepaClient) CreateUser(userName, email, password string) (*User, error) {
	userRequest := &types.UserReq{
		Name:     userName,
		Email:    email,
		Password: password,
	}

	createdUser := &User{}
	apiURL := trimURLSlash(routes.RouteUserCreate)
	if err := c.Do(http.MethodPost, apiURL, userRequest, createdUser); err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (c *bepaClient) GetSecret(userUUID *uuid.UUID) (*UserSecret, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserSecretGet), replaceDict)

	var secret UserSecret
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

func (c *bepaClient) CreateUserTokenByCreds(email, password string) (*UserToken, error) {
	tokenRequest := &types.UserTokenByCredsReq{
		Email:    email,
		Password: password,
	}

	createdToken := &UserToken{}
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

func (c *bepaClient) GetUserByName(userName string, workspaceUUID *uuid.UUID) (*User, error) {
	replaceDict := map[string]string{
		userEmailPlaceholder:     userName,
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceUserGetByEmail), replaceDict)

	user := &User{}
	if err := c.Do(http.MethodGet, apiURL, nil, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *bepaClient) GetMySelf() (*User, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOne), replaceDict)

	user := &User{}
	if err := c.Do(http.MethodGet, apiURL, nil, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *bepaClient) GetUser(userUUID *uuid.UUID) (*User, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOne), replaceDict)

	user := &User{}
	if err := c.Do(http.MethodGet, apiURL, nil, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *bepaClient) GetUsers() ([]*User, error) {
	users := []*User{}
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

	user := &User{}
	if err := c.Do(http.MethodDelete, apiURL, nil, user); err != nil {
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

func (c *bepaClient) InviteUser(workspaceUUID *uuid.UUID, email string) (*InvitationInfo, error) {
	inviteReq := &types.InviteUserReq{
		Email: email,
	}
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	invitationInfo := &InvitationInfo{}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceInvite), replaceDict)
	err := c.Do(http.MethodPost, apiURL, inviteReq, invitationInfo)
	return invitationInfo, err
}

func (c *bepaClient) JoinByInvitationToken(server, name, password, invitationToken string) (*User, error) {
	joinReq := &types.UserAcceptInvitationReq{
		Name:     name,
		Password: password,
	}
	replaceDict := map[string]string{
		userInvitationTokenPlaceholder: invitationToken,
	}
	if err := c.SetServerURL(server); err != nil {
		return nil, err
	}
	joinedUser := &User{}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserSetPassword), replaceDict)
	err := c.Do(http.MethodPost, apiURL, joinReq, joinedUser)
	return joinedUser, err
}

func (c *bepaClient) SetConfigDefaultUserData(context, token, userUUID, email string) error {
	if context == "" {
		context = "default"
	}
	viper.Set(fmt.Sprintf("contexts.%s.token", context), token)
	viper.Set(fmt.Sprintf("contexts.%s.user-uuid", context), userUUID)
	viper.Set(fmt.Sprintf("contexts.%s.user", context), email)
	viper.Set(fmt.Sprintf("contexts.%s.addr", context), c.GetServerURL())
	c.accessToken = token
	c.userUUID = userUUID
	return persistClientConfigFile()
}

func (c *bepaClient) SetCurrentContext(context string) error {
	contexts := viper.GetStringMap("contexts")
	if _, ok := contexts[context]; ok {
		viper.Set("current-context", context)
		if err := persistClientConfigFile(); err == nil {
			fmt.Printf("set default context to %s\n", context)
			return nil
		}
	}
	return fmt.Errorf("could not find context %s", context)
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
