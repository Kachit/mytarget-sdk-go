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

func buildStubClient() *Client {
	return NewClient(buildStubConfig(), nil)
}

func loadStubResponseData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func buildStubResponse(json string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(json))
	return &http.Response{Body: body}
}
