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
	ErrReadConfig = errors.New("Error read config")
)

func NewConfig() (*Config, error) {
	var cfg Config

	_ = godotenv.Load()

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, ErrReadConfig
	}

	return &cfg, nil
}
