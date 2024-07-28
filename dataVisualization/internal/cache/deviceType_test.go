package cache

import (
	"testing"
	"time"

	"manager/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newDeviceTypeCache() *gotest.Cache {
	record1 := &model.DeviceType{}
	record1.ID = 1
	record2 := &model.DeviceType{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewDeviceTypeCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_deviceTypeCache_Set(t *testing.T) {
	c := newDeviceTypeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.DeviceType)
	err := c.ICache.(DeviceTypeCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(DeviceTypeCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_deviceTypeCache_Get(t *testing.T) {
	c := newDeviceTypeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.DeviceType)
	err := c.ICache.(DeviceTypeCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(DeviceTypeCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(DeviceTypeCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_deviceTypeCache_MultiGet(t *testing.T) {
	c := newDeviceTypeCache()
	defer c.Close()

	var testData []*model.DeviceType
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.DeviceType))
	}

	err := c.ICache.(DeviceTypeCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(DeviceTypeCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.DeviceType))
	}
}

func Test_deviceTypeCache_MultiSet(t *testing.T) {
	c := newDeviceTypeCache()
	defer c.Close()

	var testData []*model.DeviceType
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.DeviceType))
	}

	err := c.ICache.(DeviceTypeCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_deviceTypeCache_Del(t *testing.T) {
	c := newDeviceTypeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.DeviceType)
	err := c.ICache.(DeviceTypeCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_deviceTypeCache_SetCacheWithNotFound(t *testing.T) {
	c := newDeviceTypeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.DeviceType)
	err := c.ICache.(DeviceTypeCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewDeviceTypeCache(t *testing.T) {
	c := NewDeviceTypeCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewDeviceTypeCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewDeviceTypeCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
