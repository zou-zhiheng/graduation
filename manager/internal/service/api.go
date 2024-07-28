package service

import (
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	managerV1 "manager/api/manager/v1"
	"manager/internal/model"
)

// buildApiDaoParams
//
//	@Description: 构建接口查询参数
//	@Author zzh
//	@receiver s
//	@param req
//	@return *query.Params
func (s *manager) buildApiDaoParams(req *managerV1.ApiGetReq) *query.Params {
	userParams := &query.Params{}

	userParams.Page = int(req.CurrPage)
	userParams.Size = int(req.PageSize)

	if req.Id != 0 {
		userParams.Columns = append(userParams.Columns, query.Column{
			Name:  "id",
			Value: req.Id,
		})
	}
	//名称
	if len(req.Name) != 0 {
		userParams.Columns = append(userParams.Columns, query.Column{
			Name:  "name",
			Value: req.Name,
		})
	}

	if len(req.Url) != 0 {
		userParams.Columns = append(userParams.Columns, query.Column{
			Name:  "url",
			Value: req.Url,
		})
	}

	if len(req.Method) != 0 {
		userParams.Columns = append(userParams.Columns, query.Column{
			Name:  "method",
			Value: req.Method,
		})
	}

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

	if userParams.Page != 0 {
		userParams.Page--
	}

	return userParams
}

// buildApiGetRes
//
//	@Description: 构建接口查询响应参数
//	@Author zzh
//	@receiver s
//	@param apis
//	@return apiRes
func (s *manager) buildApiGetRes(apis []*model.Api) (apiRes []*managerV1.Api) {
	for _, api := range apis {
		apiRes = append(apiRes, &managerV1.Api{
			Id:     api.ID,
			Name:   api.Name,
			Url:    api.Url,
			Method: api.Method,
			Desc:   api.Desc,
		})
	}
	return
}
