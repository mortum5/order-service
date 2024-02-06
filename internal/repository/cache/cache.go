package cache

import (
	"errors"
	"log"

	"github.com/mortum5/order-service/internal/entity"
	"github.com/patrickmn/go-cache"
)

type Repository interface {
	SaveOrder(entity.Order) error
	GetOrder(string) (entity.Order, error)
	ListOrders(int64) ([]entity.Order, error)
}

type OrderRepositoryWithCache struct {
	repo  Repository
	cache *cache.Cache
}

func New(repo Repository, cache *cache.Cache) OrderRepositoryWithCache {
	return OrderRepositoryWithCache{
		repo:  repo,
		cache: cache,
	}
}

func (cr OrderRepositoryWithCache) RestoreCacheFromStorage() {
	go func() {
		orders, _ := cr.repo.ListOrders(100)

		log.Println("cache: start cache recover")
		for _, order := range orders {
			key := order.OrderUID
			cr.cache.Set(key, order, 0)
		}
		log.Printf("cache: %d orders were restored", len(orders))
	}()
}

func (cr OrderRepositoryWithCache) SaveOrder(order entity.Order) (err error) {
	key := order.OrderUID
	if err = cr.cache.Add(key, order, 0); err != nil {
		return
	}
	err = cr.repo.SaveOrder(order)
	return
}

func (cr OrderRepositoryWithCache) GetOrder(id string) (order entity.Order, err error) {
	if obj, ok := cr.cache.Get(id); ok {
		if order, ok = obj.(entity.Order); !ok {
			err = errors.New("cache repo: cannot convert to order type")
		}
		log.Printf("cache hit with id=%s", id)
		return
	}

	order, err = cr.repo.GetOrder(id)
	if err == nil {
		log.Printf("cache miss with id=%s", id)
	}
	return
}
