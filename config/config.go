package config

import (
	"errors"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	HttpConfig HttpConfig
}

type HttpConfig struct {
	Address string `env:"HTTP_ADDR"`
}

var (
	ErrLoadingConfig = errors.New("Error loading .env file")
	ErrReadConfig    = errors.New("Error read config")
)

func NewConfig() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		return nil, ErrLoadingConfig
	}

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, ErrReadConfig
	}

	return &cfg, nil
}
