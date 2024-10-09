package middlewares

import (
	"net/http"
	"time"
)

type RetryTransport struct {
	Base http.RoundTripper
}

func NewRetryTransport(base http.RoundTripper) *RetryTransport {
	return &RetryTransport{Base: base}
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	for attempt := 0; attempt < 4; attempt++ {
		resp, err = t.Base.RoundTrip(req)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != 420 && resp.StatusCode != 429 {
			return resp, nil
		}

		time.Sleep(500 * time.Millisecond)
	}
	return resp, err
}
