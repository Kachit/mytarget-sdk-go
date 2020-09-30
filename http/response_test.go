package http

import (
	"github.com/kachit/mytarget-sdk-go/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_HTTP_Response_IsSuccessTrue(t *testing.T) {
	response := &Response{raw: stubs.BuildStubResponseFromFile(http.StatusOK, "../stubs/data/management/campaigns/list.success.json")}
	assert.True(t, response.IsSuccess())
}

func Test_HTTP_Response_IsSuccessFalse(t *testing.T) {
	response := &Response{raw: stubs.BuildStubResponseFromFile(http.StatusBadRequest, "../stubs/data/management/campaigns/list.success.json")}
	assert.False(t, response.IsSuccess())
}

func Test_HTTP_GetRawResponse(t *testing.T) {
	rsp := stubs.BuildStubResponseFromFile(http.StatusOK, "../stubs/data/management/campaigns/list.success.json")
	response := &Response{raw: rsp}
	assert.NotEmpty(t, response.GetRawResponse())
}

func Test_HTTP_GetError(t *testing.T) {
	rsp := stubs.BuildStubResponseFromFile(http.StatusBadRequest, "../stubs/data/common/auth/invalid_token.error.json")
	response := &Response{raw: rsp}
	result, _ := response.GetError()
	assert.Equal(t, "Unknown access token", result.Error.Message)
	assert.Equal(t, "invalid_token", result.Error.Code)
}

func Test_HTTP_GetRawBody(t *testing.T) {
	rsp := stubs.BuildStubResponseFromFile(http.StatusBadRequest, "../stubs/data/management/campaigns/list.success.json")
	response := &Response{raw: rsp}
	assert.NotEmpty(t, response.GetRawBody())
}

func Test_HTTP_NewResponse(t *testing.T) {
	rsp := stubs.BuildStubResponseFromFile(http.StatusOK, "../stubs/data/management/campaigns/list.success.json")
	response := NewResponse(rsp)
	assert.NotEmpty(t, response)
}
