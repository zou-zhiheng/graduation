package cache

import (
	"testing"
	"time"

	"manager/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newCompanyUserCache() *gotest.Cache {
	record1 := &model.CompanyUser{}
	record1.ID = 1
	record2 := &model.CompanyUser{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewCompanyUserCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_companyUserCache_Set(t *testing.T) {
	c := newCompanyUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CompanyUser)
	err := c.ICache.(CompanyUserCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(CompanyUserCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_companyUserCache_Get(t *testing.T) {
	c := newCompanyUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CompanyUser)
	err := c.ICache.(CompanyUserCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CompanyUserCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(CompanyUserCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_companyUserCache_MultiGet(t *testing.T) {
	c := newCompanyUserCache()
	defer c.Close()

	var testData []*model.CompanyUser
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.CompanyUser))
	}

	err := c.ICache.(CompanyUserCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CompanyUserCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.CompanyUser))
	}
}

func Test_companyUserCache_MultiSet(t *testing.T) {
	c := newCompanyUserCache()
	defer c.Close()

	var testData []*model.CompanyUser
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.CompanyUser))
	}

	err := c.ICache.(CompanyUserCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_companyUserCache_Del(t *testing.T) {
	c := newCompanyUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CompanyUser)
	err := c.ICache.(CompanyUserCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_companyUserCache_SetCacheWithNotFound(t *testing.T) {
	c := newCompanyUserCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.CompanyUser)
	err := c.ICache.(CompanyUserCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewCompanyUserCache(t *testing.T) {
	c := NewCompanyUserCache(&model.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewCompanyUserCache(&model.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewCompanyUserCache(&model.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
