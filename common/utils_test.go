package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Utils_ArrayToStringSingle(t *testing.T) {
	result := ArrayToString([]int{1}, ",")
	expected := "1"
	assert.Equal(t, expected, result)
}

func Test_Utils_ArrayToStringMultiple(t *testing.T) {
	result := ArrayToString([]int{1, 2, 3}, ",")
	expected := "1,2,3"
	assert.Equal(t, expected, result)
}
