package marketing

import (
	"fmt"
	"github.com/kachit/mytarget-sdk-go/common"
	"github.com/kachit/mytarget-sdk-go/http"
	"github.com/kachit/mytarget-sdk-go/resources"
	"time"
)

type StatisticsResource struct {
	*resources.ResourceAbstract
}

type StatsResource string
type StatsDimension string
type StatsFilterMetrics string

const (
	StatsResourceCampaigns StatsResource = "campaigns"
	StatsResourceBanners   StatsResource = "banners"
	StatsResourceUsers     StatsResource = "users"
)

const (
	StatsDimensionDay     StatsDimension = "day"
	StatsDimensionSummary StatsDimension = "summary"
)

const (
	StatsFilterMetricsBase StatsFilterMetrics = "base"
	StatsFilterMetricsAll  StatsFilterMetrics = "all"
)

type StatsFilter struct {
	DateFrom time.Time            //Start date. Only for day.json
	DateTo   time.Time            //	End date (included in the report). Only for day.json
	Metrics  []StatsFilterMetrics //List of metric sets. Available sets are: all, base, events, video, viral, uniques, tps, playable, romi.
	Ids      []int                //	List of banner, campaign, or user identifiers.
}

func (f *StatsFilter) Build() map[string]interface{} {
	params := make(map[string]interface{})
	params["date_from"] = f.DateFrom.Format("2006-01-02")
	params["date_to"] = f.DateTo.Format("2006-01-02")
	if len(f.Ids) > 0 {
		params["id"] = common.ArrayToString(f.Ids, ",")
	}
	return params
}

func (f *StatsFilter) IsValid() error {
	if f.DateFrom.IsZero() {
		return fmt.Errorf("StatsFilter@IsValid: %v", "DateFrom is required")
	}
	if f.DateTo.IsZero() {
		return fmt.Errorf("StatsFilter@IsValid: %v", "DateTo is required")
	}
	return nil
}

/**
* @see https://target.my.com/adv/api-marketing/doc/stat-v2
* @return StatsCollection[]Campaigns
 */
func (sr *StatisticsResource) GetCampaignStatsDaily(filter *StatsFilter) (*http.Response, error) {
	return sr.GetStats(StatsResourceCampaigns, StatsDimensionDay, filter)
}

/**
* @see https://target.my.com/adv/api-marketing/doc/stat-v2
* @return StatsCollection
 */
func (sr *StatisticsResource) GetStats(resource StatsResource, dimension StatsDimension, filter *StatsFilter) (*http.Response, error) {
	path := fmt.Sprintf("api/v2/statistics/%s/%s.json", string(resource), string(dimension))
	err := filter.IsValid()
	if err != nil {
		return nil, err
	}
	return sr.Get(path, filter.Build())
}

type StatsCollection struct {
	Items []*StatsItem          `json:"items"`
	Total *StatsCollectionTotal `json:"total"`
}

type StatsCollectionTotal struct {
	Base *StatsItemRowBase `json:"base"`
}

type StatsItem struct {
	Id    int             `json:"id"`
	Rows  []*StatsItemRow `json:"rows"`
	Total *StatsItemTotal `json:"total"`
}

type StatsItemRow struct {
	Date common.CustomDate `json:"date"`
	Base *StatsItemRowBase `json:"base"`
}

type StatsItemRowBase struct {
	Shows  int                  `json:"shows"`  //the number of impressions (times when the ad was displayed to users)
	Clicks int                  `json:"clicks"` //the number of clicks on the ad
	Goals  int                  `json:"goals"`  //the number of acheived goals (Top@Mail.ru goals for websites and installations for mobile apps)
	Spent  common.CustomFloat64 `json:"spent"`  // charged sum
	Cpm    common.CustomFloat64 `json:"cpm"`    //average cost per 1000 views
	Cpc    common.CustomFloat64 `json:"cpc"`    //average cost per 1 click
	Cpa    common.CustomFloat64 `json:"cpa"`    //average cost per 1 goal achievement
	Ctr    float64              `json:"ctr"`    //click-through rate; ratio of users who clicked on the ad to the total number of those who viewed it
	Cr     float64              `json:"cr"`     //ratio of acheived goals to the number of clicks
}

type StatsItemTotal struct {
	Base *StatsItemRowBase `json:"base"`
}
