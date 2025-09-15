package interceptor

import (
	"fmt"
	"time"

	"github.com/sony/gobreaker"
	"github.com/sotoon/iam-client/pkg/models"
)

type CircuitBreakerInterceptor struct {
	abortOnFailure bool
	cb             *gobreaker.CircuitBreaker
}

var CircuteBreakerForJust429 = gobreaker.NewCircuitBreaker(gobreaker.Settings{
	Name:        "circuteBreakerForJust429",
	MaxRequests: 0,                // unlimited concurrent requests
	Interval:    10 * time.Second, // check status every 10 seconds
	Timeout:     20 * time.Second, // how long to wait before closing circuit after it's opened
	ReadyToTrip: func(counts gobreaker.Counts) bool {
		return counts.ConsecutiveFailures > 0 // instant circute openinig
	},
	IsSuccessful: func(err error) bool {
		return err.Error() != "429"
	},
})

// NewCircuitBreakerInterceptor creates a new circuit breaker interceptor
// faced status codes (codes>=400) are passed as simple string to ReadyToTrip function.
func NewCircuitBreakerInterceptor(cb *gobreaker.CircuitBreaker, abortOnFailure bool) *CircuitBreakerInterceptor {
	if cb == nil {
		panic("cb should not be nil")
	}

	return &CircuitBreakerInterceptor{
		abortOnFailure: abortOnFailure,
		cb:             cb,
	}
}

// BeforeRequest checks if the circuit breaker is open
func (c *CircuitBreakerInterceptor) BeforeRequest(data InterceptorData) InterceptorData {
	if c.cb.State() == gobreaker.StateOpen {
		if c.abortOnFailure {
			panic(models.ErrCircuitBreakerOpen)
		}
		data.Error = models.ErrCircuitBreakerOpen
		return data
	}

	return data
}

// AfterResponse records failures in the circuit breaker
func (c *CircuitBreakerInterceptor) AfterResponse(data InterceptorData) InterceptorData {
	if data.Response != nil && data.Response.StatusCode >= 400 {
		c.cb.Execute(func() (interface{}, error) {
			return nil, fmt.Errorf("%d", data.Response.StatusCode)
		})

		if c.abortOnFailure {
			panic(models.ErrTooManyRequests)
		}

		data.Error = models.ErrTooManyRequests
		return data
	}

	return data
}
