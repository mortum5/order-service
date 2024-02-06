package postgres

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mortum5/order-service/internal/config"
)

func New(config config.Config) (*pgxpool.Pool, error) {
	hostAndPort := net.JoinHostPort(config.DBHost, strconv.Itoa(config.DBPort))
	dbSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		config.DBUser,
		config.DBPass,
		hostAndPort,
		config.DBName,
	)

	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		return nil, fmt.Errorf("postres connection error: %v", err)
	}

	return pool, nil
}
