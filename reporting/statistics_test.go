package reporting

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Reporting_Statistics_StatisticsPartnersPadsFilter_Build(t *testing.T) {
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

func Test_Reporting_Statistics_StatisticsPartnersPadsFilter_IsValidSuccess(t *testing.T) {
	filter := StatisticsPartnersPadsFilter{}
	filter.DateFrom = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	filter.Ids = []int{1, 2, 3}
	assert.Nil(t, filter.IsValid())
}

func Test_Reporting_Statistics_StatisticsPartnersPadsFilter_IsValidFailedDateFrom(t *testing.T) {
	filter := StatisticsPartnersPadsFilter{}
	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	filter.Ids = []int{1, 2, 3}
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "StatisticsPartnersPadsFilter@IsValid: DateFrom is required", err.Error())
}

func Test_Reporting_Statistics_StatisticsPartnersPadsFilter_IsValidFailedDateTo(t *testing.T) {
	filter := StatisticsPartnersPadsFilter{}
	filter.DateFrom = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	filter.Ids = []int{1, 2, 3}
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "StatisticsPartnersPadsFilter@IsValid: DateTo is required", err.Error())
}

func Test_Reporting_Statistics_StatisticsPartnersPadsFilter_IsValidFailedIds(t *testing.T) {
	filter := StatisticsPartnersPadsFilter{}
	filter.DateFrom = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "StatisticsPartnersPadsFilter@IsValid: Ids is required", err.Error())
}

func Test_Reporting_Statistics_StatisticsPadsWithSitesFilter_Build(t *testing.T) {
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

func Test_Reporting_Statistics_StatisticsPadsWithSitesFilter_IsValidSuccess(t *testing.T) {
	filter := StatisticsPadsWithSitesFilter{}
	filter.DateFrom = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	filter.Pads = []int{1, 2, 3}
	assert.Nil(t, filter.IsValid())
}

func Test_Reporting_Statistics_StatisticsPadsWithSitesFilter_IsValidFailedDateFrom(t *testing.T) {
	filter := StatisticsPadsWithSitesFilter{}
	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	filter.Pads = []int{1, 2, 3}
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "StatisticsPadsWithSitesFilter@IsValid: DateFrom is required", err.Error())
}

func Test_Reporting_Statistics_StatisticsPadsWithSitesFilter_IsValidFailedDateTo(t *testing.T) {
	filter := StatisticsPadsWithSitesFilter{}
	filter.DateFrom = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	filter.Pads = []int{1, 2, 3}
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "StatisticsPadsWithSitesFilter@IsValid: DateTo is required", err.Error())
}

func Test_Reporting_Statistics_StatisticsPadsWithSitesFilter_IsValidFailedIds(t *testing.T) {
	filter := StatisticsPadsWithSitesFilter{}
	filter.DateFrom = time.Date(2020, time.Month(1), 10, 0, 0, 0, 0, time.UTC)
	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "StatisticsPadsWithSitesFilter@IsValid: Pads is required", err.Error())
}

//func Test_Reporting_Statistics_StatisticsResource_GetPadsWithSitesListInvalidDateFrom(t *testing.T) {
//	filter := &StatisticsPadsWithSitesFilter{}
//	client := mytarget_sdk.BuildStubClient()
//	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
//	filter.Pads = []int{1, 2, 3}
//	_, err := client.ReportsStatistics().GetPadsWithSitesList(filter)
//	assert.Error(t, err)
//	assert.Equal(t, "StatisticsPadsWithSitesFilter@IsValid: DateFrom is required", err.Error())
//}
//
//func Test_Reporting_Statistics_StatisticsResource_GetPartnersPadsListInvalidDateFrom(t *testing.T) {
//	filter := &StatisticsPartnersPadsFilter{}
//	client := mytarget_sdk.BuildStubClient()
//	filter.DateTo = time.Date(2020, time.Month(1), 20, 0, 0, 0, 0, time.UTC)
//	filter.Ids = []int{1, 2, 3}
//	_, err := client.ReportsStatistics().GetPartnersPadsList(filter)
//	assert.Error(t, err)
//	assert.Equal(t, "StatisticsPartnersPadsFilter@IsValid: DateFrom is required", err.Error())
//}
