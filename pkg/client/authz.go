package client

import (
	"net/http"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
)

func (c *bepaClient) Authorize(identity, userType, action, object string) error {

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
	_, err = proccessRequest(req)
	return err
}
