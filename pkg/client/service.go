package client

import (
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"
	"net/http"
)

func (c *bepaClient) GetService(server, name string) (*types.Service, error) {

	replaceDict := map[string]string{
		serviceNamePlaceholder: name,
	}
	if err := c.SetServerURL(server); err != nil {
		return nil, err
	}
	service := &types.Service{}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceGetOne), replaceDict)
	err := c.Do(http.MethodGet, apiURL, nil, service)
	return service, err
}
func (c *bepaClient) GetAllServices(server string) (*[]types.Service, error) {

	if err := c.SetServerURL(server); err != nil {
		return nil, err
	}
	services := &[]types.Service{}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceGetAll), nil)
	err := c.Do(http.MethodGet, apiURL, nil, services)
	return services, err
}