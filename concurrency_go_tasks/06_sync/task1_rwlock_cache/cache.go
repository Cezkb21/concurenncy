package cache

import "sync"

// Cache представляет потокобезопасный кэш.
type Cache struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

// New создаёт новый кэш.
func New() *Cache {
	return &Cache{data: make(map[string]interface{})}
}

// Set сохраняет значение по ключу.
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

// Get возвращает значение по ключу и признак его наличия.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	val, ok := c.data[key]
	c.mu.RUnlock()
	if ok {
		return val, true
	}
	return nil, false
}
