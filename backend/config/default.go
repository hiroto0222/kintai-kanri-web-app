package config

import (
	"github.com/spf13/viper"
)

// Config stores all configurations of the application
// The values are read by viper from a config file or env variables
type Config struct {
	DBDriver string `mapstructure:"POSTGRES_DRIVER"`
	DBSource string `mapstructure:"POSTGRES_SOURCE"`
	Port     string `mapstructure:"PORT"`
	Origin   string `mapstructure:"ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	// load env vars from app.env
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".local")

	// if any, override .local.env vars with os env vars
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	return
}
