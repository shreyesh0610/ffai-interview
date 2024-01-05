package utils

import (
	"sync"
)

type Cache struct {
	sync.RWMutex
	IsWordValid map[string]bool
}

var cache *Cache

func init() {
	cache = &Cache{
		IsWordValid: make(map[string]bool),
	}
}

func GetCache() *Cache {
	return cache
}

func (c *Cache) Get(word string) (is_valid bool, found bool) {
	c.RLock()
	defer c.RUnlock()
	is_valid, found = c.IsWordValid[word]
	return is_valid, found
}

func (c *Cache) Set(word string, is_valid bool) {
	c.Lock()
	defer c.Unlock()
	c.IsWordValid[word] = is_valid
}
