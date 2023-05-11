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
	fmt.Println(ginMode)
	if ginMode == "release" {
		// Use environment variables from OS
		viper.AutomaticEnv()
	} else {
		// Use config file and override with OS environment variables
		viper.AddConfigPath(path)
		viper.SetConfigType("env")
		viper.SetConfigName(".local")

		if err = viper.ReadInConfig(); err != nil {
			return
		}

		viper.AutomaticEnv()
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	return
}
