package service

import (
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	managerV1 "manager/api/manager/v1"
	"manager/internal/common"
	"manager/internal/model"
)

// buildUserDaoParams
//
//	@Description: 构建用户查询参数
//	@Author zzh
//	@receiver s
//	@param req
//	@return *query.Params
func (s *manager) buildUserDaoParams(req *managerV1.UserGetReq) *query.Params {
	userParams := &query.Params{}

	userParams.Page = int(req.CurrPage)
	userParams.Size = int(req.PageSize)

	userParams.Columns = append(userParams.Columns, query.Column{
		Name:  "is_valid",
		Value: common.Enable.Uint32(),
	})

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
	//手机号
	if len(req.Phone) != 0 {
		userParams.Columns = append(userParams.Columns, query.Column{
			Name:  "phone",
			Value: req.Phone,
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
func (s *manager) buildUserGetRes(users []*model.User) (userRes []*managerV1.User) {
	for _, user := range users {
		roleIds, _ := common.IdToArray(user.RoleId)
		var roleId uint64
		if len(roleIds) != 0 {
			roleId = roleIds[0]
		}
		userRes = append(userRes, &managerV1.User{
			Id:        user.ID,
			Name:      user.Name,
			Account:   user.Account,
			Sex:       user.Sex,
			Phone:     user.Phone,
			AvatarUrl: user.AvatarUrl,
			IsValid:   uint32(user.IsValid),
			RoleId:    roleId,
		})
	}
	return
}
