package resources

import (
	"fmt"
	mytarget_http "github.com/kachit/mytarget-sdk-go/http"
)

type ResourceAbstract struct {
	tr *mytarget_http.Transport
}

func (r *ResourceAbstract) Get(path string, query map[string]interface{}) (*mytarget_http.Response, error) {
	rsp, err := r.tr.Get(path, query)
	if err != nil {
		return nil, fmt.Errorf("ResourceAbstract@get request: %v", err)
	}
	return mytarget_http.NewResponse(rsp), nil
}

func (r *ResourceAbstract) Post(path string, body map[string]interface{}, query map[string]interface{}) (*mytarget_http.Response, error) {
	rsp, err := r.tr.Post(path, body, query)
	if err != nil {
		return nil, fmt.Errorf("ResourceAbstract@post request: %v", err)
	}
	return mytarget_http.NewResponse(rsp), nil
}

func NewResourceAbstract(transport *mytarget_http.Transport) *ResourceAbstract {
	return &ResourceAbstract{tr: transport}
}
