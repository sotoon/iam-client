package interceptor

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
	"github.com/sotoon/iam-client/pkg/client"
)

type CircuitBreakerInterceptor struct {
	cb *gobreaker.CircuitBreaker
}

// NewCircuitBreakerInterceptor creates a new circuit breaker interceptor
func NewCircuitBreakerInterceptor() *CircuitBreakerInterceptor {
	// Create circuit breaker settings
	cbSettings := gobreaker.Settings{
		Name:        "HTTP_REQUEST",
		MaxRequests: 0,                // unlimited concurrent requests
		Interval:    10 * time.Second, // check status every 10 seconds
		Timeout:     20 * time.Second, // how long to wait before closing circuit after it's opened
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 0
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("Circuit breaker '%s' changed from '%v' to '%v'", name, from, to)
			// When circuit closes again, log a message
			if from == gobreaker.StateOpen && to == gobreaker.StateHalfOpen {
				log.Printf("Circuit breaker is testing the connection again after timeout")
			}
		},
	}

	return &CircuitBreakerInterceptor{
		cb: gobreaker.NewCircuitBreaker(cbSettings),
	}
}

// BeforeRequest checks if the circuit breaker is open
func (c *CircuitBreakerInterceptor) BeforeRequest(req *http.Request) (*http.Request, error, bool) {
	// Check if circuit breaker is open
	if c.cb.State() == gobreaker.StateOpen {
		// Return error with retry=true to indicate we should retry later
		return req, errors.New("circuit breaker is open"), true
	}

	// Pass the request through unchanged, no retry needed
	return req, nil, false
}

// AfterResponse records failures in the circuit breaker
func (c *CircuitBreakerInterceptor) AfterResponse(resp *http.Response, successCode int) (*http.Response, error, bool) {
	// Check for rate limit response
	if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
		// Record failure in the circuit breaker
		c.cb.Execute(func() (interface{}, error) {
			return nil, client.ErrTooManyRequests
		})

		// Return error with retry flag set to true
		return resp, client.ErrCircuitBreakerOpen, true
	}

	// No retry needed for other responses
	return resp, nil, false
}
