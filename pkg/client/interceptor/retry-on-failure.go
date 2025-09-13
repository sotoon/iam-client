package interceptor

import (
	"context"
	"net/http"

	"github.com/sotoon/iam-client/pkg/models"
)

type RetryInterceptor struct {
}

func NewRetryInterceptor() *RetryInterceptor {
	return &RetryInterceptor{}
}

// BeforeRequest does nothing for retry interceptor
func (r *RetryInterceptor) BeforeRequest(ctx context.Context, req *http.Request) BeforeRequestData {
	return BeforeRequestData{
		Request: req,
		Retry:   false,
		Context: ctx,
	}
}

// AfterResponse checks if the response should be retried
func (r *RetryInterceptor) AfterResponse(ctx context.Context, resp *http.Response, successCode int) AfterResponseData {
	if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
		return AfterResponseData{
			Error:   models.ErrTooManyRequests,
			Retry:   true,
			Context: ctx,
		}
	}

	return AfterResponseData{
		Response: resp,
		Retry:    false,
		Context:  ctx,
	}
}
