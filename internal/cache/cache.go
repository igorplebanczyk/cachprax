package cache

import (
	goCache "github.com/patrickmn/go-cache"
	"time"
)

type Cache struct {
	internalCache *goCache.Cache
}

func NewCache(expireAfter time.Duration, purgeAfter time.Duration) *Cache {
	return &Cache{
		internalCache: goCache.New(expireAfter, purgeAfter),
	}
}

func (cache *Cache) IsCached(cacheKey string) bool {
	_, found := cache.internalCache.Get(cacheKey)
	return found
}

func (cache *Cache) GetCached(cacheKey string) []byte {
	data, found := cache.internalCache.Get(cacheKey)
	if !found {
		return nil
	}
	return data.([]byte)
}

func (cache *Cache) SetCached(cacheKey string, data []byte) {
	cache.internalCache.Set(cacheKey, data, goCache.NoExpiration)
}
