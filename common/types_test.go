package common

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Common_Types_CustomFloat64_UnmarshalJSONStringSuccess(t *testing.T) {
	c := CustomFloat64{}
	_ = c.UnmarshalJSON([]byte(`"10.10"`))
	assert.Equal(t, 10.10, c.Value())
}

func Test_Common_Types_CustomFloat64_UnmarshalJSONStringError(t *testing.T) {
	c := CustomFloat64{}
	err := c.UnmarshalJSON([]byte(`"qwerty"`))
	assert.Error(t, err)
	assert.Equal(t, "CustomFloat64@UnmarshalJSON: invalid character 'q' looking for beginning of value", err.Error())
}

func Test_Common_Types_CustomFloat64_UnmarshalJSONFloatSuccess(t *testing.T) {
	c := CustomFloat64{}
	_ = c.UnmarshalJSON([]byte(`10.10`))
	assert.Equal(t, 10.10, c.Value())
}

func Test_Common_Types_CustomFloat64_UnmarshalJSONFloatError(t *testing.T) {
	c := CustomFloat64{}
	err := c.UnmarshalJSON([]byte(`qwerty`))
	assert.Error(t, err)
	assert.Equal(t, "CustomFloat64@UnmarshalJSON: invalid character 'q' looking for beginning of value", err.Error())
}

func Test_Common_Types_CustomFloat64_MarshalJSONSuccess(t *testing.T) {
	c := CustomFloat64{}
	c.Float64 = 10.10
	result, err := c.MarshalJSON()
	expected, _ := json.Marshal(c.Float64)
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_Common_Types_CustomFloat64_MarshalJSONEmpty(t *testing.T) {
	c := CustomFloat64{}
	result, err := c.MarshalJSON()
	expected, _ := json.Marshal(c.Float64)
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_Common_Types_CustomTimestamp_UnmarshalJSONSuccess(t *testing.T) {
	c := CustomTimestamp{}
	format := "2006-01-02 15:04:05"
	str := `"2020-09-10 15:15:15"`
	_ = c.UnmarshalJSON([]byte(str))
	assert.Equal(t, "2020-09-10 15:15:15", c.Value().Format(format))
}

func Test_Common_Types_CustomTimestamp_UnmarshalJSONError(t *testing.T) {
	c := CustomTimestamp{}
	str := `123456`
	err := c.UnmarshalJSON([]byte(str))
	assert.Error(t, err)
	assert.Equal(t, "CustomTimestamp@UnmarshalJSON: json: cannot unmarshal number into Go value of type string", err.Error())
}

func Test_Common_Types_CustomTimestamp_UnmarshalJSONErrorTimeParse(t *testing.T) {
	c := CustomTimestamp{}
	str := `"2006/01/02"`
	err := c.UnmarshalJSON([]byte(str))
	assert.Error(t, err)
	assert.Equal(t, `CustomTimestamp@UnmarshalJSON ParseTime: parsing time "2006/01/02" as "2006-01-02 15:04:05": cannot parse "/01/02" as "-"`, err.Error())
}

func Test_Common_Types_CustomTimestamp_MarshalJSONSuccess(t *testing.T) {
	c := CustomTimestamp{}
	c.Timestamp = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`"2020-01-20 00:00:00"`), result)
	assert.Nil(t, err)
}

func Test_Common_Types_CustomTimestamp_MarshalJSONEmpty(t *testing.T) {
	c := CustomTimestamp{}
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(``), result)
	assert.Nil(t, err)
}

//func Test_Common_Types_CustomTimestamp_MarshalJSONError(t *testing.T) {
//	c := CustomTimestamp{}
//	result, err := c.MarshalJSON()
//	assert.Equal(t, []byte(``), result)
//	assert.Nil(t, err)
//}

func Test_Common_Types_CustomDate_UnmarshalJSONSuccess(t *testing.T) {
	c := CustomDate{}
	format := "2006-01-02"
	str := `"2020-09-10"`
	_ = c.UnmarshalJSON([]byte(str))
	assert.Equal(t, "2020-09-10", c.Value().Format(format))
}

func Test_Common_Types_CustomDate_UnmarshalJSONError(t *testing.T) {
	c := CustomDate{}
	str := `123456`
	err := c.UnmarshalJSON([]byte(str))
	assert.Error(t, err)
	assert.Equal(t, "CustomDate@UnmarshalJSON: json: cannot unmarshal number into Go value of type string", err.Error())
}

func Test_Common_Types_CustomDate_UnmarshalJSONErrorTimeParse(t *testing.T) {
	c := CustomDate{}
	str := `"2006/01/02"`
	err := c.UnmarshalJSON([]byte(str))
	assert.Error(t, err)
	assert.Equal(t, `CustomDate@UnmarshalJSON ParseTime: parsing time "2006/01/02" as "2006-01-02": cannot parse "/01/02" as "-"`, err.Error())
}

func Test_Common_Types_CustomDate_MarshalJSONSuccess(t *testing.T) {
	c := CustomDate{}
	c.Date = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(`"2020-01-20"`), result)
	assert.Nil(t, err)
}

func Test_Common_Types_CustomDate_MarshalJSONEmpty(t *testing.T) {
	c := CustomDate{}
	result, err := c.MarshalJSON()
	assert.Equal(t, []byte(``), result)
	assert.Nil(t, err)
}

//func Test_Common_Types_CustomDate_MarshalJSONError(t *testing.T) {
//	c := CustomDate{}
//	result, err := c.MarshalJSON()
//	assert.Equal(t, []byte(``), result)
//	assert.Nil(t, err)
//}
