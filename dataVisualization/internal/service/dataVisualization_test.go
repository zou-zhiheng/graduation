package service

import (
	"context"
	dataVisualizationV1 "dataVisualization/api/dataVisualization/v1"
	"dataVisualization/configs"
	"dataVisualization/internal/cache"
	"dataVisualization/internal/config"
	"dataVisualization/internal/dao"
	"dataVisualization/internal/model"
	"fmt"
	"testing"
)

var service *dataVisualization

func init() {
	err := config.Init(configs.Path("dataVisualization.yml"))
	if err != nil {
		panic(err)
	}
	service = &dataVisualization{
		deviceDao:      dao.NewDeviceDao(model.GetDB(), cache.NewDeviceCache(model.GetCacheType())),
		userDeviceDao:  dao.NewUserDeviceDao(model.GetDB(), cache.NewUserDeviceCache(model.GetCacheType())),
		deviceTypeDao:  dao.NewDeviceTypeDao(model.GetDB(), cache.NewDeviceTypeCache(model.GetCacheType())),
		deviceDaoModel: dao.NewDeviceDataModelDao(model.GetDB(), cache.NewDeviceDataModelCache(model.GetCacheType())),
	}
}

func Test_DeviceDataGet(t *testing.T) {
	res, err := service.DeviceDataGet(context.Background(), &dataVisualizationV1.DeviceDataGetReq{
		UserId:   2,
		Code:     "e2",
		CurrPage: 1,
		PageSize: 10,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func TestDataVisualization_DeviceDataCurve(t *testing.T) {
	res, err := service.DeviceDataCurve(context.Background(), &dataVisualizationV1.DeviceDataCurveReq{
		UserId:     1,
		//DeviceCode: "e2",
		Interval:   dataVisualizationV1.IntervalType_Month,
		ChartType:  []dataVisualizationV1.ChartType{dataVisualizationV1.ChartType_Pip, dataVisualizationV1.ChartType_Line},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
