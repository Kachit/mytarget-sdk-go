package management

import (
	"github.com/kachit/mytarget-sdk-go/http"
	"github.com/kachit/mytarget-sdk-go/resources"
)

type Factory struct {
	Transport *http.Transport
}

func (f *Factory) Campaigns() *CampaignsResource {
	resource := resources.NewResourceAbstract(f.Transport)
	return &CampaignsResource{ResourceAbstract: resource}
}
