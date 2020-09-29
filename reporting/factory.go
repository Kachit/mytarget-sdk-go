package reporting

import mytarget_sdk "github.com/kachit/mytarget-sdk-go"

type Factory struct {
	Transport *mytarget_sdk.Transport
}

func (f *Factory) Statistics() *StatisticsResource {
	resource := mytarget_sdk.NewResourceAbstract(f.Transport)
	return &StatisticsResource{ResourceAbstract: resource}
}
