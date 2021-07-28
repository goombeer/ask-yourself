package config

import (
	"fmt"

	"github.com/alexsasharegan/dotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Environment Environment `default:"local" envconfig:"environment"`
	Host        string      `envconfig:"host"`
	Port        string      `envconfig:"port"`
}

func (s ServerConfig) IsProd() bool {
	return s.Environment == "production"
}

func (s ServerConfig) IsLocal() bool {
	return s.Environment == "local"
}

func LoadConfigWithDotenv() Config {
	err := dotenv.Load(fmt.Sprintf("%s/.env", PrjRoot))
	if err != nil {
		panic(err)
	}
	return LoadConfig()
}

func LoadConfig() Config {
	var conf Config
	err := envconfig.Process("", &conf)
	if err != nil {
		panic(err)
	}
	return conf
}