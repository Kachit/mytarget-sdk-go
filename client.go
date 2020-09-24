package mytarget_sdk

import "net/http"

type Client struct {
	transport *Transport
}

func (c *Client) Statistics() *StatisticsResource {
	resource := newResourceAbstract(c.transport)
	return &StatisticsResource{ResourceAbstract: resource}
}

func NewClient(config *Config, cl *http.Client) *Client {
	if cl == nil {
		cl = &http.Client{}
	}
	transport := newHttpTransport(config, cl)
	return &Client{transport}
}
