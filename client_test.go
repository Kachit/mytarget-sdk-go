package mytarget

import (
	"github.com/kachit/mytarget-sdk-go/stubs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_GetMarketingAPI(t *testing.T) {
	client := NewClientFromConfig(stubs.BuildStubConfig(), nil)
	result := client.Marketing()
	assert.NotEmpty(t, result)
}

func Test_Client_GetReportingAPI(t *testing.T) {
	client := NewClientFromConfig(stubs.BuildStubConfig(), nil)
	result := client.Reporting()
	assert.NotEmpty(t, result)
}

func Test_Client_GetManagementAPI(t *testing.T) {
	client := NewClientFromConfig(stubs.BuildStubConfig(), nil)
	result := client.Management()
	assert.NotEmpty(t, result)
}

func Test_Client_NewClientFromConfig(t *testing.T) {
	client := NewClientFromConfig(stubs.BuildStubConfig(), nil)
	assert.NotEmpty(t, client)
}

func Test_Client_NewClientFromCredentials(t *testing.T) {
	client := NewClientFromCredentials("access token", nil)
	assert.NotEmpty(t, client)
}
