package http

import (
	"github.com/kachit/mytarget-sdk-go/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_HTTP_Response_IsSuccessTrue(t *testing.T) {
	response := &Response{raw: stubs.BuildStubResponseFromFile(http.StatusOK, "stubs/data/management/campaigns/list.success.json")}
	assert.True(t, response.IsSuccess())
}

func Test_HTTP_Response_IsSuccessFalse(t *testing.T) {
	response := &Response{raw: stubs.BuildStubResponseFromFile(http.StatusBadRequest, "stubs/data/management/campaigns/list.success.json")}
	assert.False(t, response.IsSuccess())
}

func Test_HTTP_NewResponse(t *testing.T) {
	rsp := stubs.BuildStubResponseFromFile(http.StatusBadRequest, "stubs/data/management/campaigns/list.success.json")
	response := NewResponse(rsp)
	assert.NotEmpty(t, response)
}
