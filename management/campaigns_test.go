package management

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Management_Campaigns_CampaignsListFilter_BuildByDefault(t *testing.T) {
	filter := CampaignsListFilter{}
	expected := make(map[string]interface{})
	result := filter.Build()
	assert.Equal(t, expected, result)
}

func Test_Management_Campaigns_CampaignsListFilter_BuildByFilled(t *testing.T) {
	filter := CampaignsListFilter{}
	filter.Limit = 20
	filter.Offset = 10
	expected := make(map[string]interface{})
	expected["limit"] = 20
	expected["offset"] = 10
	result := filter.Build()
	assert.Equal(t, expected, result)
}
