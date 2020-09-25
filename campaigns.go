package mytarget_sdk

type CampaignsResource struct {
	*ResourceAbstract
}

type CampaignsListFilter struct {
	Limit  int
	Offset int
}

func (f *CampaignsListFilter) Build() map[string]interface{} {
	params := make(map[string]interface{})
	if f.Limit != 0 {
		params["limit"] = f.Limit
	}
	if f.Offset != 0 {
		params["offset"] = f.Offset
	}
	return params
}

/**
 * @see https://target.my.com/doc/api/ru/resource/Campaigns
 * @return CampaignsCollection
 */
func (cr *CampaignsResource) GetList(filter *CampaignsListFilter) (*Response, error) {
	if filter == nil {
		filter = &CampaignsListFilter{}
	}
	return cr.get("api/v2/campaigns.json", filter.Build())
}

type CampaignsCollection struct {
	Count  int                `json:"count"`
	Offset int                `json:"offset"`
	Items  []*CampaignsObject `json:"items"`
}

type CampaignsObject struct {
	Id          int    `json:"id"`
	PackageId   int    `json:"package_id"`
	Name        string `json:"name"`
	LastUpdated string `json:"last_updated"`
}
