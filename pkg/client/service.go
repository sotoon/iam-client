package client

import (
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/routes"
	"net/http"
)

func (c *bepaClient) GetService(server, name string) (*Service, error) {

	replaceDict := map[string]string{
		serviceNamePlaceholder: name,
	}
	if err := c.SetServerURL(server); err != nil {
		return nil, err
	}
	service := &Service{}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceGetOne), replaceDict)
	err := c.Do(http.MethodGet, apiURL, nil, service)
	return service, err
}
func (c *bepaClient) GetAllServices(server string) (*[]Service, error) {

	if err := c.SetServerURL(server); err != nil {
		return nil, err
	}
	services := &[]Service{}
	apiURL := substringReplace(trimURLSlash(routes.RouteServiceGetAll), nil)
	err := c.Do(http.MethodGet, apiURL, nil, services)
	return services, err
}