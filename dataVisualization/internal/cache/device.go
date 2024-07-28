package cache

import (
	"context"
	"dataVisualization/internal/model"
	"strings"
	"time"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"
)

const (
	// cache prefix key, must end with a colon
	deviceCachePrefixKey = "device:"
	// DeviceExpireTime expire time
	DeviceExpireTime = 5 * time.Minute
)

var _ DeviceCache = (*deviceCache)(nil)

// DeviceCache cache interface
type DeviceCache interface {
	Set(ctx context.Context, id uint64, data *model.Device, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Device, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Device, error)
	MultiSet(ctx context.Context, data []*model.Device, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// deviceCache define a cache struct
type deviceCache struct {
	cache cache.Cache
}

// NewDeviceCache new a cache
func NewDeviceCache(cacheType *model.CacheType) DeviceCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Device{}
		})
		return &deviceCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Device{}
		})
		return &deviceCache{cache: c}
	}

	return nil // no cache
}

// GetDeviceCacheKey cache key
func (c *deviceCache) GetDeviceCacheKey(id uint64) string {
	return deviceCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *deviceCache) Set(ctx context.Context, id uint64, data *model.Device, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetDeviceCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *deviceCache) Get(ctx context.Context, id uint64) (*model.Device, error) {
	var data *model.Device
	cacheKey := c.GetDeviceCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *deviceCache) MultiSet(ctx context.Context, data []*model.Device, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetDeviceCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *deviceCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Device, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetDeviceCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Device)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Device)
	for _, id := range ids {
		val, ok := itemMap[c.GetDeviceCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *deviceCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetDeviceCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *deviceCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetDeviceCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
