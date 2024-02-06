package repository

import (
	"context"
	"encoding/json"

	"github.com/mortum5/order-service/internal/entity"
	sqlc "github.com/mortum5/order-service/pkg/sqlc/generate"
)

type OrderRepository struct {
	queries *sqlc.Queries
}

func New(queries *sqlc.Queries) OrderRepository {
	return OrderRepository{
		queries: queries,
	}
}

func (repo OrderRepository) SaveOrder(order entity.Order) (err error) {
	json, err := json.Marshal(order)
	if err != nil {
		return
	}

	params := sqlc.CreateOrderParams{
		ID:   order.OrderUID,
		Data: json,
	}
	_, err = repo.queries.CreateOrder(context.Background(), params)
	return err
}

func (repo OrderRepository) GetOrder(id string) (order entity.Order, err error) {
	sqlcOrder, err := repo.queries.GetOrder(context.Background(), id)
	if err != nil {
		return
	}

	err = json.Unmarshal(sqlcOrder.Data, &order)
	if err != nil {
		return
	}

	return order, nil
}

func (repo OrderRepository) ListOrders(limit int64) (orders []entity.Order, err error) {
	sqlcOrders, err := repo.queries.ListOrders(context.Background(), int32(limit))
	if err != nil {
		return
	}

	var order entity.Order

	for _, sqlcOrd := range sqlcOrders {
		err = json.Unmarshal(sqlcOrd.Data, &order)
		if err != nil {
			return
		}
		orders = append(orders, order)
	}
	return
}
