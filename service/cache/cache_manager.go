package cache

import (
	"sync"
)

// cache管理接口
type ICacheManager interface {
	GetDefaultCache() ICache

	GetCache(key string) ICache
	SetCache(key string, cache ICache)
	RemoveCache(key string)
}

type cacheManager struct {
	defaultCache ICache

	rwLock        sync.RWMutex
	registedCache map[string]ICache
}

func NewCacheManager() ICacheManager {
	cacheManager := &cacheManager{
		defaultCache:  NewCache(),
		registedCache: make(map[string]ICache),
	}
	return cacheManager
}

func (c *cacheManager) GetDefaultCache() ICache {
	return c.defaultCache
}

func (c *cacheManager) GetCache(key string) ICache {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	return c.registedCache[key]
}

func (c *cacheManager) SetCache(key string, cache ICache) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	c.registedCache[key] = cache
}

func (c *cacheManager) RemoveCache(key string) {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	delete(c.registedCache, key)
}
