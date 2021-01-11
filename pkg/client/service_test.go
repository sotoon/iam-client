package client

import (
	"regexp"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/internal/pkg/testutils"
	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/types"

	"github.com/bxcodec/faker"
)

func TestGetService(t *testing.T) {
	var object types.Service
	var serviceName string
	faker.FakeData(&serviceName)
	config := testutils.TestConfig{
		Object:           &object,
		Params:           []interface{}{serviceName},
		ParamNames:       []string{"Name"},
		ParamsInURL:      []interface{}{serviceName},
		URLregexp:        regexp.MustCompile(`^/api/v1/service/(.+)/$`),
		ClientMethodName: "GetService",
	}

	testutils.DoTestReadAPI(t, config)
}

func TestGetAllServices(t *testing.T) {

	services := []types.Service{}
	config := testutils.TestConfig{
		Object:          &services,
		URLregexp:        regexp.MustCompile(`/service/`),
		ClientMethodName: "GetAllServices",
	}
	testutils.DoTestListingAPI(t, config)

}
