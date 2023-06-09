package httptools

import (
	"net/http"
	"time"
)

const (
	DefaultRetryCount = 5
	DefaultRetryDelay = 2 // seconds
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type retryclient struct {
	client  *http.Client
	retries int
	delay   time.Duration
	allowed []int
}

func NewRetryClient(r, d int) HTTPClient {
	return &retryclient{
		client:  &http.Client{},
		retries: r,
		delay:   time.Duration(d) * time.Second,
		allowed: []int{
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusTooManyRequests,
		},
	}
}

func (c *retryclient) isTemporaryError(cd int) bool {
	for _, code := range c.allowed {
		if cd == code {
			return true
		}
	}
	return false
}

func (c *retryclient) Do(r *http.Request) (*http.Response, error) {
	var res *http.Response
	var err error
	for i := 0; i < c.retries; i++ {
		res, err = c.client.Do(r)
		if err != nil {
			continue
		}
		if c.isTemporaryError(res.StatusCode) {
			time.Sleep(c.delay)
		} else {
			break
		}
	}
	return res, err
}

var RetryClient HTTPClient = NewRetryClient(DefaultRetryCount, DefaultRetryDelay)
