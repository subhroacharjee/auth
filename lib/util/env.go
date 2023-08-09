package util

import (
	"os"
	"strings"
	"text/template"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	Timezone string
}

func GetDbDSN() (*string, error) {
	t, err := template.New("DSN").Parse("host={{.Host}} user={{.User}} password={{.Password}} dbname={{.DBName}} port={{.Port}} sslmode={{.SSLMode}} timezone={{.Timezone}}")
	if err != nil {
		return nil, err
	}
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
		Timezone: os.Getenv("DB_TIMEZONE"),
	}

	var dnsBuilder strings.Builder
	err = t.Execute(&dnsBuilder, dbConfig)
	if err != nil {
		return nil, err
	}
	res := dnsBuilder.String()
	return &res, nil
}
