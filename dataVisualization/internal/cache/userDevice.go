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
	userDeviceCachePrefixKey = "userDevice:"
	// UserDeviceExpireTime expire time
	UserDeviceExpireTime = 5 * time.Minute
)

var _ UserDeviceCache = (*userDeviceCache)(nil)

// UserDeviceCache cache interface
type UserDeviceCache interface {
	Set(ctx context.Context, id uint64, data *model.UserDevice, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.UserDevice, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserDevice, error)
	MultiSet(ctx context.Context, data []*model.UserDevice, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// userDeviceCache define a cache struct
type userDeviceCache struct {
	cache cache.Cache
}

// NewUserDeviceCache new a cache
func NewUserDeviceCache(cacheType *model.CacheType) UserDeviceCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserDevice{}
		})
		return &userDeviceCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.UserDevice{}
		})
		return &userDeviceCache{cache: c}
	}

	return nil // no cache
}

// GetUserDeviceCacheKey cache key
func (c *userDeviceCache) GetUserDeviceCacheKey(id uint64) string {
	return userDeviceCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *userDeviceCache) Set(ctx context.Context, id uint64, data *model.UserDevice, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetUserDeviceCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *userDeviceCache) Get(ctx context.Context, id uint64) (*model.UserDevice, error) {
	var data *model.UserDevice
	cacheKey := c.GetUserDeviceCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *userDeviceCache) MultiSet(ctx context.Context, data []*model.UserDevice, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetUserDeviceCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *userDeviceCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.UserDevice, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetUserDeviceCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.UserDevice)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.UserDevice)
	for _, id := range ids {
		val, ok := itemMap[c.GetUserDeviceCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *userDeviceCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserDeviceCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *userDeviceCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetUserDeviceCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
