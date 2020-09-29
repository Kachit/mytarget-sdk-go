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
	DateFrom time.Time
	DateTo   time.Time
	Metrics  []StatsFilterMetrics
	Ids      []int
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

func (sr *StatisticsResource) GetCampaignStatsDaily(filter *StatsFilter) (*http.Response, error) {
	return sr.GetStats(StatsResourceCampaigns, StatsDimensionDay, filter)
}

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
	Shows  int                  `json:"shows"`
	Clicks int                  `json:"clicks"`
	Goals  int                  `json:"goals"`
	Spent  common.CustomFloat64 `json:"spent"`
	Cpm    common.CustomFloat64 `json:"cpm"`
	Cpc    common.CustomFloat64 `json:"cpc"`
	Cpa    common.CustomFloat64 `json:"cpa"`
	Ctr    float64              `json:"ctr"`
	Cr     float64              `json:"cr"`
}

type StatsItemTotal struct {
	Base *StatsItemRowBase `json:"base"`
}
