package mytarget_sdk

import (
	"time"
)

type StatisticsPartnersPadsFilter struct {
	DateFrom time.Time //Start date
	DateTo   time.Time //End date (included in the report)
	Ids      []int     //List of pads or pad_group identifiers. You can specify different ids separating by comma, up to 150 ids.
}

func (f *StatisticsPartnersPadsFilter) Build() map[string]interface{} {
	params := make(map[string]interface{})
	params["id"] = arrayToString(f.Ids, ",")
	params["date_from"] = f.DateFrom.Format("2006-01-02")
	params["date_to"] = f.DateTo.Format("2006-01-02")
	return params
}

type StatisticsPadsWithSitesFilter struct {
	DateFrom time.Time //Start date
	DateTo   time.Time //End date (included in the report)
	Pads     []int
}

func (f *StatisticsPadsWithSitesFilter) Build() map[string]interface{} {
	params := make(map[string]interface{})
	params["pads"] = arrayToString(f.Pads, ",")
	params["date_from"] = f.DateFrom.Format("2006-01-02")
	params["date_to"] = f.DateTo.Format("2006-01-02")
	return params
}

type StatisticsResource struct {
}

func (sr *StatisticsResource) GetPartnersPadsList(filter *StatisticsPartnersPadsFilter) {

}

func (sr *StatisticsResource) GetPadsWithSitesList(filter *StatisticsPadsWithSitesFilter) {

}
