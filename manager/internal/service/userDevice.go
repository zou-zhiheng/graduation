package service

import (
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	managerV1 "manager/api/manager/v1"
	"manager/internal/common"
	"manager/internal/model"
)

// buildUserDeviceDaoParams
//
//	@Description: 构建用户设备查询参数
//	@Author zzh
//	@receiver s
//	@param req
//	@return *query.Params
func (s *manager) buildUserDeviceDaoParams(req *managerV1.UserDeviceGetReq) *query.Params {
	params := &query.Params{}

	params.Page = int(req.CurrPage)
	params.Size = int(req.PageSize)

	//用户名
	if len(req.UserName) != 0 {
		params.Columns = append(params.Columns, query.Column{
			Name:  "user_name",
			Value: req.UserName,
		})
	}

	//设备名称
	if len(req.DeviceName) != 0 {
		params.Columns = append(params.Columns, query.Column{
			Name:  "device_name",
			Value: req.DeviceName,
		})
	}

	//设备编号
	if len(req.DeviceCode) != 0 {
		params.Columns = append(params.Columns, query.Column{
			Name:  "device_code",
			Value: req.DeviceCode,
		})
	}

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

	return params
}

// buildUserDeviceGetRes
//
//	@Description: 构建用户设备信息返回参数
//	@Author zzh
//	@receiver s
//	@param userDevices
//	@return userDeviceRes
func (s *manager) buildUserDeviceGetRes(userDevices []*model.UserDevice) (userDeviceRes []*managerV1.UserDevice) {
	for _, userDevice := range userDevices {
		userDeviceRes = append(userDeviceRes, &managerV1.UserDevice{
			Id:         userDevice.ID,
			DeviceName: userDevice.DeviceName,
			DeviceCode: userDevice.DeviceCode,
			UserName:   userDevice.UserName,
			CreateTime: userDevice.CreateTime.Format(common.TimeFormat),
		})
	}
	return
}
