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
	companyUserCachePrefixKey = "companyUser:"
	// CompanyUserExpireTime expire time
	CompanyUserExpireTime = 5 * time.Minute
)

var _ CompanyUserCache = (*companyUserCache)(nil)

// CompanyUserCache cache interface
type CompanyUserCache interface {
	Set(ctx context.Context, id uint64, data *model.CompanyUser, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.CompanyUser, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.CompanyUser, error)
	MultiSet(ctx context.Context, data []*model.CompanyUser, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetCacheWithNotFound(ctx context.Context, id uint64) error
}

// companyUserCache define a cache struct
type companyUserCache struct {
	cache cache.Cache
}

// NewCompanyUserCache new a cache
func NewCompanyUserCache(cacheType *model.CacheType) CompanyUserCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.CompanyUser{}
		})
		return &companyUserCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.CompanyUser{}
		})
		return &companyUserCache{cache: c}
	}

	return nil // no cache
}

// GetCompanyUserCacheKey cache key
func (c *companyUserCache) GetCompanyUserCacheKey(id uint64) string {
	return companyUserCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *companyUserCache) Set(ctx context.Context, id uint64, data *model.CompanyUser, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetCompanyUserCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *companyUserCache) Get(ctx context.Context, id uint64) (*model.CompanyUser, error) {
	var data *model.CompanyUser
	cacheKey := c.GetCompanyUserCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *companyUserCache) MultiSet(ctx context.Context, data []*model.CompanyUser, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetCompanyUserCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *companyUserCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.CompanyUser, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetCompanyUserCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.CompanyUser)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.CompanyUser)
	for _, id := range ids {
		val, ok := itemMap[c.GetCompanyUserCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *companyUserCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetCompanyUserCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheWithNotFound set empty cache
func (c *companyUserCache) SetCacheWithNotFound(ctx context.Context, id uint64) error {
	cacheKey := c.GetCompanyUserCacheKey(id)
	err := c.cache.SetCacheWithNotFound(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}
