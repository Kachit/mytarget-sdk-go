package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kachit/mytarget-sdk-go/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type RequestBuilder struct {
	cfg *config.Config
}

func (rb *RequestBuilder) buildUri(path string, query map[string]interface{}) (uri *url.URL, err error) {
	u, err := url.Parse(rb.cfg.Uri)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder@buildUri parse: %v", err)
	}
	u.Path = "/" + path
	if query != nil {
		q := u.Query()
		for k, v := range query {
			q.Set(k, fmt.Sprintf("%v", v))
		}
		u.RawQuery = q.Encode()
	}
	return u, err
}

func (rb *RequestBuilder) buildHeaders() http.Header {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+rb.cfg.AccessToken)
	return headers
}

func (rb *RequestBuilder) buildBody(data map[string]interface{}) (io.Reader, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder@buildBody json convert: %v", err)
	}
	return bytes.NewBuffer(b), nil
}

type Transport struct {
	http *http.Client
	rb   *RequestBuilder
}

func (t *Transport) Request(method string, path string, query map[string]interface{}, body map[string]interface{}) (resp *http.Response, err error) {
	//build uri
	uri, err := t.rb.buildUri(path, query)
	if err != nil {
		return nil, fmt.Errorf("transport@request build uri: %v", err)
	}
	//build body
	bodyReader, err := t.rb.buildBody(body)
	if err != nil {
		return nil, fmt.Errorf("transport@request build request body: %v", err)
	}
	//build request
	req, err := http.NewRequest(strings.ToUpper(method), uri.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("transport@request new request error: %v", err)
	}
	//build headers
	req.Header = t.rb.buildHeaders()
	return t.http.Do(req)
}

func (t *Transport) Get(path string, query map[string]interface{}) (resp *http.Response, err error) {
	return t.Request("GET", path, query, nil)
}

func (t *Transport) Delete(path string, query map[string]interface{}) (resp *http.Response, err error) {
	return t.Request("DELETE", path, query, nil)
}

func (t *Transport) Put(path string, body map[string]interface{}, query map[string]interface{}) (resp *http.Response, err error) {
	return t.Request("PUT", path, query, body)
}

func (t *Transport) Post(path string, body map[string]interface{}, query map[string]interface{}) (resp *http.Response, err error) {
	return t.Request("POST", path, query, body)
}

func NewHttpTransport(config *config.Config, h *http.Client) *Transport {
	if h == nil {
		h = &http.Client{}
	}
	return &Transport{http: h, rb: &RequestBuilder{cfg: config}}
}
