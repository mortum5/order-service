package main

import (
	"log/slog"

	_ "github.com/mortum5/order-service/docs"
	"github.com/mortum5/order-service/internal/app"
	"github.com/mortum5/order-service/internal/config"
	"github.com/mortum5/order-service/pkg/logger"
)

// @title Order Service Application
// @version 0.1
// @host localhost:8080
// @BasePath /.
func main() {
	logger := logger.New()
	slog.SetDefault(logger)

	config, err := config.LoadConfig("config/")
	if err != nil {
		panic(err)
	}

	app.Run(config)
}
