package config

import (
	"github.com/spf13/viper"
)

// Config contains app cfg variables
type Config struct {
	DB_URI    string
	DB_NAME   string
	PORT      string
	APP_ENV   string
	AUTH_HOST string
}

// New returns Config
func New() (*Config, error) {
	viper.SetDefault("DB_URI", "localhost:27032")
	viper.SetDefault("DB_NAME", "analytics_local")
	viper.SetDefault("PORT", "9090")
	viper.SetDefault("APP_ENV", "dev")
	viper.SetDefault("AUTH_HOST", "localhost:9091")

	viper.AutomaticEnv()

	cfg := new(Config)
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
