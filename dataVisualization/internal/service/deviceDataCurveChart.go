package service

import (
	"context"
	dataVisualizationV1 "dataVisualization/api/dataVisualization/v1"
	"dataVisualization/internal/common"
	"dataVisualization/internal/model"
	"encoding/json"
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
)

type DataPip struct {
	Name  string
	Value float32
}

const (
	baseValue = 1
)

func (s *dataVisualization) getDeviceDataPipChart(ctx context.Context, deviceIds []uint64) (pipInfo *dataVisualizationV1.DataLine, err error) {

	//查询设备类型
	var deviceTypeMap = make(map[uint64]*DataPip)
	pipInfo = &dataVisualizationV1.DataLine{}
	for _, deviceId := range deviceIds {
		//查询设备
		var deviceInfo = &model.Device{}
		deviceInfo, err = s.deviceDao.GetByID(ctx, deviceId)
		if err != nil {
			return nil, err
		}
		//查询设备类型
		if _, ok := deviceTypeMap[deviceInfo.DeviceTypeID]; !ok {
			var deviceTypeInfo = &model.DeviceType{}
			deviceTypeInfo, err = s.deviceTypeDao.GetByID(context.Background(), deviceInfo.DeviceTypeID)
			if err != nil {
				return nil, err
			}
			deviceTypeMap[deviceInfo.DeviceTypeID] = &DataPip{
				Name:  deviceTypeInfo.Name,
				Value: baseValue,
			}
		} else {
			deviceTypeMap[deviceInfo.DeviceTypeID].Value++
		}
	}

	for _, val := range deviceTypeMap {
		pipInfo.Key = append(pipInfo.Key, val.Name)
		pipInfo.Value = append(pipInfo.Value, val.Value)
	}

	return pipInfo, err
}

func (s *dataVisualization) getElectLine(ctx context.Context, deviceIds []uint64) (elect *dataVisualizationV1.DataLine, volt *dataVisualizationV1.DataLine, err error) {

	//选择一个数据量最大的设备
	var deviceInfo, device *model.Device
	var maxCount int64
	var count int64
	for _, deviceId := range deviceIds {
		if deviceId != 1 {
			continue
		}
		device, err = s.deviceDao.GetByID(ctx, deviceId)
		if err != nil {
			return
		}
		count, err = s.deviceDaoModel.CountDeviceDataByColumns(context.Background(), &query.Params{}, device.Code)
		if err != nil {
			return
		}
		if maxCount < count {
			deviceInfo = device
			maxCount = count
		}
	}

	if deviceInfo != nil && maxCount != 0 {
		elect = &dataVisualizationV1.DataLine{}
		volt = &dataVisualizationV1.DataLine{}
		var deviceDataInfo []*model.DeviceDataModel
		deviceDataInfo, _, err = s.deviceDaoModel.GetDeviceDataByColumns(context.Background(), &query.Params{Size: 30}, deviceInfo.Code)
		if err != nil {
			return
		}

		//开始处理数据
		for _, deviceData := range deviceDataInfo {
			//转为数组
			var dataInfo []*dataVisualizationV1.DataDetail
			_ = json.Unmarshal([]byte(deviceData.Data), &dataInfo)
			elect.Key = append(elect.Key, deviceData.CreateTime.Format(common.TimeLayOut))
			elect.Value = append(elect.Value, dataInfo[0].Value)
			volt.Key = append(volt.Key, deviceData.CreateTime.Format(common.TimeLayOut))
			volt.Value = append(volt.Value, dataInfo[1].Value)
		}
	}

	return

}
