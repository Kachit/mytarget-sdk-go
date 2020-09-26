package mytarget_sdk

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func buildStubConfig() *Config {
	cfg := &Config{
		Uri:         "https://github.com",
		AccessToken: "qwerty",
	}
	return cfg
}

func buildStubClient() *Client {
	return NewClient(buildStubConfig(), nil)
}

func loadStubResponseData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func buildStubResponseFromString(statusCode int, json string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(json))
	return &http.Response{Body: body, StatusCode: statusCode}
}

func buildStubResponseFromData(statusCode int, path string) *http.Response {
	data, _ := loadStubResponseData(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode}
}
