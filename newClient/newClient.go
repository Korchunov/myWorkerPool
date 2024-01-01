package newClient

import (
	"errors"
	"net/http"
	"time"
)

type Client struct {
	C *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout is required parameter")
	}
	return &Client{
		C: &http.Client{
			Timeout: timeout,
		},
	}, nil
}
