package service

import (
	"context"
	dataVisualizationV1 "dataVisualization/api/dataVisualization/v1"
	"dataVisualization/internal/common"
	"dataVisualization/internal/model"
	"encoding/json"
	"errors"
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"time"
)

// buildApiDaoParams
//
//	@Description: 构建接口查询参数
//	@Author zzh
//	@receiver s
//	@param req
//	@return *query.Params
func (s *dataVisualization) buildDeviceDataParams(req *dataVisualizationV1.DeviceDataGetReq) *query.Params {
	userParams := &query.Params{}

	userParams.Page = int(req.CurrPage)
	userParams.Size = int(req.PageSize)

	if len(req.StartTime) != 0 {
		userParams.Columns = append(userParams.Columns, query.Column{
			Name:  "create_time",
			Value: req.StartTime,
			Exp:   ">=",
		})
	}

	if len(req.EndTime) != 0 {
		userParams.Columns = append(userParams.Columns, query.Column{
			Name:  "create_time",
			Value: req.EndTime,
			Exp:   "<=",
		})
	}

	if userParams.Size == 0 {
		userParams.Size = 10
	}

	if userParams.Page != 0 {
		userParams.Page--
	}

	return userParams
}

// buildDeviceDataGetRes
//
//	@Description: 构建设备数据查询参数
//	@Author zzh
//	@receiver s
//	@param deviceDataInfo
//	@return deviceDataGetRes
func (s *dataVisualization) buildDeviceDataGetRes(deviceDataInfo []*model.DeviceDataModel) (deviceDataGetRes []*dataVisualizationV1.DeviceData) {

	for _, deviceData := range deviceDataInfo {
		dataDetailInfo := []*model.DataDetail{}
		_ = json.Unmarshal([]byte(deviceData.Data), &dataDetailInfo)
		var detail []*dataVisualizationV1.DataDetail
		for _, dataDetail := range dataDetailInfo {
			detail = append(detail, &dataVisualizationV1.DataDetail{
				Key:   dataDetail.Key,
				Value: dataDetail.Value,
				Unit:  dataDetail.Unit,
			})
		}
		deviceDataGetRes = append(deviceDataGetRes, &dataVisualizationV1.DeviceData{
			Data:       detail,
			CreateTime: deviceData.CreateTime.Format(common.TimeLayOut),
		})
	}

	return
}

func (s *dataVisualization) getDeviceDataCurveByInterval(ctx context.Context, deviceIds []uint64, interval dataVisualizationV1.IntervalType) (*dataVisualizationV1.DataLine, error) {

	//记录每个设备的参数列表
	var paramsMap = make(map[string][]*DeviceDataCurveParams)
	var dataLen int
	//构建每个设备的参数列表
	for _, deviceId := range deviceIds {
		//查询设备信息
		deviceInfo, err := s.deviceDao.GetByID(context.Background(), deviceId)
		if err != nil {
			return nil, err
		}
		paramsMap[deviceInfo.Code], err = s.buildCountDeviceDataByColumnsParams(context.Background(), interval, deviceInfo.Code)
		if err != nil {
			return nil, err
		}
		if dataLen == 0 {
			dataLen = len(paramsMap[deviceInfo.Code])
		} else {
			if dataLen != len(paramsMap[deviceInfo.Code]) {
				return nil, errors.New("dataLen error")
			}
		}
	}
	var lineInfo = &dataVisualizationV1.DataLine{}
	//res := [][]string{{"createTime", "count"}}
	//查询每个设备数据量并相加，分钟
	for i := 0; i < dataLen; i++ {
		var timeStr string
		var count int
		for deviceCode, params := range paramsMap {
			//查询当前设备的数据
			//是否需要查询数据库
			if params[i].ToDB {
				dataCount, err := s.deviceDaoModel.CountDeviceDataByColumns(context.Background(), &query.Params{Columns: []query.Column{
					{
						Name:  "createTime",
						Value: params[i].EndTime,
						Exp:   "<",
					},
					{
						Name:  "createTime",
						Value: params[i].StartTime,
						Exp:   ">=",
					},
				},
				}, deviceCode)
				if err != nil {
					return nil, err
				}
				count += int(dataCount)
			}
			if len(timeStr) == 0 {
				timeStr = params[i].StartTime.Format(common.TimeLayOut)
			}
		}
		lineInfo.Key = append(lineInfo.Key, timeStr)
		lineInfo.Value = append(lineInfo.Value, float32(count))
		//res = append(res, []string{timeStr, fmt.Sprintf("%d", count)})

	}

	return lineInfo, nil
}

type DeviceDataCurveParams struct {
	ToDB      bool //是否需要查询数据库
	StartTime time.Time
	EndTime   time.Time
}

func (s *dataVisualization) buildCountDeviceDataByColumnsParams(ctx context.Context, interval dataVisualizationV1.IntervalType, tableName string) (params []*DeviceDataCurveParams, err error) {

	var base, loop int
	switch interval {
	case dataVisualizationV1.IntervalType_Day:
		base = 60
		loop = 24
	case dataVisualizationV1.IntervalType_Week:
		base = 60 * 24
		loop = 7
	case dataVisualizationV1.IntervalType_Month:
		base = 60 * 24 * 7
		loop = 6
	default:
		return nil, errors.New("interval not support")
	}

	endTime := time.Now()
	startTime := endTime.Add(-1 * time.Duration(base) * time.Minute)

	//查询数据表最早一条数据
	deviceDataInfos, _, err := s.deviceDaoModel.GetDeviceDataByColumns(context.Background(), &query.Params{Size: 1}, tableName)
	if err != nil {
		return nil, err
	}

	if len(deviceDataInfos) != 1 {
		return nil, errors.New("unknown error")
	}

	deviceDataInfo := deviceDataInfos[0]

	var toDB = true
	for i := 0; i < loop; i++ {
		param := &DeviceDataCurveParams{}
		if toDB {

			if deviceDataInfo.CreateTime.Sub(endTime) < time.Minute {
				if deviceDataInfo.CreateTime.Sub(startTime) > time.Minute {
					startTime = *deviceDataInfo.CreateTime
				}
			} else {
				toDB = false
			}
		}
		param.ToDB = toDB
		param.StartTime = startTime
		param.EndTime = endTime
		//更新startTime、endTime
		startTime = startTime.Add(-24 * time.Hour)
		endTime = endTime.Add(-24 * time.Hour)

		params = append(params, param)
	}

	return
}
