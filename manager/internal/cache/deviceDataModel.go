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
	deviceDataModelCachePrefixKey = "deviceDataModel:"
	// DeviceDataModelExpireTime expire time
	DeviceDataModelExpireTime = 5 * time.Minute
)

var _ DeviceDataModelCache = (*deviceDataModelCache)(nil)

// DeviceDataModelCache cache interface
type DeviceDataModelCache interface {
	Set(ctx context.Context, id uint64, data *model.DeviceDataModel, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.DeviceDataModel, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.DeviceDataModel, error)
	MultiSet(ctx context.Context, data []*model.DeviceDataModel, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// deviceDataModelCache define a cache struct
type deviceDataModelCache struct {
	cache cache.Cache
}

// NewDeviceDataModelCache new a cache
func NewDeviceDataModelCache(cacheType *model.CacheType) DeviceDataModelCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.DeviceDataModel{}
		})
		return &deviceDataModelCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.DeviceDataModel{}
		})
		return &deviceDataModelCache{cache: c}
	}

	return nil // no cache
}

// GetDeviceDataModelCacheKey cache key
func (c *deviceDataModelCache) GetDeviceDataModelCacheKey(id uint64) string {
	return deviceDataModelCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *deviceDataModelCache) Set(ctx context.Context, id uint64, data *model.DeviceDataModel, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetDeviceDataModelCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *deviceDataModelCache) Get(ctx context.Context, id uint64) (*model.DeviceDataModel, error) {
	var data *model.DeviceDataModel
	cacheKey := c.GetDeviceDataModelCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *deviceDataModelCache) MultiSet(ctx context.Context, data []*model.DeviceDataModel, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetDeviceDataModelCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *deviceDataModelCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.DeviceDataModel, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetDeviceDataModelCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.DeviceDataModel)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.DeviceDataModel)
	for _, id := range ids {
		val, ok := itemMap[c.GetDeviceDataModelCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *deviceDataModelCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetDeviceDataModelCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *deviceDataModelCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetDeviceDataModelCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
