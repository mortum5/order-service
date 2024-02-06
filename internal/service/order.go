package service

import (
	"github.com/mortum5/order-service/internal/config"
	"github.com/mortum5/order-service/internal/consumer"
	"github.com/mortum5/order-service/internal/entity"
)

const (
	LIMIT int64 = 100
)

type Repository interface {
	SaveOrder(entity.Order) error
	GetOrder(string) (entity.Order, error)
}

type OrderService struct {
	Config     config.Config
	consumer   consumer.NatsConsumer
	repository Repository
}

func New(config config.Config, cons consumer.NatsConsumer, repo Repository) OrderService {
	return OrderService{
		Config:     config,
		consumer:   cons,
		repository: repo,
	}
}

// Take all objects from cosumer and save it to persistent storage.
func (s OrderService) Run() {
	go func() {
		for order := range s.consumer.Update() {
			s.repository.SaveOrder(order)
		}
	}()
}

// Return object from persistent storage or cache if enabled.
func (s OrderService) Get(id string) (entity.Order, error) {
	order, err := s.repository.GetOrder(id)
	return order, err
}
