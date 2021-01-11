package testutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"testing"

	"git.cafebazaar.ir/infrastructure/bepa-client/pkg/client"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/require"
)

var uuidregex = `\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b`

const (
	//TestAccessToken is a dummy token for tests
	TestAccessToken = "test_access_token"
	//TestWorkspace is a dummy workspace name for tests
	TestWorkspace = "test_workspace"
	//TestUserUUID is dummy text represents a sample uuid for tests
	TestUserUUID = "00000000-00000000-00000000-00000000"
)

//NewTestClient creates a testing client
func NewTestClient(s *httptest.Server) client.Client {
	targets := []string{s.URL}
	c, _ := client.NewClient(TestAccessToken, targets, TestWorkspace, TestUserUUID)
	return c
}

//WriteObject serializes object to json and writes it as http response body
func WriteObject(w *http.ResponseWriter, object interface{}) {
	json.NewEncoder(*w).Encode(object)
}

//ReadObject de-serializes object to json and writes it as http response body
func ReadObject(r *http.Request) (interface{}, error) {
	var object interface{}
	err := json.NewDecoder(r.Body).Decode(&object)
	return object, err
}

func isWriteMethod(method string) bool {
	writeMethods := []string{http.MethodPatch, http.MethodPost, http.MethodPut}
	for _, val := range writeMethods {
		if val == method {
			return true
		}
	}
	return false
}

//TestHandlerFunc callback type for custom checks
type TestHandlerFunc func(*http.ResponseWriter, *http.Request, *regexp.Regexp) bool

func checkRegex(t *testing.T, path string, regex *regexp.Regexp) {
	require.True(t,
		regex.MatchString(path),
		fmt.Sprintf("url regex %v not matched with %s", regex, path),
	)
}

func checkParamsInURL(t *testing.T, params []interface{}, path string, regex *regexp.Regexp) {
	matches := regex.FindStringSubmatch(path)[1:]
	for i, v := range matches {
		require.Equal(t, v, fmt.Sprint(params[i]))
	}
}

func createClientParams(params []interface{}) []reflect.Value {
	if len(params) == 0 {
		return nil
	}
	reflectedParams := make([]reflect.Value, len(params))
	for i, v := range params {
		reflectedParams[i] = reflect.ValueOf(v)
	}
	return reflectedParams
}

func getReturnedObject(v reflect.Value) interface{} {

	if v.Kind() == reflect.Ptr {
		return v.Elem().Interface()
	}
	return v.Interface()
}

func toString(value interface{}) string {
	return fmt.Sprintf("%s", value)
}
func getReturnedErr(v reflect.Value) error {
	if v.IsNil() {
		return nil
	}
	return v.Interface().(error)
}

func toReflectedValue(val interface{}) reflect.Value {
	reflected := reflect.ValueOf(val)
	if reflected.Kind() == reflect.Ptr {
		reflected = reflected.Elem()
	}
	return reflected
}

func dereferenced(val reflect.Value) reflect.Value {
	reflected := val
	if reflected.Kind() == reflect.Ptr {
		reflected = reflected.Elem()
	}
	return reflected
}

type httpTestCase struct {
	statusCode int
	err        error
}

func createTestServer(t *testing.T, config *TestConfig, tc *httpTestCase, method string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, r.Method, method)
		if isWriteMethod(r.Method) {
			reflectedObject := toReflectedValue(config.Object)
			for i, paramName := range config.ParamNames {
				if paramName != "" {

					val := toReflectedValue(config.Params[i])
					field := reflectedObject.FieldByName(paramName)
					require.Equal(t, val.Kind(), field.Kind())
					field.Set(val)
				}
			}
		}
		if config.URLregexp != nil {

			checkRegex(t, r.URL.Path, config.URLregexp)
			if config.CustomHandlerTest != nil {
				if config.CustomHandlerTest(&w, r, config.URLregexp) {
					return
				}
			}
		}

		if config.ParamsInURL != nil {
			checkParamsInURL(t, config.ParamsInURL, r.URL.Path, config.URLregexp)
		}

		w.WriteHeader(tc.statusCode)
		if tc.statusCode == 201 || tc.statusCode == 200 {
			WriteObject(&w, config.Object)
		}
	}))

}

func matchFields(t *testing.T, config *TestConfig, value reflect.Value) {
	value = dereferenced(value)
	reflectedObject := toReflectedValue(config.Object)
	for _, paramName := range config.ParamNames {
		require.Equal(t,
			value.FieldByName(paramName).String(),
			reflectedObject.FieldByName(paramName).String(),
			fmt.Sprintf("values is diffrent %v!=%v", value, reflectedObject),
		)
	}
}

func callMethod(client *client.Client, name string, params []interface{}) ([]reflect.Value, error) {
	paramsReflected := createClientParams(params)
	method := reflect.ValueOf(*client).MethodByName(name)
	values := method.Call(paramsReflected)

	valuesLen := len(values)
	if valuesLen == 0 {
		return nil, nil
	}
	err := getReturnedErr(values[valuesLen-1])
	return values[:valuesLen], err
}

