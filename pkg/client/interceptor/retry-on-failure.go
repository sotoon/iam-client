package interceptor

import (
	"net/http"

	"github.com/sotoon/iam-client/pkg/client"
)

// RetryInterceptor implements retry logic with exponential backoff
type RetryInterceptor struct {
}

// NewRetryInterceptor creates a new retry interceptor with default settings
func NewRetryInterceptor() *RetryInterceptor {
	return &RetryInterceptor{}
}

// BeforeRequest does nothing for retry interceptor
func (r *RetryInterceptor) BeforeRequest(req *http.Request) (*http.Request, error, bool) {
	// Nothing to do before the request
	return req, nil, false
}

// AfterResponse checks if the response should be retried
func (r *RetryInterceptor) AfterResponse(resp *http.Response, successCode int) (*http.Response, error, bool) {
	// Check for rate limit errors (429)
	if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
		// Return error with retry=true to indicate we should retry
		return resp, client.ErrTooManyRequests, true
	}

	// No retry needed
	return resp, nil, false
}
