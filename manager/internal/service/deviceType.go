package service

import (
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	managerV1 "manager/api/manager/v1"
	"manager/internal/model"
)

// buildDeviceTypeDaoParams
//
//	@Description: 构建设备类型查询参数
//	@Author zzh
//	@receiver s
//	@param req
//	@return *query.Params
func (s *manager) buildDeviceTypeDaoParams(req *managerV1.DeviceTypeGetReq) *query.Params {
	params := &query.Params{}

	params.Page = int(req.CurrPage)
	params.Size = int(req.PageSize)

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

	return params
}

// buildDeviceTypeGetRes
//
//	@Description: 构建设备类型查询返回参数
//	@Author zzh
//	@receiver s
//	@param deviceTypes
//	@return deviceTypeRes
func (s *manager) buildDeviceTypeGetRes(deviceTypes []*model.DeviceType) (deviceTypeRes []*managerV1.DeviceType) {
	for _, deviceType := range deviceTypes {
		deviceTypeRes = append(deviceTypeRes, &managerV1.DeviceType{
			Id:       deviceType.ID,
			Name:     deviceType.Name,
			Describe: deviceType.Describe,
		})
	}
	return
}
