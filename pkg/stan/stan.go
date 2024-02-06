package stan

import (
	"fmt"

	"github.com/mortum5/order-service/internal/config"
	"github.com/nats-io/stan.go"
)

func New(config config.Config) (stan.Conn, error) {
	sc, err := stan.Connect(
		config.ClusterID,
		config.ClientID,
		stan.NatsURL(config.NatsURL),
	)

	if err != nil {
		return nil, fmt.Errorf("stan connect failed: %v", err)
	}

	return sc, nil
}
