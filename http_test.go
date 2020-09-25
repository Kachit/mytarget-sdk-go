package mytarget_sdk

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRequestBuilder_BuildUriWithoutQueryParams(t *testing.T) {
	cfg := buildStubConfig()
	builder := RequestBuilder{cfg: cfg}
	uri, err := builder.buildUri("qwerty", nil)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty", uri.String())
	assert.Nil(t, err)
}

func TestRequestBuilder_BuildUriWithQueryParams(t *testing.T) {
	cfg := buildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	uri, err := builder.buildUri("qwerty", data)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty?bar=baz&foo=bar", uri.String())
	assert.Nil(t, err)
}

func TestRequestBuilder_BuildHeaders(t *testing.T) {
	cfg := buildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	headers := builder.buildHeaders()
	assert.NotEmpty(t, headers)
	assert.Equal(t, "application/json", headers.Get("Content-Type"))
	assert.Equal(t, "Bearer "+cfg.AccessToken, headers.Get("Authorization"))
}

func TestRequestBuilder_BuildBody(t *testing.T) {
	cfg := buildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	body, _ := builder.buildBody(data)
	assert.NotEmpty(t, body)
}

func TestTransport_Request(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := buildStubConfig()
	transport := newHttpTransport(cfg, nil)

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo",
		httpmock.NewStringResponder(http.StatusOK, buildStubResponseData()))

	resp, _ := transport.request("GET", "foo", nil, nil)
	assert.NotEmpty(t, resp)
}
