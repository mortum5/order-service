package main

import (
	"encoding/json"
	"flag"
	"log/slog"

	"github.com/mortum5/order-service/pkg/utils"

	"github.com/nats-io/stan.go"
)

const (
	clusterID string = "test-cluster"
	clientID  string = "test-producer"
)

func main() {
	var n = flag.Int("c", 10, "count of random order send into nats")
	flag.Parse()

	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL("nats://nats:4222"),
	)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := sc.Close()
		if err != nil {
			panic(err)
		}
	}()

	for i := 0; i < *n; i++ {
		order := utils.GenerateNewOrder()
		str, err := json.Marshal(order)
		if err != nil {
			slog.Error("producer: json marshal", err)
		}

		err = sc.Publish("orders", str)
		if err != nil {
			slog.Error("producer: publish", err)
		}
	}
}
