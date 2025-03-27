package config

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Port string
	}

	Database struct {
		Path string
	}
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")

	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
