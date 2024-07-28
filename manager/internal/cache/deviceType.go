package cache

import (
	"context"
	"strings"
	"time"

	"manager/internal/model"

	"github.com/zhufuyi/sponge/pkg/cache"
	"github.com/zhufuyi/sponge/pkg/encoding"
	"github.com/zhufuyi/sponge/pkg/utils"
)

const (
	// cache prefix key, must end with a colon
	deviceTypeCachePrefixKey = "deviceType:"
	// DeviceTypeExpireTime expire time
	DeviceTypeExpireTime = 5 * time.Minute
)

var _ DeviceTypeCache = (*deviceTypeCache)(nil)

// DeviceTypeCache cache interface
type DeviceTypeCache interface {
	Set(ctx context.Context, id uint64, data *model.DeviceType, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.DeviceType, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.DeviceType, error)
	MultiSet(ctx context.Context, data []*model.DeviceType, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// deviceTypeCache define a cache struct
type deviceTypeCache struct {
	cache cache.Cache
}

// NewDeviceTypeCache new a cache
func NewDeviceTypeCache(cacheType *model.CacheType) DeviceTypeCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.DeviceType{}
		})
		return &deviceTypeCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.DeviceType{}
		})
		return &deviceTypeCache{cache: c}
	}

	return nil // no cache
}

// GetDeviceTypeCacheKey cache key
func (c *deviceTypeCache) GetDeviceTypeCacheKey(id uint64) string {
	return deviceTypeCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *deviceTypeCache) Set(ctx context.Context, id uint64, data *model.DeviceType, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetDeviceTypeCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *deviceTypeCache) Get(ctx context.Context, id uint64) (*model.DeviceType, error) {
	var data *model.DeviceType
	cacheKey := c.GetDeviceTypeCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *deviceTypeCache) MultiSet(ctx context.Context, data []*model.DeviceType, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetDeviceTypeCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *deviceTypeCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.DeviceType, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetDeviceTypeCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.DeviceType)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.DeviceType)
	for _, id := range ids {
		val, ok := itemMap[c.GetDeviceTypeCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *deviceTypeCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetDeviceTypeCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *deviceTypeCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetDeviceTypeCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
