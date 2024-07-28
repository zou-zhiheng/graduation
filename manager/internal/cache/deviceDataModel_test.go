package cache

import (
	"testing"
	"time"

	"manager/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newDeviceDataModelCache() *gotest.Cache {
	record1 := &model.DeviceDataModel{}
	record1.ID = 1
	record2 := &model.DeviceDataModel{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewDeviceDataModelCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_deviceDataModelCache_Set(t *testing.T) {
	c := newDeviceDataModelCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.DeviceDataModel)
	err := c.ICache.(DeviceDataModelCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(DeviceDataModelCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_deviceDataModelCache_Get(t *testing.T) {
	c := newDeviceDataModelCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.DeviceDataModel)
	err := c.ICache.(DeviceDataModelCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(DeviceDataModelCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(DeviceDataModelCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_deviceDataModelCache_MultiGet(t *testing.T) {
	c := newDeviceDataModelCache()
	defer c.Close()

	var testData []*model.DeviceDataModel
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.DeviceDataModel))
	}

	err := c.ICache.(DeviceDataModelCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(DeviceDataModelCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.DeviceDataModel))
	}
}

func Test_deviceDataModelCache_MultiSet(t *testing.T) {
	c := newDeviceDataModelCache()
	defer c.Close()

	var testData []*model.DeviceDataModel
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.DeviceDataModel))
	}

	err := c.ICache.(DeviceDataModelCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_deviceDataModelCache_Del(t *testing.T) {
	c := newDeviceDataModelCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.DeviceDataModel)
	err := c.ICache.(DeviceDataModelCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_deviceDataModelCache_SetCacheWithNotFound(t *testing.T) {
	c := newDeviceDataModelCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.DeviceDataModel)
	err := c.ICache.(DeviceDataModelCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewDeviceDataModelCache(t *testing.T) {
	c := NewDeviceDataModelCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewDeviceDataModelCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewDeviceDataModelCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
