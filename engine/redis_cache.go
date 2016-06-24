package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/briandowns/stock-exchange/config"
	"github.com/briandowns/stock-exchange/models"

	"github.com/garyburd/redigo/redis"
)

// RedisCache
type RedisCache struct {
	sync.Locker // synchronize access to this data
	*redis.Pool
	*config.Config
}

// NewRedisCache
func NewRedisCache(config *config.Config) *RedisCache {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", config.Cache.Redis.Host, config.Cache.Redis.Port))
			if err != nil {
				return nil, err
			}
			// leaving this here in case there's auth needed later
			/*if _, err := c.Do("AUTH", ""); err != nil {
				c.Close()
				return nil, err
			}*/
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return &RedisCache{
		&sync.Mutex{},
		pool,
		config,
	}
}

// Build will build the symbol cache
func (r *RedisCache) Build() error {
	log.Print("Building symbol cache...")
	r.Lock()
	defer r.Unlock()

	c := r.Pool.Get()
	defer c.Close()

	cache, err := generateSymbolData("data/" + r.Config.Exchange + "_symbols.json")
	if err != nil {
		return err
	}

	// iterate over the symbol data and add to cache
	for _, symbol := range cache {
		b, err := json.Marshal(symbol)
		if err != nil {
			return err
		}
		_, err = c.Do("SET", symbol.Symbol, b)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	log.Println("Building symbol cache complete!")
	return nil
}

// Get gets the value from the cache
func (r *RedisCache) Get(key []byte) (models.Company, error) {
	r.Lock()
	defer r.Unlock()

	c := r.Pool.Get()
	defer c.Close()

	result, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		return models.Company{}, nil
	}

	var company models.Company
	decoder := json.NewDecoder(bytes.NewReader(result))
	if err := decoder.Decode(&company); err != nil {
		return models.Company{}, err
	}

	return company, nil
}

// Add will add a given entry to the cache
func (r *RedisCache) Add(key []byte, value models.Company) error {
	r.Lock()
	defer r.Unlock()

	c := r.Pool.Get()
	defer c.Close()

	return nil
}

// Entries will retrieve all entries in the cache
func (r *RedisCache) Entries() ([]models.Company, error) {
	c := r.Pool.Get()
	defer c.Close()

	keys, err := redis.Strings(c.Do("KEYS", "*"))
	if err != nil {
		return nil, err
	}

	var symbols []models.Company
	for _, key := range keys {
		result, err := redis.Bytes(c.Do("GET", key))
		if err != nil {
			return nil, nil
		}
		var symbol models.Company
		decoder := json.NewDecoder(bytes.NewReader(result))
		if err := decoder.Decode(&symbol); err != nil {
			return nil, err
		}
		symbols = append(symbols, symbol)
	}

	return symbols, nil
}

// Flush will flush all of the keys in the database
func (r *RedisCache) Flush() error {
	r.Lock()
	defer r.Unlock()

	c := r.Pool.Get()
	defer c.Close()

	_, err := c.Do("FLUSHDB")
	if err != nil {
		return nil
	}
	return nil
}
