package client

import (
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	"net/http"
)

func (c *bepaClient) CreateTokenWithToken(server, secret string) (*UserToken, error) {
	userTokenreq := &types.UserTokenReq{
		Secret:     secret,
	}
	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	if err := c.SetServerURL(server); err != nil {
		return nil, err
	}
	userToken := &UserToken{}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserTokenCreateByToken), replaceDict)
	err := c.Do(http.MethodPost, apiURL, userTokenreq, userToken)
	return userToken, err
}

func (c *bepaClient) GetUserToken(server, user_token_uuid string) (*UserToken, error) {

	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
		userTokenUUIDPlaceholder: user_token_uuid,
	}
	if err := c.SetServerURL(server); err != nil {
		return nil, err
	}
	userToken := &UserToken{}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserTokenGetOne), replaceDict)
	err := c.Do(http.MethodGet, apiURL, nil, userToken)
	return userToken, err
}
func (c *bepaClient) GetAllUserToken(server string) (*[]UserToken, error) {

	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
	}
	if err := c.SetServerURL(server); err != nil {
		return nil, err
	}
	userTokens := &[]UserToken{}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserTokenGetAll), replaceDict)
	err := c.Do(http.MethodGet, apiURL, nil, userTokens)
	return userTokens, err
}

func (c *bepaClient) DeleteUserToken(server, user_token_uuid string) error {

	replaceDict := map[string]string{
		userUUIDPlaceholder: c.userUUID,
		userTokenUUIDPlaceholder: user_token_uuid,
	}
	if err := c.SetServerURL(server); err != nil {
		return err
	}
	apiURL := substringReplace(trimURLSlash(routes.RouteUserTokenDelete), replaceDict)
	return c.Do(http.MethodDelete, apiURL, nil, nil)

}
