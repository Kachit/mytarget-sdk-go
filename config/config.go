package config

const UriProd = "https://target.my.com"

type Config struct {
	Uri         string
	AccessToken string
}

func NewConfig(accessToken string) *Config {
	return &Config{Uri: UriProd, AccessToken: accessToken}
}
