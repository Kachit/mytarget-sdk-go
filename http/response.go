package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	raw *http.Response
}

func (r *Response) IsSuccess() bool {
	return r.raw.StatusCode < http.StatusMultipleChoices
}

func (r *Response) GetRawResponse() *http.Response {
	return r.raw
}

func (r *Response) GetRawBody() string {
	body, _ := r.ReadBody()
	return string(body)
}

func (r *Response) Unmarshal(v interface{}) error {
	data, err := r.ReadBody()
	if err != nil {
		return fmt.Errorf("Response@Unmarshal read body: %v", err)
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		return fmt.Errorf("Response@Unmarshal parse json error: %v", err)
	}
	return nil
}

func (r *Response) ReadBody() ([]byte, error) {
	defer r.raw.Body.Close()
	return ioutil.ReadAll(r.raw.Body)
}

func (r *Response) GetError() (*ErrorResult, error) {
	var result ErrorResult
	err := r.Unmarshal(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func NewResponse(raw *http.Response) *Response {
	return &Response{raw: raw}
}

type ErrorResult struct {
	Error *ErrorResultError `json:"error"`
}

type ErrorResultError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
