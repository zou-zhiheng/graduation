package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/logger"
	managerV1 "manager/api/manager/v1"
	"manager/internal/common"
	"manager/internal/model"
	"sync"
	"time"
)

// buildApiDaoParams
//
//	@Description: 构建设备查询参数
//	@Author zzh
//	@receiver s
//	@param req
//	@return *query.Params
func (s *manager) buildDeviceDaoParams(ctx context.Context, req *managerV1.DeviceGetReq) (*query.Params, error) {
	params := &query.Params{}

	params.Page = int(req.CurrPage)
	params.Size = int(req.PageSize)

	if len(req.DeviceTypeName) != 0 {
		deviceTypeCondition := &query.Conditions{Columns: []query.Column{
			{
				Name:  "name",
				Value: req.DeviceTypeName,
			},
		}}
		deviceTypeInfo, err := s.deviceDao.GetByCondition(ctx, deviceTypeCondition)
		if err != nil {
			return nil, err
		}
		params.Columns = append(params.Columns, query.Column{
			Name:  "device_type_id",
			Value: deviceTypeInfo.ID,
		})
	}
	if req.Id != 0 {
		params.Columns = append(params.Columns, query.Column{
			Name:  "id",
			Value: req.Id,
		})
	}
	//名称
	if len(req.Name) != 0 {
		params.Columns = append(params.Columns, query.Column{
			Name:  "name",
			Value: req.Name,
		})
	}

	//编号
	if len(req.Code) != 0 {
		params.Columns = append(params.Columns, query.Column{
			Name:  "code",
			Value: req.Code,
		})
	}

	////设备类型ID
	//if deviceTypeId != 0 {
	//	params.Columns = append(params.Columns, query.Column{
	//		Name:  "device_type_id",
	//		Value: deviceTypeId,
	//	})
	//}

	if len(req.StartTime) != 0 {
		params.Columns = append(params.Columns, query.Column{
			Name:  "create_time",
			Value: req.StartTime,
			Exp:   ">=",
		})
	}

	if len(req.EndTime) != 0 {
		params.Columns = append(params.Columns, query.Column{
			Name:  "create_time",
			Value: req.EndTime,
			Exp:   "<=",
		})
	}

	if params.Page != 0 {
		params.Page--
	}

	return params, nil
}

// buildApiGetRes
//
//	@Description: 构建接口查询响应参数
//	@Author zzh
//	@receiver s
//	@param apis
//	@return apiRes
func (s *manager) buildDeviceGetRes(ctx context.Context, devices []*model.Device) (deviceRes []*managerV1.Device, err error) {

	var deviceTypeNameMap = make(map[uint64]string)
	var ok bool

	for _, device := range devices {
		if _, ok = deviceTypeNameMap[device.DeviceTypeID]; !ok {
			//查找设备类型名称
			var deviceType *model.DeviceType
			deviceType, err = s.deviceTypeDao.GetByID(ctx, device.DeviceTypeID)
			if err != nil {
				return nil, err
			}

			deviceTypeNameMap[device.DeviceTypeID] = deviceType.Name
		}
		deviceRes = append(deviceRes, &managerV1.Device{
			Id:             device.ID,
			Name:           device.Name,
			Code:           device.Code,
			DeviceTypeName: deviceTypeNameMap[device.DeviceTypeID],
			DeviceTypeId:   device.DeviceTypeID,
			State:          uint32(device.State),
			CheckTime:      float32(device.CheckTime),
		})
	}
	return
}

