package management

import (
	"github.com/kachit/mytarget-sdk-go/config"
	"github.com/kachit/mytarget-sdk-go/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Management_Factory_GetCampaignsResource(t *testing.T) {
	factory := Factory{Transport: http.NewHttpTransport(&config.Config{}, nil)}
	assert.NotEmpty(t, factory.Campaigns())
}
