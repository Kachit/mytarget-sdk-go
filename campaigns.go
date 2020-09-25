package mytarget_sdk

type CampaignsResource struct {
	*ResourceAbstract
}

type CampaignsListFilter struct {
}

func (f *CampaignsListFilter) Build() map[string]interface{} {
	params := make(map[string]interface{})
	return params
}

/**
 * @see https://target.my.com/doc/api/ru/resource/Campaigns
 * @return CampaignsCollection
 */
func (sr *StatisticsResource) GetList(filter *CampaignsListFilter) (*Response, error) {
	return sr.get("api/v2/campaigns.json", filter.Build())
}

type CampaignsCollection struct {
	Count  int64              `json:"count"`
	Offset int64              `json:"offset"`
	Items  []*CampaignsObject `json:"items"`
}

type CampaignsObject struct {
	Id          int64  `json:"id"`
	PackageId   int64  `json:"package_id"`
	Name        string `json:"name"`
	LastUpdated string `json:"last_updated"`
}
