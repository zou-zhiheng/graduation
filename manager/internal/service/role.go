package service

import (
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	managerV1 "manager/api/manager/v1"
	"manager/internal/common"
	"manager/internal/model"
)

func (s *manager) buildRoleDaoParams(req *managerV1.RoleGetReq) *query.Params {
	userParams := &query.Params{}

	userParams.Page = int(req.CurrPage)
	userParams.Size = int(req.PageSize)

	if req.Id != 0 {
		userParams.Columns = append(userParams.Columns, query.Column{
			Name:  "id",
			Value: req.Id,
		})
	}
	//账户
	if len(req.Name) != 0 {
		userParams.Columns = append(userParams.Columns, query.Column{
			Name:  "name",
			Value: req.Name,
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

	return userParams
}

// buildUserGetRes
//
//	@Description: 构建查询用户请求参数
//	@Author zzh
//	@receiver s
//	@param users
//	@return userRes
func (s *manager) buildRoleGetRes(roles []*model.Role) (roleRes []*managerV1.Role) {
	for _, role := range roles {
		apis, _ := common.IdToArray(role.Api)
		roleRes = append(roleRes, &managerV1.Role{
			Id:   role.ID,
			Name: role.Name,
			Code: role.Code,
			Apis: apis,
			Desc: role.Desc,
		})
	}
	return
}
