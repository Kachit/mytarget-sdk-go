package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Common_Types_CustomFloat64_UnmarshalJSONString(t *testing.T) {
	c := CustomFloat64{}
	_ = c.UnmarshalJSON([]byte(`"10.10"`))
	assert.Equal(t, 10.10, c.Value())
}

func Test_Common_Types_CustomFloat64_UnmarshalJSONFloat(t *testing.T) {
	c := CustomFloat64{}
	_ = c.UnmarshalJSON([]byte(`10.10`))
	assert.Equal(t, 10.10, c.Value())
}

func Test_Common_Types_CustomTimestamp_UnmarshalJSON(t *testing.T) {
	c := CustomTimestamp{}
	format := "2006-01-02 15:04:05"
	str := `"2020-09-10 15:15:15"`
	_ = c.UnmarshalJSON([]byte(str))
	assert.Equal(t, "2020-09-10 15:15:15", c.Value().Format(format))
}

func Test_Common_Types_CustomDate_UnmarshalJSON(t *testing.T) {
	c := CustomDate{}
	format := "2006-01-02"
	str := `"2020-09-10"`
	_ = c.UnmarshalJSON([]byte(str))
	assert.Equal(t, "2020-09-10", c.Value().Format(format))
}
