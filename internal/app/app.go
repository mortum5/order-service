package app

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/mortum5/order-service/internal/config"
	"github.com/mortum5/order-service/internal/consumer"
	server "github.com/mortum5/order-service/internal/controller/http"
	cacheRepo "github.com/mortum5/order-service/internal/repository/cache"
	repository "github.com/mortum5/order-service/internal/repository/sqlc"
	"github.com/mortum5/order-service/internal/service"
	"github.com/mortum5/order-service/pkg/go-cache"
	"github.com/mortum5/order-service/pkg/postgres"
	sqlc "github.com/mortum5/order-service/pkg/sqlc/generate"
	"github.com/mortum5/order-service/pkg/stan"
	"github.com/mortum5/order-service/pkg/utils"
)

func Run(config config.Config) {
	// Get postgres connection
	pool, err := postgres.New(config)
	if err != nil {
		utils.Panic(err)
	}
	defer pool.Close()

	// Establish nats streaming connection
	sc, err := stan.New(config)
	if err != nil {
		utils.Panic(err)
	}
	defer sc.Close()

	// Create new querier object from sqlc
	queries := sqlc.New(pool)

	// Create repository for order
	repository := repository.New(queries)

	// Create cache
	cache := cache.New()

	// Create cached version of repository
	repo := cacheRepo.New(repository, cache)

	// Restore cache from persistent storage
	repo.RestoreCacheFromStorage()

	// Create nats consumer
	cons := consumer.New(config, sc)
	err = cons.Run()
	if err != nil {
		utils.Panic(err)
	}
	defer func() {
		err = cons.Close()
		if err != nil {
			utils.Panic(err)
		}
	}()

	// Create order service
	service := service.New(config, cons, repo)

	service.Run()

	// Create http server
	server := server.New(service)

	server.Run()

	// Gracefull shutdown handle
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, os.Interrupt)

	select {
	// Signal interrupt
	case <-exit:
		log.Println("closed by signal")
	case err := <-server.Error():
		slog.Error("server error", slog.Any("err", err.Error()))
	}
}
