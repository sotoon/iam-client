package interceptor

import (
	"context"
	"net/http"
)

type BeforeRequestData struct {
	Request *http.Request
	Error   error
	Retry   bool
	Context context.Context
}

type AfterResponseData struct {
	Response *http.Response
	Error    error
	Retry    bool
	Context  context.Context
}

type ClientInterceptor interface {
	// BeforeRequest is called before executing an HTTP request
	// It can modify the request or perform pre-request actions
	// Returns: modified request, error, and a boolean indicating if retry is needed
	BeforeRequest(ctx context.Context, req *http.Request) BeforeRequestData

	// AfterResponse is called after receiving an HTTP response
	// It can modify the response or perform post-response actions
	// Returns: modified response, error, and a boolean indicating if retry is needed
	AfterResponse(ctx context.Context, resp *http.Response, successCode int) AfterResponseData
}
