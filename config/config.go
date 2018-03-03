package config

import "os"

// Config contains app config variables
type Config struct {
	DB_URI    string
	DB_NAME   string
	PORT      string
	APP_ENV   string
	AUTH_HOST string
}

// New returns Config
func New() Config {
	return Config{
		DB_URI:    getenvOrDefault("DB_URI", "localhost:27032"),
		DB_NAME:   getenvOrDefault("DB_NAME", "analytics_local"),
		PORT:      getenvOrDefault("PORT", "9090"),
		APP_ENV:   getenvOrDefault("APP_ENV", "dev"),
		AUTH_HOST: getenvOrDefault("AUTH_HOST", "localhost:9091"),
	}
}

// getenvOrDefault returns fallback if env var doesn't exists
func getenvOrDefault(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
