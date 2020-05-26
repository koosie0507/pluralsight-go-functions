package http

import "net/http"

type RoundTripCounter struct {
	count int
}

func (rt *RoundTripCounter) RoundTrip(r *http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}
