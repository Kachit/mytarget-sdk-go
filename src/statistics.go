package mytarget_sdk

import "time"

type StatisticsPartnersPadsFilter struct {
	DateFrom time.Time //Start date
	DateTo   time.Time //End date (included in the report)
	Id       string    //List of pads or pad_group identifiers. You can specify different ids separating by comma, up to 150 ids.
}

type StatisticsPadsWithSitesFilter struct {
	DateFrom time.Time //Start date
	DateTo   time.Time //End date (included in the report)
	Pads     string    //List of site identifiers.
}

type StatisticsResource struct {
}

func (sr *StatisticsResource) GetPartnersPadsList(filter *StatisticsPartnersPadsFilter) {

}

func (sr *StatisticsResource) GetPadsWithSitesList(filter *StatisticsPadsWithSitesFilter) {

}
