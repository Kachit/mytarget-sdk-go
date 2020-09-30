package resources

import (
	"github.com/jarcoal/httpmock"
	mytarget_sdk "github.com/kachit/mytarget-sdk-go/http"
	"github.com/kachit/mytarget-sdk-go/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_Resources_Resource_Get(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := stubs.BuildStubConfig()
	transport := mytarget_sdk.NewHttpTransport(cfg, nil)
	resource := NewResourceAbstract(transport)

	body, _ := stubs.LoadStubResponseData("../stubs/data/management/campaigns/list.success.json")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := resource.Get("foo", nil)
	assert.NotEmpty(t, resp)
}

func Test_Resources_Resource_Post(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := stubs.BuildStubConfig()
	transport := mytarget_sdk.NewHttpTransport(cfg, nil)
	resource := NewResourceAbstract(transport)

	body, _ := stubs.LoadStubResponseData("../stubs/data/management/campaigns/list.success.json")

	httpmock.RegisterResponder("POST", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := resource.Post("foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_Resources_NewResourceAbstract(t *testing.T) {
	cfg := stubs.BuildStubConfig()
	transport := mytarget_sdk.NewHttpTransport(cfg, nil)
	resource := NewResourceAbstract(transport)
	assert.NotEmpty(t, resource)
}
