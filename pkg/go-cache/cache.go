package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func New() *cache.Cache {
	return cache.New(5*time.Minute, 10*time.Minute)
}
