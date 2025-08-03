package client

import (
	"encoding/json"
	"errors"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/routes"
	"github.com/sotoon/iam-client/pkg/types"
)

func (c *iamClient) CreateUser(userName, email, password string) (*types.User, error) {
	userRequest := &types.UserReq{
		Name:     userName,
		Email:    email,
		Password: password,
	}

	createdUser := &types.User{}
	apiURL := trimURLSlash(routes.RouteUserCreate)
	if err := c.Do(http.MethodPost, apiURL, 0, userRequest, createdUser); err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (c *iamClient) GetSecret(userUUID *uuid.UUID) (*types.UserSecret, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserSecretGet), replaceDict)

	var secret types.UserSecret
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &secret); err != nil {
		return nil, err
	}
	return &secret, nil
}

func (c *iamClient) RevokeSecret(userUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserSecretPost), replaceDict)

	return c.Do(http.MethodPost, apiURL, 0, nil, nil)
}

func (c *iamClient) CreateUserTokenByCreds(email, password string) (*types.UserToken, error) {
	tokenRequest := &types.UserTokenByCredsReq{
		Email:    email,
		Password: password,
	}

	createdToken := &types.UserToken{}
	apiURL := trimURLSlash(routes.RouteUserTokenCreateByCreds)
	if err := c.Do(http.MethodPost, apiURL, 201, tokenRequest, createdToken); err != nil {
		responseError := err.(*types.RequestExecutionError)
		if responseError.StatusCode == 200 {
			challengeRequired := &types.AuthnChallengeRequiredResponse{}
			decodeErr := json.Unmarshal(responseError.Data, &challengeRequired)
			if decodeErr != nil {
				return nil, decodeErr
			}

			return nil, challengeRequired
		}

		return nil, err
	}
	return createdToken, nil
}

func (c *iamClient) CreateUserTokenByChallenge(challengeToken, challengeAnswer string) (*types.UserToken, error) {
	tokenRequest := &types.AuthnChallengeRequest{
		ChallengeToken:  challengeToken,
		ChallengeAnswer: challengeAnswer,
	}

	createdToken := &types.UserToken{}
	apiURL := trimURLSlash(routes.RouteUserTokenByChallenge)
	if err := c.Do(http.MethodPost, apiURL, 201, tokenRequest, createdToken); err != nil {
		responseError := err.(*types.RequestExecutionError)
		if responseError.StatusCode == 200 {
			challengeRequired := &types.AuthnChallengeRequiredResponse{}
			decodeErr := json.Unmarshal(responseError.Data, &challengeRequired)
			if decodeErr != nil {
				return nil, decodeErr
			}

			return nil, challengeRequired
		}

		return nil, err
	}
	return createdToken, nil
}

func (c *iamClient) UpdateUser(userUUID *uuid.UUID, name, email, password string) error {
	userUpdateReq := &types.UserUpdateReq{
		Name:     name,
		Email:    email,
		Password: password,
	}

	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserUpdate), replaceDict)

	return c.Do(http.MethodPatch, apiURL, 0, userUpdateReq, nil)
}

func (c *iamClient) GetUserByEmail(email string, workspaceUUID *uuid.UUID) (*types.User, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	params := map[string]string{
		"email": email,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceGetUsers), replaceDict)

	var users []types.User
	if err := c.DoSimple(http.MethodGet, apiURL, params, nil, &users); err != nil {
		return nil, err
	}
	if len(users) > 0 {
		// because email is unique
		return &users[0], nil
	} else {
		return nil, errors.New("User not found")
	}
}

func (c *iamClient) GetUserByName(userName string, workspaceUUID *uuid.UUID) (*types.User, error) {
	replaceDict := map[string]string{
		userEmailPlaceholder:     userName,
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceUserGetByEmail), replaceDict)

	user := &types.User{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *iamClient) GetMySelf() (*types.User, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOne), replaceDict)

	user := &types.User{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *iamClient) GetUser(userUUID *uuid.UUID) (*types.User, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserGetOne), replaceDict)

	user := &types.User{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *iamClient) GetUsers() ([]*types.User, error) {
	users := []*types.User{}
	apiURL := trimURLSlash(routes.RouteUserGetAll)
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *iamClient) DeleteUser(userUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}

func (c *iamClient) DeleteMySelf() error {
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserDelete), replaceDict)

	if err := c.Do(http.MethodDelete, apiURL, 0, nil, nil); err != nil {
		return err
	}
	return nil
}

func (c *iamClient) AddUserToWorkspace(userUUID, workspaceUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		userUUIDPlaceholder:      userUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserAppendWorkspace), replaceDict)
	return c.Do(http.MethodPost, apiURL, 0, nil, nil)
}

