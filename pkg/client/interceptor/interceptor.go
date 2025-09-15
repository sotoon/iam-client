package interceptor

import (
	"context"
	"net/http"
)

type InterceptorData struct {
	ID             string // this ID will remain the same among all interceptors of an iteration
	InitialRequest *http.Request
	Request        *http.Request
	Response       *http.Response
	Context        context.Context
	Error          error
}

type ClientInterceptor interface {
	// BeforeRequest is called before executing an HTTP request
	// It can modify the request or perform pre-request actions
	//
	// NOTE: moidfiying InterceptorData.Response at this step will prevent calling AfterResponse
	//  and directly returns this response as response of server
	BeforeRequest(InterceptorData) InterceptorData

	// AfterResponse is called after receiving an HTTP response
	// It can modify the response or perform post-response actions
	AfterResponse(InterceptorData) InterceptorData
}
