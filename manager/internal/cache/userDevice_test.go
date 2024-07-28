package cache

import (
	"testing"
	"time"

	"manager/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newUserDeviceCache() *gotest.Cache {
	record1 := &model.UserDevice{}
	record1.ID = 1
	record2 := &model.UserDevice{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewUserDeviceCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_userDeviceCache_Set(t *testing.T) {
	c := newUserDeviceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserDevice)
	err := c.ICache.(UserDeviceCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(UserDeviceCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_userDeviceCache_Get(t *testing.T) {
	c := newUserDeviceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserDevice)
	err := c.ICache.(UserDeviceCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserDeviceCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(UserDeviceCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_userDeviceCache_MultiGet(t *testing.T) {
	c := newUserDeviceCache()
	defer c.Close()

	var testData []*model.UserDevice
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserDevice))
	}

	err := c.ICache.(UserDeviceCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserDeviceCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.UserDevice))
	}
}

func Test_userDeviceCache_MultiSet(t *testing.T) {
	c := newUserDeviceCache()
	defer c.Close()

	var testData []*model.UserDevice
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserDevice))
	}

	err := c.ICache.(UserDeviceCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userDeviceCache_Del(t *testing.T) {
	c := newUserDeviceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserDevice)
	err := c.ICache.(UserDeviceCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userDeviceCache_SetCacheWithNotFound(t *testing.T) {
	c := newUserDeviceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserDevice)
	err := c.ICache.(UserDeviceCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewUserDeviceCache(t *testing.T) {
	c := NewUserDeviceCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewUserDeviceCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewUserDeviceCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
