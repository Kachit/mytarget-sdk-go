package mytarget_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_StatisticsResource(t *testing.T) {
	client := buildStubClient()
	result := client.Statistics()
	assert.NotEmpty(t, result)
}

func Test_Client_CampaignsResource(t *testing.T) {
	client := buildStubClient()
	result := client.Campaigns()
	assert.NotEmpty(t, result)
}
