package client

import (
	"net/http"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
)

func (c *bepaClient) Authorize(identity, action, object string) error {
	for i := 0; i < c.loadBalancer.TargetsLen()*2; i++ {
		req, target, err := c.NewRequest(http.MethodGet, trimURLSlash(routes.RouteAuthz), nil)

		if err != nil {
			return err
		}

		query := req.URL.Query()
		query.Set("identity", identity)
		query.Set("object", object)
		query.Set("action", action)

		req.URL.RawQuery = query.Encode()
		_, err = proccessRequest(req, target)

		if loopBreaker, err := checkErrorsAndPenaltyReward(err, target); err != nil && loopBreaker {
			return err
		}
		return nil
	}
	return errRetriesExceeded
}
