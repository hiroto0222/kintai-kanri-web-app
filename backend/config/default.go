package config

import (
	"fmt"
	"os"
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

	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	// Check if running in production mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		// Use environment variables from OS only
		config.DBDriver = os.Getenv("POSTGRES_DRIVER")
		config.DBSource = os.Getenv("POSTGRES_SOURCE")
		config.Port = os.Getenv("PORT")
		config.Origin = os.Getenv("ORIGIN")
		config.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")
		config.AccessTokenDuration, err = time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
		if err != nil {
			return
		}
		config.RefreshTokenDuration, err = time.ParseDuration(os.Getenv("REFRESH_TOKEN_DURATION"))
		if err != nil {
			return
		}
	} else {
		// Use config file and override with OS environment variables
		viper.AddConfigPath(path)
		viper.SetConfigType("env")
		viper.SetConfigName(".local")

		if err = viper.ReadInConfig(); err != nil {
			return
		}

		viper.AutomaticEnv()

		if err = viper.Unmarshal(&config); err != nil {
			return
		}
	}

	fmt.Println(config)

	return
}
