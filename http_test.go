package mytarget_sdk

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_HTTP_RequestBuilder_BuildUriWithoutQueryParams(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}
	uri, err := builder.buildUri("qwerty", nil)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty", uri.String())
	assert.Nil(t, err)
}

func Test_HTTP_RequestBuilder_BuildUriWithQueryParams(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	uri, err := builder.buildUri("qwerty", data)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty?bar=baz&foo=bar", uri.String())
	assert.Nil(t, err)
}

func Test_HTTP_RequestBuilder_BuildHeaders(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	headers := builder.buildHeaders()
	assert.NotEmpty(t, headers)
	assert.Equal(t, "application/json", headers.Get("Content-Type"))
	assert.Equal(t, "Bearer "+cfg.AccessToken, headers.Get("Authorization"))
}

func Test_HTTP_RequestBuilder_BuildBody(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	body, _ := builder.buildBody(data)
	assert.NotEmpty(t, body)
}

func Test_HTTP_Transport_Request(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := newHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/campaigns/list.success.json")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.request("GET", "foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestGET(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := newHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/campaigns/list.success.json")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.get("foo", nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestPOST(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := newHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/campaigns/list.success.json")

	httpmock.RegisterResponder("POST", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.post("foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestPUT(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := newHttpTransport(cfg, nil)

	body, _ := LoadStubResponseData("stubs/campaigns/list.success.json")

	httpmock.RegisterResponder("PUT", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.put("foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Response_IsSuccessTrue(t *testing.T) {
	response := &Response{raw: BuildStubResponseFromFile(http.StatusOK, "stubs/campaigns/list.success.json")}
	assert.True(t, response.IsSuccess())
}

func Test_HTTP_Response_IsSuccessFalse(t *testing.T) {
	response := &Response{raw: BuildStubResponseFromFile(http.StatusBadRequest, "stubs/campaigns/list.success.json")}
	assert.False(t, response.IsSuccess())
}
