package interceptor

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

type ClientHandler interface {
	ProcessRequest(httpRequest *http.Request, successCode int, id string) (*http.Response, error)
}

type Backoff interface {
	TimeToWait(iteration int) time.Duration
}

type RetryInterceptor struct {
	maxRetries      int
	maxBackoff      time.Duration
	cache           *cache.Cache
	clientHandler   ClientHandler
	backoffStrategy Backoff
}

type retryInternalData struct {
	retryCount int
}

func NewRetryInterceptor(maxRetries int, maxBackoffDuration time.Duration, clientHandler ClientHandler, backoffStrategy Backoff) *RetryInterceptor {
	return &RetryInterceptor{
		maxRetries:      maxRetries,
		maxBackoff:      maxBackoffDuration,
		cache:           cache.New(time.Minute, time.Minute*15),
		clientHandler:   clientHandler,
		backoffStrategy: backoffStrategy,
	}
}

func (e *RetryInterceptor) BeforeRequest(data InterceptorData) InterceptorData {
	if data.Error != nil {
		e.sleep(data)
		response, err := e.clientHandler.ProcessRequest(data.Request, 0, data.ID)
		if err != nil || response.StatusCode >= 400 {
			data.Error = err
			return e.BeforeRequest(data)
		}
		data.Error = nil
		data.Response = response
	}
	return data
}

func (e *RetryInterceptor) AfterResponse(data InterceptorData) InterceptorData {

	if data.Error != nil || data.Response.StatusCode >= 400 {
		e.sleep(data)
		response, err := e.clientHandler.ProcessRequest(data.InitialRequest, 0, data.ID)
		data.Response = response
		if err != nil || response.StatusCode >= 400 {
			data.Error = err
			return e.AfterResponse(data)
		}
		data.Error = nil
	}
	return data
}

func (e *RetryInterceptor) sleep(data InterceptorData) {
	internalData, found := e.cache.Get(data.ID)
	var d retryInternalData
	if found {
		d = internalData.(retryInternalData)
		if d.retryCount >= e.maxRetries {
			panic(data.Error)
		}
		d.retryCount++
		e.cache.Set(data.ID, d, cache.DefaultExpiration)
	} else {
		d := retryInternalData{retryCount: 1}
		e.cache.Set(data.ID, d, cache.DefaultExpiration)
	}
	time.Sleep(e.backoffStrategy.TimeToWait(d.retryCount))
}

/////////////////////////////////////////

type BackoffStrategyExpnential struct {
	baseDuration time.Duration
	maxBackoff   time.Duration
}

func NewRetryInterceptor_ExponentialBackoff(baseDuration, maxBackoff time.Duration) BackoffStrategyExpnential {
	return BackoffStrategyExpnential{
		baseDuration: baseDuration,
		maxBackoff:   maxBackoff,
	}
}

func (b BackoffStrategyExpnential) TimeToWait(iteration int) time.Duration {
	return exponentialBackoff(iteration, b.baseDuration, b.maxBackoff)
}

// exponentialBackoff calculates the backoff duration with jitter for retries
func exponentialBackoff(retry int, initialBackoff, maxBackoff time.Duration) time.Duration {
	if retry <= 0 {
		return initialBackoff
	}

	backoff := initialBackoff * time.Duration(1<<uint(retry))

	jitter := time.Duration(rand.Int63n(int64(backoff) / 2))
	backoff = backoff + jitter

	if backoff > maxBackoff {
		backoff = maxBackoff
	}

	return backoff
}

type BackoffStrategyLinier struct {
	baseDuration time.Duration
}

func NewRetryInterceptor_BackoffStrategyLinier(baseDuration time.Duration) BackoffStrategyExpnential {
	return BackoffStrategyExpnential{
		baseDuration: baseDuration,
	}
}

func (b BackoffStrategyLinier) TimeToWait(iteration int) time.Duration {
	return b.baseDuration
}