func (c *iamClient) RemoveUserFromWorkspace(userUUID, workspaceUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		userUUIDPlaceholder:      userUUID.String(),
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserLeaveWorkspace), replaceDict)
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}

func (c *iamClient) SetMyPassword(password string) error {
	userUpdateReq := &types.UserUpdateReq{
		Password: password,
	}
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserUpdate), replaceDict)
	return c.Do(http.MethodPatch, apiURL, 0, userUpdateReq, nil)
}

func (c *iamClient) SetMyName(name string) error {
	userUpdateReq := &types.UserUpdateReq{
		Name: name,
	}
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserUpdate), replaceDict)
	return c.Do(http.MethodPatch, apiURL, 0, userUpdateReq, nil)
}

func (c *iamClient) SetMyEmail(email string) error {
	userUpdateReq := &types.UserUpdateReq{
		Email: email,
	}
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserUpdate), replaceDict)
	return c.Do(http.MethodPatch, apiURL, 0, userUpdateReq, nil)
}

func (c *iamClient) InviteUser(workspaceUUID *uuid.UUID, email string) (*types.InvitationInfo, error) {
	inviteReq := &types.InviteUserReq{
		Email: email,
	}
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	invitationInfo := &types.InvitationInfo{}
	apiURL := substringReplace(trimURLSlash(routes.RouteWorkspaceInvite), replaceDict)
	err := c.Do(http.MethodPost, apiURL, 0, inviteReq, invitationInfo)
	return invitationInfo, err
}

func (c *iamClient) JoinByInvitationToken(name, password, invitationToken string) (*types.User, error) {
	joinReq := &types.UserAcceptInvitationReq{
		Name:     name,
		Password: password,
	}
	replaceDict := map[string]string{
		userInvitationTokenPlaceholder: invitationToken,
	}

	joinedUser := &types.User{}
	apiURL := substringReplace(trimURLSlash(routes.RouteAcceptInvitation), replaceDict)
	err := c.Do(http.MethodPost, apiURL, 0, joinReq, joinedUser)
	return joinedUser, err
}

func (c *iamClient) SuspendUserInWorkspace(workspaceUUID *uuid.UUID, userUUID *uuid.UUID) error {

	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		userUUIDPlaceholder:      userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteSuspendUserInWorkspace), replaceDict)

	return c.DoMinimal(http.MethodPut, apiURL, nil)
}

func (c *iamClient) ActivateUserInWorkspace(workspaceUUID *uuid.UUID, userUUID *uuid.UUID) error {

	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		userUUIDPlaceholder:      userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteActivateUserInWorkspace), replaceDict)

	return c.DoMinimal(http.MethodPut, apiURL, nil)
}

func (c *iamClient) SuspendUser(userUUID *uuid.UUID) error {

	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserSuspend), replaceDict)

	return c.Do(http.MethodPut, apiURL, 0, nil, nil)
}

func (c *iamClient) ActivateUser(userUUID *uuid.UUID) error {

	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserActivate), replaceDict)

	return c.Do(http.MethodPut, apiURL, 0, nil, nil)
}

func (c *iamClient) ResetPassword(email string) error {
	resetRequest := &types.ResetPasswordReq{
		Email: email,
	}

	apiURL := trimURLSlash(routes.RouteUserResetPassword)
	return c.Do(http.MethodPost, apiURL, 0, resetRequest, nil)
}

func (c *iamClient) ChangePassword(token, password string) error {
	changeRequest := &types.ChangePasswordReq{
		Password: password,
	}

	replaceDict := map[string]string{
		userTokenUUIDPlaceholder: token,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserChangePassword), replaceDict)
	return c.Do(http.MethodPost, apiURL, 0, changeRequest, nil)
}

func (c *iamClient) GetUserDetailList(workspaceUUID uuid.UUID) ([]*types.User, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserDetailedList), replaceDict)

	users := []*types.User{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *iamClient) GetUserDetail(workspaceUUID, userUUID uuid.UUID) (*types.User, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
		userUUIDPlaceholder:      userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserDetailedDetail), replaceDict)

	user := &types.User{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (c *iamClient) GetUserOtp(userUUID uuid.UUID) (*types.UserOtp, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserOtpGet), replaceDict)

	otp := &types.UserOtp{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, otp); err != nil {
		return nil, err
	}
	return otp, nil
}

func (c *iamClient) CreateUserOtp(userUUID uuid.UUID) (*types.UserOtp, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserOtpPost), replaceDict)

	otp := &types.UserOtp{}
	if err := c.Do(http.MethodPost, apiURL, 0, nil, otp); err != nil {
		return nil, err
	}
	return otp, nil
}

func (c *iamClient) DeleteUserOtp(userUUID uuid.UUID) error {
	replaceDict := map[string]string{
		userUUIDPlaceholder: userUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserOtpDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}
