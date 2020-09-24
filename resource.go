package mytarget_sdk

type ResourceAbstract struct {
	tr *Transport
}

func newResourceAbstract(transport *Transport) *ResourceAbstract {
	return &ResourceAbstract{tr: transport}
}
