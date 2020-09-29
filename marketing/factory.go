package marketing

import (
	"github.com/kachit/mytarget-sdk-go/http"
	"github.com/kachit/mytarget-sdk-go/resources"
)

type Factory struct {
	Transport *http.Transport
}

func (f *Factory) Statistics() *StatisticsResource {
	resource := resources.NewResourceAbstract(f.Transport)
	return &StatisticsResource{ResourceAbstract: resource}
}