func getType(v reflect.Value) string {
	for t := v.Type(); ; {
		switch t.Kind() {
		case reflect.Ptr, reflect.Slice:
			t = t.Elem()
		default:
			return t.String()
		}
	}
}

//TestConfig is a configuration declaration for listing api tests
type TestConfig struct {
	Object            interface{}
	ClientMethodName  string
	CustomHandlerTest TestHandlerFunc
	Params            []interface{}
	ParamsInURL       []interface{}
	ParamNames        []string
	URLregexp         *regexp.Regexp
}

//DoTestListingAPI does tests for a listing api
func DoTestListingAPI(t *testing.T, config TestConfig) {
	kind := reflect.TypeOf(config.Object).Elem().Kind()
	require.Equal(t, kind, reflect.Slice, fmt.Sprintf("object should be a slice, but %v is given.", kind))
	faker.FakeData(config.Object)
	testCases := []httpTestCase{
		{200, nil},
		{403, client.ErrForbidden},
	}
	for _, tc := range testCases {
		s := createTestServer(t, &config, &tc, http.MethodGet)

		c := NewTestClient(s)
		values, err := callMethod(&c, config.ClientMethodName, config.Params)

		if tc.statusCode == 200 {
			returnedObjects := getReturnedObject(values[0])
			reflectedObjects := reflect.ValueOf(config.Object)
			require.Equal(t, getType(values[0]), getType(reflectedObjects))
			require.NoError(t, err)
			require.Len(t, returnedObjects, reflectedObjects.Elem().Len())
		} else {
			require.Error(t, err)
			require.EqualError(t, err, tc.err.Error())
		}

	}
}

//DoTestReadAPI does tests for a read api
func DoTestReadAPI(t *testing.T, config TestConfig) {
	faker.FakeData(config.Object)
	testCases := []httpTestCase{
		{200, nil},
		{403, client.ErrForbidden},
		{404, client.ErrNotFound},
	}

	for _, tc := range testCases {
		s := createTestServer(t, &config, &tc, http.MethodGet)

		c := NewTestClient(s)
		values, err := callMethod(&c, config.ClientMethodName, config.Params)
		if tc.statusCode == 200 {
			require.NoError(t, err)
			value := values[0]
			require.Equal(t, getType(value), getType(reflect.ValueOf(config.Object)))

			matchFields(t, &config, value)

		} else {
			require.Error(t, err)
			require.EqualError(t, err, tc.err.Error())
		}

	}
}

//DoTestDeleteAPI tests
func DoTestDeleteAPI(t *testing.T, config TestConfig) {
	testCases := []httpTestCase{
		{204, nil},
		{403, client.ErrForbidden},
		{404, client.ErrNotFound},
	}
	for _, tc := range testCases {
		s := createTestServer(t, &config, &tc, http.MethodDelete)

		c := NewTestClient(s)
		_, err := callMethod(&c, config.ClientMethodName, config.Params)
		if tc.statusCode == 204 {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
			require.EqualError(t, err, tc.err.Error())
		}

	}

}

//DoTestCreateAPI test
func DoTestCreateAPI(t *testing.T, config TestConfig) {
	faker.FakeData(config.Object)

	testCases := []httpTestCase{
		{201, nil},
		{403, client.ErrForbidden},
		{400, client.ErrBadRequest},
	}
	for _, tc := range testCases {
		s := createTestServer(t, &config, &tc, http.MethodPost)

		c := NewTestClient(s)
		values, err := callMethod(&c, config.ClientMethodName, config.Params)

		if tc.statusCode == 201 {
			value := values[0]
			require.Equal(t, getType(value), getType(reflect.ValueOf(config.Object)))
			matchFields(t, &config, value)
			require.NoError(t, err)
		} else {
			require.Error(t, err)
			require.EqualError(t, err, tc.err.Error())
		}
	}
}

func DoTestUpdateAPI(t *testing.T, config TestConfig, httpMethod string) {
	if config.Object != nil {
		faker.FakeData(config.Object)
	}
	// reflectedObject := reflect.ValueOf(config.Object)
	// for i, paramName := range config.ParamNames {
	// 	if paramName != "" {
	// 		fmt.Println(paramName, reflectedObject)

	// 		val := reflect.ValueOf(config.Params[i])
	// 		fmt.Println(val)

	// 		reflectedObject.Elem().FieldByName(paramName).Set(val)
	// 	}
	// }

	testCases := []httpTestCase{
		{200, nil},
		{403, client.ErrForbidden},
		{400, client.ErrBadRequest},
	}
	for _, tc := range testCases {
		s := createTestServer(t, &config, &tc, httpMethod)

		c := NewTestClient(s)
		_, err := callMethod(&c, config.ClientMethodName, config.Params)
		if tc.statusCode == 200 {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
			require.EqualError(t, err, tc.err.Error())
		}

	}
}
