package mytarget_sdk

import (
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

func buildStubResponseData() string {
	return `{"result":true, "data":{}, "meta":{}, "errors":[]}`
}

func buildStubResponse(json string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(json))
	return &http.Response{Body: body}
}
