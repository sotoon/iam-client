package client

import (
	"net/http"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
)

func (c *bepaClient) Authorize(identity, userType, action, object string) error {
	c.log("authorizing %v", identity)
	req, err := c.NewRequest(http.MethodGet, trimURLSlash(routes.RouteAuthz), nil)

	if err != nil {
		return err
	}

	query := req.URL.Query()
	query.Set("identity", identity)
	query.Set("user_type", userType)
	query.Set("object", object)
	query.Set("action", action)

	req.URL.RawQuery = query.Encode()
	data, statusCode, errRes := proccessRequest(req, 0)
	if errRes == nil {
		c.log("user %v is authorized", identity)
		return nil
	}

	c.log("user %v is not authorized", identity)
	return &types.RequestExecutionError{
		Err:        errRes,
		StatusCode: statusCode,
		Data:       data,
	}
}

func (c *bepaClient) IdentifyAndAuthorize(token, action, object string) error {
	c.log("identifying and authorizing")

	req := &types.IdentifyAndAuthorizeReq{
		Token:  token,
		Action: action,
		Object: object,
	}
	var avoidedRes interface{}
	err := c.Do(http.MethodPost, trimURLSlash(routes.RouteIdentifyAndAuthorize), 0, req, avoidedRes)
	if err == nil {
		c.log("token is authorized")
		return nil
	}

	c.log("token is not authorized")
	return err
}
