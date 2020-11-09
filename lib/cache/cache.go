package cache

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Cache describe cache behaviours
type Cache interface {
	Put(key string, value interface{}) error
	PutWithExpire(key string, value interface{}, ttl int64) error
	GetValue(key string) ([]byte, error)
	IsAvailable(key string) bool
	Delete(key string) error
}

// NMCache implements Cache interface
type NMCache struct {
	Duration int64
	Data     map[string]NMCacheData
	Mu       sync.RWMutex
}

// NMCacheData store cache data with last updated time
type NMCacheData struct {
	Time  int64
	Value []byte
}

// NewNMCache initialize new struct netmonkcache
func NewNMCache(duration int) Cache {
	return &NMCache{
		Duration: int64(duration),
		Data:     make(map[string]NMCacheData),
	}
}

// PutWithExpire stores value to cache with specified key in specified time
func (c *NMCache) PutWithExpire(key string, value interface{}, ttl int64) error {
	result, err := json.Marshal(value)
	if err != nil {
		return err
	}

	cacheData := NMCacheData{
		Time:  time.Now().Unix(),
		Value: result,
	}
	c.Mu.Lock()
	c.Data[key] = cacheData
	c.Mu.Unlock()
	return nil
}

// Put stores value to cache with specified key
func (c *NMCache) Put(key string, value interface{}) error {
	result, err := json.Marshal(value)
	if err != nil {
		return err
	}

	cacheData := NMCacheData{
		Time:  time.Now().Unix(),
		Value: result,
	}
	c.Mu.Lock()
	c.Data[key] = cacheData
	c.Mu.Unlock()
	return nil
}

// IsAvailable checks if the cache with specified key within duration time is available
func (c *NMCache) IsAvailable(key string) bool {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	if _, ok := c.Data[key]; ok {
		lastUpdated, _ := c.GetLastUpdated(key)
		interval := time.Now().Unix() - lastUpdated
		if interval <= c.GetDuration() {
			return true
		}
	}

	return false
}

// GetValue fetch data value from cache with specified key
func (c *NMCache) GetValue(key string) ([]byte, error) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	if v, ok := c.Data[key]; ok {
		return v.Value, nil
	}

	return nil, fmt.Errorf("Cache key not found")
}

// GetLastUpdated fetch time value from cache with specified key
func (c *NMCache) GetLastUpdated(key string) (int64, error) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	if v, ok := c.Data[key]; ok {
		return v.Time, nil
	}

	return 0, fmt.Errorf("Cache key not found")
}

// GetDuration get cache duration time
func (c *NMCache) GetDuration() int64 {
	return c.Duration
}

// Delete removes data from cache with specified key
func (c *NMCache) Delete(key string) error {
	c.Mu.Lock()
	delete(c.Data, key)
	c.Mu.Unlock()
	return nil
}
