package mytarget_sdk

import "fmt"

type ResourceAbstract struct {
	tr *Transport
}

func (r *ResourceAbstract) Get(path string, query map[string]interface{}) (*Response, error) {
	rsp, err := r.tr.get(path, query)
	if err != nil {
		return nil, fmt.Errorf("ResourceAbstract@get request: %v", err)
	}
	return &Response{raw: rsp}, nil
}

func (r *ResourceAbstract) Post(path string, body map[string]interface{}, query map[string]interface{}) (*Response, error) {
	rsp, err := r.tr.post(path, body, query)
	if err != nil {
		return nil, fmt.Errorf("ResourceAbstract@post request: %v", err)
	}
	return &Response{raw: rsp}, nil
}

func NewResourceAbstract(transport *Transport) *ResourceAbstract {
	return &ResourceAbstract{tr: transport}
}
