package utils

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/mortum5/order-service/internal/entity"
)

func GenerateNewOrder() *entity.Order {
	var order entity.Order
	gofakeit.Struct(&order)

	return &order
}
