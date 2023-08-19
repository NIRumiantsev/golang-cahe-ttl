package cache

import (
	"errors"
	"sync"
	"time"
)

type Data map[string]interface{}

type Cache struct {
	data  Data
	mutex *sync.Mutex
}

func (cache *Cache) Set(key string, value interface{}, ttl time.Duration) {
	cache.mutex.Lock()
	cache.data[key] = value
	cache.mutex.Unlock()

	go func() {
		time.Sleep(ttl)
		cache.Delete(key)
	}()
}

func (cache *Cache) Get(key string) (interface{}, error) {
	if cache.Has(key) {
		return cache.data[key], nil
	}
	return nil, errors.New("value does not exist")
}

func (cache *Cache) Delete(key string) {
	delete(cache.data, key)
}

func (cache *Cache) Has(key string) bool {
	return cache.data[key] != nil
}

func New() Cache {
	return Cache{
		data:  make(Data),
		mutex: new(sync.Mutex),
	}
}
