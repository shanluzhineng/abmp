package cache

import (
	"time"

	"github.com/shanluzhineng/abmp/pkg/core/cache"
)

type CacheOption cache.Option[string, any]

// 设置默认的过期时间
func WithTTL(ttl time.Duration) CacheOption {
	return cache.WithTTL[string, any](ttl)
}

// 设置当key不存在时，告诉缓存如何加载,所有的key都使用同一个回调
func WithLoader(loaderFunc func(key string) (any, time.Duration)) CacheOption {
	funcLoader := cache.LoaderFunc[string, any](
		func(internalCache *cache.Cache[string, any], internalKey string) *cache.Item[string, any] {
			//调用缓存
			newValue, ttl := loaderFunc(internalKey)
			cacheItemValue := internalCache.Set(internalKey, newValue, ttl)
			return cacheItemValue
		},
	)
	loader := &cache.SuppressedLoader[string, any]{
		Loader: funcLoader,
	}
	return cache.WithLoader[string, any](loader)
}

// cache接口
type ICache interface {
	Get(key string) any
	//设置key的值，没有有效期
	Set(key string, value any) any
	//设置key的值，有有效期
	SetWithTTL(key string, value any, ttl time.Duration) any
	Delete(key string)
	DeleteAll()

	Len() int
	Keys() []string
	Items() map[string]any
}

type defaultCache struct {
	*cache.Cache[string, any]
}

// new cache
func NewCache(opts ...CacheOption) ICache {
	cacheOpts := make([]cache.Option[string, any], len(opts))
	for i, eachOpt := range opts {
		cacheOpts[i] = eachOpt
	}

	newCache := &defaultCache{
		Cache: cache.New(cacheOpts...),
	}
	go newCache.Start()
	return newCache
}

func (c *defaultCache) Get(key string) any {
	cValue := c.Cache.Get(key)
	if cValue == nil {
		return nil
	}
	return cValue.Value()
}

func (c *defaultCache) Set(key string, value any) any {
	return c.SetWithTTL(key, value, cache.NoTTL)
}

// 设置值
func (c *defaultCache) SetWithTTL(key string, value any, ttl time.Duration) any {
	cValue := c.Cache.Set(key, value, ttl)
	return cValue.Value()
}

func (c *defaultCache) Delete(key string) {
	c.Cache.Delete(key)
}

func (c *defaultCache) DeleteAll() {
	c.Cache.DeleteAll()
}

func (c *defaultCache) Len() int {
	return c.Cache.Len()
}

func (c *defaultCache) Keys() []string {
	return c.Cache.Keys()
}

func (c *defaultCache) Items() map[string]any {
	cValues := c.Cache.Items()
	values := make(map[string]any)
	for k, v := range cValues {
		values[k] = v.Value()
	}
	return values
}
