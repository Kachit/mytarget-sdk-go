package mytarget_sdk

import (
	"encoding/json"
	"errors"
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
	*ResourceAbstract
}

/**
 * https://target.my.com/help/partners/reporting_api_statistics/ru#partners
 */
func (sr *StatisticsResource) GetPartnersPadsList(filter *StatisticsPartnersPadsFilter) (*Response, error) {
	return sr.get("api/v2/statistics/partner/pads/day.json", filter.Build())
}

/**
 * https://target.my.com/help/partners/reporting_api_statistics/ru#apps
 */
func (sr *StatisticsResource) GetPadsWithSitesList(filter *StatisticsPadsWithSitesFilter) (*Response, error) {
	return sr.get("api/v2/statistics/pad_with_sites/day.json", filter.Build())
}

type StatisticsPartnersPadsResult struct {
	Total *StatisticsPartnersPadsTotal  `json:"total"`
	Items []*StatisticsPartnersPadsItem `json:"items"`
}

type StatisticsPartnersPadsTotal struct {
	Shows            int64         `json:"shows"`
	Clicks           int64         `json:"clicks"`
	Goals            int64         `json:"goals"`
	Custom           int64         `json:"custom"`
	Requests         int64         `json:"requests"`
	Responses        int64         `json:"responses"`
	RequestedBanners int64         `json:"requested_banners"`
	NoShows          int64         `json:"noshows"`
	Amount           CustomFloat64 `json:"amount"`
	Cpm              CustomFloat64 `json:"cpm"`
	Ctr              float64       `json:"ctr"`
	FillRate         float64       `json:"fill_rate"`
}

type StatisticsPartnersPadsItem struct {
	Id    int64                        `json:"id"`
	Rows  []*StatisticsPartnersPadsRow `json:"rows"`
	Total *StatisticsPartnersPadsTotal `json:"total"`
}

type StatisticsPartnersPadsRow struct {
	Date             string        `json:"date"`
	Shows            int64         `json:"shows"`
	Clicks           int64         `json:"clicks"`
	Goals            int64         `json:"goals"`
	Custom           int64         `json:"custom"`
	Requests         int64         `json:"requests"`
	Responses        int64         `json:"responses"`
	RequestedBanners int64         `json:"requested_banners"`
	NoShows          int64         `json:"noshows"`
	Amount           CustomFloat64 `json:"amount"`
	Cpm              CustomFloat64 `json:"cpm"`
	Ctr              float64       `json:"ctr"`
	FillRate         float64       `json:"fill_rate"`
}

type CustomFloat64 struct {
	Float64 float64
}

func (cf *CustomFloat64) UnmarshalJSON(data []byte) error {
	if data[0] == 34 {
		err := json.Unmarshal(data[1:len(data)-1], &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	} else {
		err := json.Unmarshal(data, &cf.Float64)
		if err != nil {
			return errors.New("CustomFloat64: UnmarshalJSON: " + err.Error())
		}
	}
	return nil
}
