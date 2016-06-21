package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/briandowns/stock-exchange/models"
	"gopkg.in/redis.v3"
)

// RedisCache
type RedisCache struct {
	Lock   sync.Locker // synchronize access to this data
	Client *redis.Client
}

// Build will build the symbol cache
func (r *RedisCache) Build() error {
	log.Print("Building symbol cache...")
	r.Lock.Lock()
	defer r.Lock.Unlock()

	cache, err := generateSymbolData()
	if err != nil {
		return err
	}
	for _, symbol := range cache {
		b, err := json.Marshal(symbol)
		if err != nil {
			return err
		}
		log.Println(b)
	}
	log.Println("Building symbol cache complete!")
	return nil
}

// Get gets the value from the cache
func (r *RedisCache) Get(key []byte) (models.Company, error) {
	return models.Company{}, nil
}

func (r *RedisCache) Add(key []byte, value models.Company) error {
	return nil
}

func (r *RedisCache) Entries() ([]models.Company, error) {
	return nil, nil
}

// NewRedisCache
func NewRedisCache() *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.99.100:6379", // THIS HAS TO BE CHANGED!
		Password: "",
		DB:       0,
	})
	return &RedisCache{
		Lock:   &sync.Mutex{},
		Client: client,
	}
}
