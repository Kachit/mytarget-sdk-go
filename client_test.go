package mytarget_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func Test_Client_StatisticsResource(t *testing.T) {
//	client := BuildStubClient()
//	result := client.ReportingStatistics()
//	assert.NotEmpty(t, result)
//}

func Test_Client_CampaignsResource(t *testing.T) {
	client := BuildStubClient()
	result := client.Campaigns()
	assert.NotEmpty(t, result)
}
