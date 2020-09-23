package mytarget_sdk

import "net/http"

type Client struct {
	config *Config
	http   *http.Client
}

func NewClient(config *Config) *Client {
	return &Client{config: config}
}
