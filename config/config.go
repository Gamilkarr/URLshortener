package config

type Config struct {
	ServerAddress string
}

func New() *Config {
	return &Config{
		ServerAddress: ":8080",
	}
}
