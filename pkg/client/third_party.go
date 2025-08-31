package client

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sotoon/iam-client/pkg/routes"
	"github.com/sotoon/iam-client/pkg/types"
)

// GetThirdPartyBulkRefreshTokens retrieves all bulk refresh tokens for a service user in a workspace for a specific third party
func (c *iamClient) GetThirdPartyBulkRefreshTokens(workspaceUUID, thirdPartyUUID, serviceUserUUID uuid.UUID) ([]*types.ThirdPartyBulkRefreshToken, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
		thirdPartyUUIDPlaceholder:  thirdPartyUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteThirdPartyBulkRefreshToken), replaceDict)

	tokens := []*types.ThirdPartyBulkRefreshToken{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &tokens); err != nil {
		return nil, err
	}
	return tokens, nil
}

// CreateThirdPartyBulkRefreshToken creates a new bulk refresh token for a service user in a workspace for a specific third party
func (c *iamClient) CreateThirdPartyBulkRefreshToken(workspaceUUID, thirdPartyUUID, serviceUserUUID uuid.UUID, refreshToken string, expiresAt *time.Time) (*types.ThirdPartyBulkRefreshToken, error) {
	replaceDict := map[string]string{
		workspaceUUIDPlaceholder:   workspaceUUID.String(),
		thirdPartyUUIDPlaceholder:  thirdPartyUUID.String(),
		serviceUserUUIDPlaceholder: serviceUserUUID.String(),
	}

	req := map[string]interface{}{
		"refresh_token": refreshToken,
	}
	
	if expiresAt != nil {
		req["expires_at"] = expiresAt
	}

	apiURL := substringReplace(trimURLSlash(routes.RouteThirdPartyBulkRefreshToken), replaceDict)
	token := &types.ThirdPartyBulkRefreshToken{}
	if err := c.Do(http.MethodPost, apiURL, 0, req, token); err != nil {
		return nil, err
	}
	return token, nil
}

// GetThirdPartyAccessTokens retrieves all access tokens for a third party in an organization
func (c *iamClient) GetThirdPartyAccessTokens(organizationUUID, thirdPartyUUID uuid.UUID) ([]*types.ThirdPartyAccessToken, error) {
	replaceDict := map[string]string{
		organizationUUIDPlaceholder: organizationUUID.String(),
		thirdPartyUUIDPlaceholder:   thirdPartyUUID.String(),
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteThirdPartyAccessToken), replaceDict)

	tokens := []*types.ThirdPartyAccessToken{}
	if err := c.Do(http.MethodGet, apiURL, 0, nil, &tokens); err != nil {
		return nil, err
	}
	return tokens, nil
}

// CreateThirdPartyAccessToken creates a new access token for a third party in an organization
func (c *iamClient) CreateThirdPartyAccessToken(organizationUUID, thirdPartyUUID uuid.UUID, accessToken string, expiresAt *time.Time) (*types.ThirdPartyAccessToken, error) {
	replaceDict := map[string]string{
		organizationUUIDPlaceholder: organizationUUID.String(),
		thirdPartyUUIDPlaceholder:   thirdPartyUUID.String(),
	}

	req := map[string]interface{}{
		"access_token": accessToken,
	}
	
	if expiresAt != nil {
		req["expires_at"] = expiresAt
	}

	apiURL := substringReplace(trimURLSlash(routes.RouteThirdPartyAccessToken), replaceDict)
	token := &types.ThirdPartyAccessToken{}
	if err := c.Do(http.MethodPost, apiURL, 0, req, token); err != nil {
		return nil, err
	}
	return token, nil
}
