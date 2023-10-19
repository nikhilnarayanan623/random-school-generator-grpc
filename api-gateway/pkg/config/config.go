package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	ApiPort           string `mapstructure:"API_PORT"`
	SchoolServiceHost string `mapstructure:"SCHOOL_SERVICE_HOST"`
	SchoolServicePort string `mapstructure:"SCHOOL_SERVICE_PORT"`
}

var envs = []string{"API_PORT", "SCHOOL_SERVICE_HOST", "SCHOOL_SERVICE_PORT"}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
