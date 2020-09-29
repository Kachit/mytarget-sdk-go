package resources

import (
	"github.com/kachit/mytarget-sdk-go/http"
	"github.com/kachit/mytarget-sdk-go/stubs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Resources_NewResourceAbstract(t *testing.T) {
	cfg := stubs.BuildStubConfig()
	transport := http.NewHttpTransport(cfg, nil)
	resource := NewResourceAbstract(transport)
	assert.NotEmpty(t, resource)
}
