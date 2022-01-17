package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Server   Server
	Settings Settings
}

type Server struct {
	Port string `envconfig:"PORT" default:"63785"`
	Host string `envconfig:"HOST" default:"0.0.0.0"`
}

type Settings struct {
	QueueUrl string `envconfig:"QUEUE_URL"`
}

func Load() (Config, error) {
	err := LoadEnvironment()
	cfg := Config{}
	err = envconfig.Process("settings", &cfg)
	return cfg, err
}
