package mytarget_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Statistics(t *testing.T) {
	cfg := buildStubConfig()
	client := NewClient(cfg, nil)
	result := client.Statistics()
	assert.NotEmpty(t, result)
}
