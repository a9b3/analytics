package config

import (
	"github.com/spf13/viper"
)

// Config contains app cfg variables
type Config struct {
	MONGO_URI     string
	MONGO_DB_NAME string
	PORT          string
	APP_ENV       string
	AUTH_HOST     string
	JWT_SECRET    string
}

// New returns Config
func New() (*Config, error) {
	viper.SetDefault("MONGO_URI", "localhost:27032")
	viper.SetDefault("MONGO_DB_NAME", "analytics_local")
	viper.SetDefault("PORT", "9090")
	viper.SetDefault("APP_ENV", "dev")
	viper.SetDefault("AUTH_HOST", "localhost:9091")

	// viper.AddConfigPath(".")
	// if err := viper.ReadInConfig(); err != nil {
	// 	return nil, err
	// }

	viper.AutomaticEnv()

	cfg := new(Config)
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
