package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
	CacheURL    string `toml:"cache_url"`
	MsgURL      string `toml:"msg_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		DatabaseURL: "",
	}
}
