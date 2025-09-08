package config

import (
	"os"
)

type Config struct {
	Port         string
	JwtSecret    string
	GmailId      string
	GmailAppPass string
	Database
}

type Database struct {
	DB_USER     string
	DB_PWD      string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     string
}

var Cfg *Config

func Init() {
	Cfg = &Config{
		Port:         os.Getenv("PORT"),
		JwtSecret:    os.Getenv("JWT_SECRET"),
		GmailId:      os.Getenv("GMAIL_ID"),
		GmailAppPass: os.Getenv("GMAIL_APP_PASS"),
		Database: Database{
			DB_USER:     os.Getenv("DB_USER"),
			DB_PWD:      os.Getenv("DB_PWD"),
			DB_DATABASE: os.Getenv("DB_DATABASE"),
			DB_HOST:     os.Getenv("DB_HOST"),
			DB_PORT:     os.Getenv("DB_PORT"),
		},
	}

	if Cfg.Port == "" {
		Cfg.Port = "8080"
	}
}