// deviceDataParse
//
//	@Description: 解析设备数据
//	@Author zzh
//	@receiver s
//	@param ctx
//	@param deviceDatas
//	@return parseData
//	@return err
func (s *manager) deviceDataParse(ctx context.Context, deviceDatas []*managerV1.DeviceData) (parseData map[string][]*model.DeviceDataModel, err error) {

	parseData = make(map[string][]*model.DeviceDataModel)
	for _, deviceData := range deviceDatas {
		var jsonByte []byte
		jsonByte, err = json.Marshal(deviceData.Data)
		if err != nil {
			return nil, err
		}

		var createTime time.Time
		if len(deviceData.CreateTime) != 0 {
			createTime, err = common.TimeParseAsCST(deviceData.CreateTime)
			if err != nil {
				return nil, err
			}
			createTime = createTime.Truncate(time.Minute)
		} else {
			createTime = time.Now().Local().Truncate(time.Minute)
		}
		parseData[deviceData.Code] = append(parseData[deviceData.Code], &model.DeviceDataModel{
			Data:       string(jsonByte),
			CreateTime: &createTime,
		})
	}

	return
}

// saveDeviceData
//
//	@Description: 保存设备解析后的数据
//	@Author zzh
//	@receiver s
//	@param ctx
//	@param deviceDatas
//	@return err
func (s *manager) saveDeviceData(ctx context.Context, deviceDatas map[string][]*model.DeviceDataModel) (err error) {

	if len(deviceDatas) == 0 {
		return errors.New("deviceDatas can not be null")
	}

	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	for code, val := range deviceDatas {
		for _, deviceData := range val {
			err = tx.WithContext(ctx).Table(code).Create(deviceData).Error
			if err != nil {
				return
			}
		}
	}

	err = tx.Commit().Error
	return
}

// checkDeviceState
//
//	@Description: 检查并更新设备状态
//	@Author zzh
//	@receiver s
func (s *manager) checkDeviceState(ctx context.Context) {

	//查询所有设备
	deviceParams := &query.Params{Columns: []query.Column{
		{
			Name:  "check_time",
			Value: 0,
			Exp:   "!=",
		},
		//{
		//	Name:  "state",
		//	Value: common.Down.Int(),
		//	Exp:   "!=",
		//},
	}}
	deviceInfo, _, err := s.deviceDao.GetByColumns(ctx, deviceParams)
	if err != nil {
		logger.Error("check device status error", logger.Err(err))
		return
	}

	var wg = sync.WaitGroup{}
	for _, device := range deviceInfo {
		device := device
		go func() {
			defer func() {
				wg.Done()
			}()
			wg.Add(1)
			s.checkDeviceStateByData(ctx, device)
		}()
	}
	wg.Wait()

}

// checkDeviceStateByData
//
//	@Description: 通过检查当前设备的最新数据是否超过当前设备的监测时间判断设备是否正常
//	@Author zzh
//	@receiver s
func (s *manager) checkDeviceStateByData(ctx context.Context, device *model.Device) {
	//查询设备最新一条数据
	deviceData, err := s.deviceDataModelDao.GetDeviceLatestData(ctx, device)
	if err != nil {
		if !errors.Is(err, model.ErrRecordNotFound) {
			logger.Error("checkDeviceStateByData.deviceDataModelDao.GetDeviceLatestData error", logger.Any("device code", device.Code), logger.Err(err))
		}
		_ = s.updateDeviceState(ctx, device, common.UnNormal.Int())
		return
	}
	var state int
	interval := time.Now().Sub(*deviceData.CreateTime)
	if interval.Minutes() > device.CheckTime {
		//是否需要修改设备状态
		if device.State == common.UnNormal.Int() || device.State == common.Down.Int() {
			return
		}
		state = common.UnNormal.Int()
	} else {
		//是否需要修改设备状态
		if device.State == common.Running.Int() {
			return
		}
		state = common.Running.Int()
	}
	//修改设备状态
	_ = s.updateDeviceState(ctx, device, state)
}

func (s *manager) updateDeviceState(ctx context.Context, device *model.Device, state int) (err error) {
	device.State = state
	err = s.deviceDao.UpdateByID(ctx, device)
	if err != nil {
		logger.Error("checkDeviceStateByData.deviceDao.UpdateByID error", logger.Any("device code", device.Code), logger.Err(err))
	}
	return
}
