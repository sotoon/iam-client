package client

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/routes"
	"github.com/sotoon/iam-client/pkg/types"
)

func (c *iamClient) GetServiceUser(workspaceUUID, serviceUserUUID *uuid.UUID) (*types.ServiceUser, error) {
	replaceDict := map[string]string{
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserGetOne), replaceDict)

	service := &types.ServiceUser{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, service); err != nil {
		return nil, err
	}
	return service, nil
}

func (c *iamClient) UpdateServiceUser(workspaceUUID, serviceUserUUID uuid.UUID, name, description string) (*types.ServiceUser, error) {
	updateReq := &types.ServiceUserUpdateReq{
		Name:        name,
		Description: description,
	}

	replaceDict := map[string]string{
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
	}

	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserUpdate), replaceDict)
	serviceUser := &types.ServiceUser{}
	if err := c.Do(http.MethodPut, apiURL, 0, updateReq, serviceUser); err != nil {
		return nil, err
	}
	return serviceUser, nil
}

func (c *iamClient) GetServiceUsers(workspaceUUID *uuid.UUID) ([]*types.ServiceUser, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	serviceUsers := []*types.ServiceUser{}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserGetALL), replaceDict)
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &serviceUsers); err != nil {
		return nil, err
	}
	return serviceUsers, nil
}

func (c *iamClient) DeleteServiceUser(workspaceUUID, serviceUserUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}

func (c *iamClient) GetServiceUserByName(workspaceName string, serviceUserName string) (*types.ServiceUser, error) {
	replaceDict := map[string]string{
		serviceUserNamePlaceholder: serviceUserName,
		workspaceNamePlaceholder:   workspaceName,
		userUUIDPlaceholder:        c.userUUID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserGetByName), replaceDict)

	serviceUser := &types.ServiceUser{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, serviceUser); err != nil {
		return nil, err
	}
	return serviceUser, nil
}

func (c *iamClient) CreateServiceUser(serviceUserName, description string, workspace *uuid.UUID) (*types.ServiceUser, error) {
	userRequest := &types.ServiceUserReq{
		Name:        serviceUserName,
		Workspace:   workspace.String(),
		Description: description,
	}
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspace.String(),
	}
	createdServiceUser := &types.ServiceUser{}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserCreate), replaceDict)
	if err := c.Do(http.MethodPost, apiURL, 0, userRequest, createdServiceUser); err != nil {
		return nil, err
	}
	return createdServiceUser, nil
}

func (c *iamClient) CreateServiceUserToken(serviceUserUUID, workspaceUUID *uuid.UUID, name string, expiresAt *time.Time) (*types.ServiceUserToken, error) {
	replaceDict := map[string]string{
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
	}
	ServiceUserToken := &types.ServiceUserToken{}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserTokenCreate), replaceDict)
	
	requestBody := map[string]interface{}{
		"name": name,
	}
	if expiresAt != nil {
		requestBody["expires_at"] = expiresAt.Format(time.RFC3339)
	}
	
	err := c.Do(http.MethodPost, apiURL, 0, requestBody, ServiceUserToken)
	return ServiceUserToken, err
}

func (c *iamClient) GetWorkspaceServiceUserTokenList(serviceUserUUID, workspaceUUID *uuid.UUID) (*[]types.ServiceUserToken, error) {

	replaceDict := map[string]string{
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
	}
	ServiceUserTokens := &[]types.ServiceUserToken{}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserTokenGetALL), replaceDict)
	err := c.Do(http.MethodGet, apiURL, 0, nil, ServiceUserTokens)
	return ServiceUserTokens, err
}

func (c *iamClient) GetWorkspaceServiceUserList(workspaceUUID uuid.UUID) ([]*types.ServiceUserWithCompactRole, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserDetailList), replaceDict)

	serviceUsers := []*types.ServiceUserWithCompactRole{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &serviceUsers); err != nil {
		return nil, err
	}
	return serviceUsers, nil
}

func (c *iamClient) GetWorkspaceServiceUserDetail(workspaceUUID, serviceUserUUID uuid.UUID) (*types.ServiceUserWithCompactRole, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserDetailGetOne), replaceDict)

	serviceUser := &types.ServiceUserWithCompactRole{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, serviceUser); err != nil {
		return nil, err
	}
	return serviceUser, nil
}

func (c *iamClient) GetWorkspaceServiceUserPublicKeyList(workspaceUUID, serviceUserUUID uuid.UUID) ([]*types.ServiceUserPublicKey, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserPublicKeyList), replaceDict)

	publicKeys := []*types.ServiceUserPublicKey{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &publicKeys); err != nil {
		return nil, err
	}
	return publicKeys, nil
}

func (c *iamClient) CreateServiceUserPublicKey(workspaceUUID, serviceUserUUID uuid.UUID, name, publicKey string) (*types.ServiceUserPublicKey, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
	}

	req := map[string]string{
		"title": name,
		"key":   publicKey,
	}

	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserPublicKeyCreate), replaceDict)
	createdPublicKey := &types.ServiceUserPublicKey{}
	if err := c.Do(http.MethodPost, apiURL, 0, req, createdPublicKey); err != nil {
		return nil, err
	}
	return createdPublicKey, nil
}

func (c *iamClient) DeleteServiceUserPublicKey(workspaceUUID, serviceUserUUID, publicKeyUUID uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
		publicKeyUUIDPlaceholder:   publicKeyUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserPublicKeyDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}

func (c *iamClient) DeleteServiceUserToken(serviceUserUUID, workspaceUUID, serviceUserTokenUUID *uuid.UUID) error {

	replaceDict := map[string]string{
		serviceUserUUIDPlaceholder:      serviceUserUUID.String(),
		workspaceUUIDPlaceholder:        workspaceUUID.String(),
		serviceUserTokenUUIDPlaceholder: serviceUserTokenUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceUserTokenDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)

}
