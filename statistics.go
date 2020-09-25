package mytarget_sdk

import (
	"fmt"
	"time"
)

type StatisticsPartnersPadsFilter struct {
	DateFrom time.Time //Start date
	DateTo   time.Time //End date (included in the report)
	Ids      []int     //List of pads or pad_group identifiers. You can specify different ids separating by comma, up to 150 ids.
}

func (f *StatisticsPartnersPadsFilter) IsValid() error {
	if f.DateFrom.IsZero() {
		return fmt.Errorf("StatisticsPartnersPadsFilter@IsValid: %v", "DateFrom is required")
	}
	if f.DateTo.IsZero() {
		return fmt.Errorf("StatisticsPartnersPadsFilter@IsValid: %v", "DateTo is required")
	}
	if len(f.Ids) == 0 {
		return fmt.Errorf("StatisticsPartnersPadsFilter@IsValid: %v", "Ids is required")
	}
	return nil
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

func (f *StatisticsPadsWithSitesFilter) IsValid() error {
	if f.DateFrom.IsZero() {
		return fmt.Errorf("StatisticsPadsWithSitesFilter@IsValid: %v", "DateFrom is required")
	}
	if f.DateTo.IsZero() {
		return fmt.Errorf("StatisticsPadsWithSitesFilter@IsValid: %v", "DateTo is required")
	}
	if len(f.Pads) == 0 {
		return fmt.Errorf("StatisticsPadsWithSitesFilter@IsValid: %v", "Pads is required")
	}
	return nil
}

func (f *StatisticsPadsWithSitesFilter) Build() map[string]interface{} {
	params := make(map[string]interface{})
	params["pads"] = arrayToString(f.Pads, ",")
	params["date_from"] = f.DateFrom.Format("2006-01-02")
	params["date_to"] = f.DateTo.Format("2006-01-02")
	return params
}

type StatisticsResource struct {
	*ResourceAbstract
}

/**
 * @see https://target.my.com/help/partners/reporting_api_statistics/ru#partners
 * @return StatisticsPartnersPadsResult
 */
func (sr *StatisticsResource) GetPartnersPadsList(filter *StatisticsPartnersPadsFilter) (*Response, error) {
	err := filter.IsValid()
	if err != nil {
		return nil, err
	}
	return sr.get("api/v2/statistics/partner/pads/day.json", filter.Build())
}

/**
 * @see https://target.my.com/help/partners/reporting_api_statistics/ru#apps
 * @return StatisticsPartnersPadsResult
 */
func (sr *StatisticsResource) GetPadsWithSitesList(filter *StatisticsPadsWithSitesFilter) (*Response, error) {
	err := filter.IsValid()
	if err != nil {
		return nil, err
	}
	return sr.get("api/v2/statistics/pad_with_sites/day.json", filter.Build())
}

type StatisticsPartnersPadsResult struct {
	Total *StatisticsPartnersPadsTotal  `json:"total"`
	Items []*StatisticsPartnersPadsItem `json:"items"`
}

type StatisticsPartnersPadsTotal struct {
	Shows            int           `json:"shows"`
	Clicks           int           `json:"clicks"`
	Goals            int           `json:"goals"`
	Custom           int           `json:"custom"`
	Requests         int           `json:"requests"`
	Responses        int           `json:"responses"`
	RequestedBanners int           `json:"requested_banners"`
	NoShows          int           `json:"noshows"`
	Amount           CustomFloat64 `json:"amount"`
	Cpm              CustomFloat64 `json:"cpm"`
	Ctr              float64       `json:"ctr"`
	FillRate         float64       `json:"fill_rate"`
}

type StatisticsPartnersPadsItem struct {
	Id    int                          `json:"id"`
	Rows  []*StatisticsPartnersPadsRow `json:"rows"`
	Total *StatisticsPartnersPadsTotal `json:"total"`
}

type StatisticsPartnersPadsRow struct {
	Date             string        `json:"date"`
	Shows            int           `json:"shows"`
	Clicks           int           `json:"clicks"`
	Goals            int           `json:"goals"`
	Custom           int           `json:"custom"`
	Requests         int           `json:"requests"`
	Responses        int           `json:"responses"`
	RequestedBanners int           `json:"requested_banners"`
	NoShows          int           `json:"noshows"`
	Amount           CustomFloat64 `json:"amount"`
	Cpm              CustomFloat64 `json:"cpm"`
	Ctr              float64       `json:"ctr"`
	FillRate         float64       `json:"fill_rate"`
}
