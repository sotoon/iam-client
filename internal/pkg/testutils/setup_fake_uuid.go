package testutils

import (
	"reflect"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/client"
	"github.com/bxcodec/faker"
	uuid "github.com/satori/go.uuid"
)

func createTestURL(route string) string {
	return client.APIURI[:len(client.APIURI)-1] + route
}

func fakeUUID(v reflect.Value) (interface{}, error) {
	val := uuid.NewV4()
	return &val, nil
}

//Registers faker provider for uuid.UUID type
var _ error = func() error {
	faker.AddProvider("uuidObject", fakeUUID)
	return nil
}()
