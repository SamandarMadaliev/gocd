package config

type Config struct {
	ProjectName string
	ModuleName  string
}

func NewConfig() *Config {
	return &Config{}
}
