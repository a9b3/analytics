package config

import "os"

// Config contains app config variables
type Config struct {
	DB_URI  string
	DB_NAME string
	PORT    string
	APP_ENV string
}

// New returns Config
func New() Config {
	return Config{
		DB_URI:  getenvOrDefault("DB_URI", "localhost"),
		DB_NAME: getenvOrDefault("DB_NAME", "analytics_local"),
		PORT:    getenvOrDefault("PORT", "9090"),
		APP_ENV: getenvOrDefault("APP_ENV", "dev"),
	}
}

func getenvOrDefault(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
