package interceptor

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
	"github.com/sotoon/iam-client/pkg/models"
)

type CircuitBreakerInterceptor struct {
	cb *gobreaker.CircuitBreaker
}

// NewCircuitBreakerInterceptor creates a new circuit breaker interceptor
// if no cb provided, will make a default one.
func NewCircuitBreakerInterceptor(cb *gobreaker.CircuitBreaker) *CircuitBreakerInterceptor {

	if cb != nil {
		return &CircuitBreakerInterceptor{
			cb: cb,
		}
	}
	// Create Default circuit breaker

	cbSettings := gobreaker.Settings{
		Name:        "HTTP_REQUEST",
		MaxRequests: 0,                // unlimited concurrent requests
		Interval:    10 * time.Second, // check status every 10 seconds
		Timeout:     20 * time.Second, // how long to wait before closing circuit after it's opened
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 0 // instant circute openinig
		},
	}

	return &CircuitBreakerInterceptor{
		cb: gobreaker.NewCircuitBreaker(cbSettings),
	}
}

// BeforeRequest checks if the circuit breaker is open
func (c *CircuitBreakerInterceptor) BeforeRequest(ctx context.Context, req *http.Request) BeforeRequestData {
	if c.cb.State() == gobreaker.StateOpen {
		return BeforeRequestData{
			Error:   errors.New("circuit breaker is open"),
			Retry:   true,
			Context: ctx,
		}
	}

	return BeforeRequestData{
		Request: req,
		Retry:   false,
		Context: ctx,
	}
}

// AfterResponse records failures in the circuit breaker
func (c *CircuitBreakerInterceptor) AfterResponse(ctx context.Context, resp *http.Response, successCode int) AfterResponseData {
	if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
		c.cb.Execute(func() (interface{}, error) {
			return nil, models.ErrTooManyRequests
		})

		return AfterResponseData{
			Error:   models.ErrCircuitBreakerOpen,
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
