package cache

import (
	"context"
	model2 "gateway/internal/model"
	"strings"
	"time"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"
)

const (
	// cache prefix key, must end with a colon
	apiCachePrefixKey = "api:"
	// ApiExpireTime expire time
	ApiExpireTime = 5 * time.Minute
)

var _ ApiCache = (*apiCache)(nil)

// ApiCache cache interface
type ApiCache interface {
	Set(ctx context.Context, id uint64, data *model2.Api, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model2.Api, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model2.Api, error)
	MultiSet(ctx context.Context, data []*model2.Api, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// apiCache define a cache struct
type apiCache struct {
	cache cache.Cache
}

// NewApiCache new a cache
func NewApiCache(cacheType *model2.CacheType) ApiCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model2.Api{}
		})
		return &apiCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model2.Api{}
		})
		return &apiCache{cache: c}
	}

	return nil // no cache
}

// GetApiCacheKey cache key
func (c *apiCache) GetApiCacheKey(id uint64) string {
	return apiCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *apiCache) Set(ctx context.Context, id uint64, data *model2.Api, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetApiCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *apiCache) Get(ctx context.Context, id uint64) (*model2.Api, error) {
	var data *model2.Api
	cacheKey := c.GetApiCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *apiCache) MultiSet(ctx context.Context, data []*model2.Api, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetApiCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *apiCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model2.Api, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetApiCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model2.Api)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model2.Api)
	for _, id := range ids {
		val, ok := itemMap[c.GetApiCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *apiCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetApiCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *apiCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetApiCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
