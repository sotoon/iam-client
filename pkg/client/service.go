package client

import (
	"net/http"

	"github.com/sotoon/iam-client/pkg/routes"
	"github.com/sotoon/iam-client/pkg/types"
)

func (c *iamClient) GetService(name string) (*types.Service, error) {

	replaceDict := map[string]string{
		serviceNamePlaceholder: name,
	}

	service := &types.Service{}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceGetOne), replaceDict)
	err := c.Do(http.MethodGet, apiURL, 0, nil, service)
	return service, err
}
func (c *iamClient) GetAllServices() (*[]types.Service, error) {

	services := &[]types.Service{}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceGetAll), nil)
	err := c.Do(http.MethodGet, apiURL, 0, nil, services)
	return services, err
}
