package main

import (
	"github.com/canter-tech/car-service/config"
	"github.com/canter-tech/car-service/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
