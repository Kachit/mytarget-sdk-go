package stubs

import (
	"bytes"
	"github.com/kachit/mytarget-sdk-go/config"
	"io/ioutil"
	"net/http"
	"strings"
)

func BuildStubConfig() *config.Config {
	cfg := &config.Config{
		Uri:         "https://github.com",
		AccessToken: "qwerty",
	}
	return cfg
}

func LoadStubResponseData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func BuildStubResponseFromString(statusCode int, json string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(json))
	return &http.Response{Body: body, StatusCode: statusCode}
}

func BuildStubResponseFromFile(statusCode int, path string) *http.Response {
	data, _ := LoadStubResponseData(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode}
}
