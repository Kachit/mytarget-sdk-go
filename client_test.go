package mytarget_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Statistics(t *testing.T) {
	client := buildStubClient()
	result := client.Statistics()
	assert.NotEmpty(t, result)
}

func TestClient_Campaigns(t *testing.T) {
	client := buildStubClient()
	result := client.Campaigns()
	assert.NotEmpty(t, result)
}
