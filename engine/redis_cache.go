package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/briandowns/stock-exchange/models"

	"github.com/garyburd/redigo/redis"
)

// RedisCache
type RedisCache struct {
	sync.Locker // synchronize access to this data
	*redis.Pool
}

// NewRedisCache
func NewRedisCache() *RedisCache {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "192.168.99.100:6379")
			if err != nil {
				return nil, err
			}
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
	}
}

// Build will build the symbol cache
func (r *RedisCache) Build() error {
	log.Print("Building symbol cache...")
	r.Lock()
	defer r.Unlock()

	c := r.Pool.Get()
	defer c.Close()

	cache, err := generateSymbolData()
	if err != nil {
		return err
	}

	// flush the cache before loading new data
	if err := r.Flush(); err != nil {
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

	var company models.Company
	result, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		return models.Company{}, nil
	}

	decoder := json.NewDecoder(bytes.NewReader(result))
	if err := decoder.Decode(&company); err != nil {
		log.Println(err)
		return models.Company{}, err
	}

	return models.Company{}, nil
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

	keys, err := c.Do("GET", "KEYS")
	if err != nil {
		return nil, err
	}

	fmt.Println(keys)
	return nil, nil
}

// Flush will flush all of the keys in the database
func (r *RedisCache) Flush() error {
	r.Lock()
	defer r.Unlock()

	c := r.Pool.Get()
	defer c.Close()

	_, err := c.Do("FLUSHDB", 0)
	if err != nil {
		return nil
	}
	return nil
}
