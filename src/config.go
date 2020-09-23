package mytarget_sdk

const URI_PROD = "https://target.my.com"

type Config struct {
	Uri   string
	Token string
}

func NewConfig() *Config {
	return &Config{Uri: URI_PROD}
}
