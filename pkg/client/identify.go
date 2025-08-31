package client

import (
	"net/http"

	"github.com/sotoon/iam-client/pkg/routes"
	"github.com/sotoon/iam-client/pkg/types"
)

func (c *iamClient) Identify(token string) (*types.UserRes, error) {
	idenReq := &types.UserTokenReq{
		Secret: token,
	}

	userRes := &types.UserRes{}
	err := c.Do(http.MethodPost, trimURLSlash(routes.RouteUserTokenIdentify), 0, idenReq, userRes)
	if err != nil {
		return nil, err
	}

	return userRes, nil
}
