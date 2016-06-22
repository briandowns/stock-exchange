package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/briandowns/stock-exchange/models"
)

// RedisCache
type RedisCache struct {
	sync.Locker // synchronize access to this data
	*redis.Pool
}

// NewRedisCache
func NewRedisCache() *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.99.100:6379", // THIS HAS TO BE CHANGED!
		Password: "",
		DB:       0,
	})
	return &RedisCache{
		&sync.Mutex{},
		client,
	}
}

// Build will build the symbol cache
func (r *RedisCache) Build() error {
	log.Print("Building symbol cache...")
	r.Lock()
	defer r.Unlock()

	cache, err := generateSymbolData()
	if err != nil {
		return err
	}
	for _, symbol := range cache {
		b, err := json.Marshal(symbol)
		if err != nil {
			return err
		}
		log.Println(string(b))
	}
	log.Println("Building symbol cache complete!")
	return nil
}

// Get gets the value from the cache
func (r *RedisCache) Get(key []byte) (models.Company, error) {
	r.Lock()
	defer r.Unlock()
	var company models.Company
	x := r.Get(key)
	if err != nil {
		return models.Company{}, nil
	}
	decoder := json.NewDecoder(strCMD)
	if err := decoder.Decode(&company); err != nil {
		return models.Company{}, err
	}
	return models.Company{}, nil
}

// Add will add a given entry to the cache
func (r *RedisCache) Add(key []byte, value models.Company) error {
	r.Lock()
	defer r.Unlock()
	// 0 value in duration position means it won't expire
	return r.Set(string(key), value, 0).Err()
}

// Entries will retrieve all entries in the cache
func (r *RedisCache) Entries() ([]models.Company, error) {
	return nil, nil
}
