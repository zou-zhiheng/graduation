package cache

import (
	"testing"
	"time"

	"manager/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newDeviceCache() *gotest.Cache {
	record1 := &model.Device{}
	record1.ID = 1
	record2 := &model.Device{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewDeviceCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_deviceCache_Set(t *testing.T) {
	c := newDeviceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Device)
	err := c.ICache.(DeviceCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(DeviceCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_deviceCache_Get(t *testing.T) {
	c := newDeviceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Device)
	err := c.ICache.(DeviceCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(DeviceCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(DeviceCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_deviceCache_MultiGet(t *testing.T) {
	c := newDeviceCache()
	defer c.Close()

	var testData []*model.Device
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Device))
	}

	err := c.ICache.(DeviceCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(DeviceCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Device))
	}
}

func Test_deviceCache_MultiSet(t *testing.T) {
	c := newDeviceCache()
	defer c.Close()

	var testData []*model.Device
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Device))
	}

	err := c.ICache.(DeviceCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_deviceCache_Del(t *testing.T) {
	c := newDeviceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Device)
	err := c.ICache.(DeviceCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_deviceCache_SetCacheWithNotFound(t *testing.T) {
	c := newDeviceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Device)
	err := c.ICache.(DeviceCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewDeviceCache(t *testing.T) {
	c := NewDeviceCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewDeviceCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewDeviceCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
