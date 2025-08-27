package config

import (
	"os"
)

type Config struct {
	Port string
}

var Cfg *Config

func Init() {
	Cfg = &Config{
		Port: os.Getenv("PORT"),
	}

	if Cfg.Port == "" {
		Cfg.Port = "8080"
	}
}
