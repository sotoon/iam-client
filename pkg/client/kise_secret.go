package client

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/routes"
	"github.com/sotoon/iam-client/pkg/types"
)

func (c *iamClient) DeleteUserKiseSecret(kiseSecretUUID *uuid.UUID) error {
	replaceDict := map[string]string{
		userUUIDPlaceholder:       c.userUUID,
		kiseSecretUUIDPlaceholder: kiseSecretUUID.String(),
		workspaceUUIDPlaceholder:  c.defaultWorkspace,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteKiseSecretDelete), replaceDict)

	err := c.Do(http.MethodDelete, apiURL, 0, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *iamClient) GetUserKiseSecrets(userUUID *uuid.UUID, workspaceUUID *uuid.UUID) ([]*types.KiseSecret, error) {
	userID := c.userUUID
	workspaceID := c.defaultWorkspace

	if userUUID != nil {
		userID = userUUID.String()
	}
	if workspaceUUID != nil {
		workspaceID = workspaceUUID.String()
	}

	replaceDict := map[string]string{
		userUUIDPlaceholder:      userID,
		workspaceUUIDPlaceholder: workspaceID,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteKiseSecretGetAll), replaceDict)

	kiseSecret := []*types.KiseSecret{}
	err := c.Do(http.MethodGet, apiURL, 0, nil, &kiseSecret)
	if err != nil {
		return nil, err
	}
	return kiseSecret, nil
}

func (c *iamClient) GetAllUserKiseSecret() ([]*types.KiseSecret, error) {
	return c.GetUserKiseSecrets(nil, nil)
}

func (c *iamClient) CreateUserKiseSecret(userUUID *uuid.UUID, workspaceUUID *uuid.UUID, title string) (*types.KiseSecret, error) {
	userID := c.userUUID
	workspaceID := c.defaultWorkspace

	if userUUID != nil {
		userID = userUUID.String()
	}
	if workspaceUUID != nil {
		workspaceID = workspaceUUID.String()
	}

	replaceDict := map[string]string{
		userUUIDPlaceholder:      userID,
		workspaceUUIDPlaceholder: workspaceID,
	}

	req := map[string]string{
		"title": title,
	}

	apiURL := substringReplace(trimURLSlash(routes.RouteKiseSecretCreate), replaceDict)
	createdKiseSecret := &types.KiseSecret{}
	if err := c.Do(http.MethodPost, apiURL, 0, req, createdKiseSecret); err != nil {
		return nil, err
	}
	return createdKiseSecret, nil
}

func (c *iamClient) CreateKiseSecretForDefaultUser() (*types.KiseSecret, error) {
	replaceDict := map[string]string{
		userUUIDPlaceholder:      c.userUUID,
		workspaceUUIDPlaceholder: c.defaultWorkspace,
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteKiseSecretCreate), replaceDict)

	createdKiseSecret := &types.KiseSecret{}
	if err := c.Do(http.MethodPost, apiURL, 0, nil, createdKiseSecret); err != nil {
		return nil, err
	}
	return createdKiseSecret, nil
}

func (c *iamClient) GetServiceUserKiseSecrets(workspaceUUID uuid.UUID) ([]*types.KiseSecret, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder: workspaceUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteKiseSecretServiceUserList), replaceDict)

	kiseSecrets := []*types.KiseSecret{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &kiseSecrets); err != nil {
		return nil, err
	}
	return kiseSecrets, nil
}

func (c *iamClient) CreateServiceUserKiseSecret(workspaceUUID, serviceUserUUID uuid.UUID, title string) (*types.KiseSecret, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
	}

	req := map[string]string{
		"title": title,
	}

	apiURL := substringReplace(trimURLSlash(routes.RouteKiseSecretServiceUserCreate), replaceDict)
	createdKiseSecret := &types.KiseSecret{}
	if err := c.Do(http.MethodPost, apiURL, 0, req, createdKiseSecret); err != nil {
		return nil, err
	}
	return createdKiseSecret, nil
}

func (c *iamClient) DeleteServiceUserKiseSecret(workspaceUUID, serviceUserUUID, kiseSecretUUID uuid.UUID) error {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
		kiseSecretUUIDPlaceholder:  kiseSecretUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteKiseSecretServiceUserDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, 0, nil, nil)
}
