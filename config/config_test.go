package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Config_NewConfig(t *testing.T) {
	config := NewConfig("access token")
	assert.NotEmpty(t, config)
	assert.Equal(t, "access token", config.AccessToken)
	assert.Equal(t, UriProd, config.Uri)
}
