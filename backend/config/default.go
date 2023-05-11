package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configurations of the application
// The values are read by viper from a config file or env variables
type Config struct {
	DBDriver string `mapstructure:"POSTGRES_DRIVER"`
	DBSource string `mapstructure:"POSTGRES_SOURCE"`
	Port     string `mapstructure:"PORT"`
	Origin   string `mapstructure:"ORIGIN"`
	Env      string `mapstructure:"ENV"`

	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	if config.Env != "production" {
		// load env vars from app.env
		viper.AddConfigPath(path)
		viper.SetConfigType("env")
		viper.SetConfigName(".local")
	}

	// if any, override .local.env vars with os env vars
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	return
}
