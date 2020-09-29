package marketing

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Marketing_Statistics_StatsFilter_BuildByDefault(t *testing.T) {
	filter := StatsFilter{}
	filter.DateFrom = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	expected := make(map[string]interface{})
	expected["date_from"] = "2020-01-10"
	expected["date_to"] = "2020-01-20"
	result := filter.Build()
	assert.Equal(t, expected, result)
}

func Test_Marketing_Statistics_StatsFilter_BuildFull(t *testing.T) {
	filter := StatsFilter{}
	filter.DateFrom = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	filter.Ids = []int{1, 2, 3}
	expected := make(map[string]interface{})
	expected["id"] = "1,2,3"
	expected["date_from"] = "2020-01-10"
	expected["date_to"] = "2020-01-20"
	result := filter.Build()
	assert.Equal(t, expected, result)
}

func Test_Marketing_Statistics_StatsFilter_IsValidSuccess(t *testing.T) {
	filter := StatsFilter{}
	filter.DateFrom = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	assert.Nil(t, filter.IsValid())
}

func Test_Marketing_Statistics_StatsFilter_IsValidFailedDateFrom(t *testing.T) {
	filter := StatsFilter{}
	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "StatsFilter@IsValid: DateFrom is required", err.Error())
}

func Test_Marketing_Statistics_StatsFilter_IsValidFailedDateTo(t *testing.T) {
	filter := StatsFilter{}
	filter.DateFrom = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "StatsFilter@IsValid: DateTo is required", err.Error())
}
