package mytarget

import (
	"github.com/kachit/mytarget-sdk-go/config"
	mytarget_http "github.com/kachit/mytarget-sdk-go/http"
	"github.com/kachit/mytarget-sdk-go/management"
	"github.com/kachit/mytarget-sdk-go/marketing"
	"github.com/kachit/mytarget-sdk-go/reporting"
	"net/http"
)

type Client struct {
	transport *mytarget_http.Transport
}

func (c *Client) Reporting() *reporting.Factory {
	return &reporting.Factory{Transport: c.transport}
}

func (c *Client) Management() *management.Factory {
	return &management.Factory{Transport: c.transport}
}

func (c *Client) Marketing() *marketing.Factory {
	return &marketing.Factory{Transport: c.transport}
}

func NewClientFromConfig(config *config.Config, cl *http.Client) *Client {
	if cl == nil {
		cl = &http.Client{}
	}
	transport := mytarget_http.NewHttpTransport(config, cl)
	return &Client{transport}
}

func NewClientFromCredentials(accessToken string, cl *http.Client) *Client {
	cfg := config.NewConfig(accessToken)
	if cl == nil {
		cl = &http.Client{}
	}
	transport := mytarget_http.NewHttpTransport(cfg, cl)
	return &Client{transport}
}
