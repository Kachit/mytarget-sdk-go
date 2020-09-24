package mytarget_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Statistics_StatisticsPartnersPadsFilter_Build(t *testing.T) {
	filter := StatisticsPartnersPadsFilter{}
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

func Test_Statistics_StatisticsPadsWithSitesFilter_Build(t *testing.T) {
	filter := StatisticsPadsWithSitesFilter{}
	filter.DateFrom = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	filter.Pads = []int{1, 2, 3}
	expected := make(map[string]interface{})
	expected["pads"] = "1,2,3"
	expected["date_from"] = "2020-01-10"
	expected["date_to"] = "2020-01-20"
	result := filter.Build()
	assert.Equal(t, expected, result)
}
