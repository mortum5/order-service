package consumer

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/mortum5/order-service/internal/config"
	"github.com/mortum5/order-service/internal/entity"
	"github.com/nats-io/stan.go"
)

type NatsConsumer struct {
	config    config.Config
	sc        stan.Conn
	sub       stan.Subscription
	orderChan chan entity.Order
}

// Create new Nats consumer which will receive all not readed message and send
// them to the result channel.
func New(config config.Config, sc stan.Conn) NatsConsumer {
	return NatsConsumer{
		config:    config,
		sc:        sc,
		orderChan: make(chan entity.Order),
	}
}

// Subscribe to nats streaming.
func (nc *NatsConsumer) Run() error {
	sub, err := nc.sc.Subscribe(
		nc.config.TopicName,
		func(m *stan.Msg) {
			var order entity.Order
			err := json.Unmarshal(m.Data, &order)
			if err != nil {
				slog.Info("consumer: json unmarshal", slog.Any("err", err.Error()))
			} else {
				nc.orderChan <- order
			}
		},
		stan.DurableName(nc.config.DurableName),
		stan.DeliverAllAvailable(),
	)

	nc.sub = sub
	if err != nil {
		return fmt.Errorf("consumer: subscribe %v", err)
	}
	return nil
}

// Closing nats conn and orders channel.
// Closing on durable connect allows us to start from the first unread message.
func (nc *NatsConsumer) Close() (err error) {
	err = nc.sub.Close()
	close(nc.orderChan)

	if err != nil {
		err = fmt.Errorf("consumer: close %v", err)
	}
	return
}

// Return channel with received orders.
func (nc *NatsConsumer) Update() chan entity.Order {
	return nc.orderChan
}
