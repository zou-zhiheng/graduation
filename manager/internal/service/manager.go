// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zhufuyi/sponge/pkg/errcode"
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"manager/internal/cache"
	"manager/internal/common"
	"manager/internal/dao"
	"manager/internal/ecode"
	"manager/internal/model"
	"math/rand"
	"strconv"
	"time"

	managerV1 "manager/api/manager/v1"

	//"manager/internal/cache"
	//"manager/internal/dao"
	//"manager/internal/ecode"
	//"manager/internal/model"

	//"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	//"github.com/zhufuyi/sponge/pkg/logger"

	"google.golang.org/grpc"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		managerV1.RegisterManagerServer(server, NewManagerServer())
	})
}

var _ managerV1.ManagerServer = (*manager)(nil)

type manager struct {
	managerV1.UnimplementedManagerServer

	// example:
	//		iDao dao.ManagerDao
	userDao            dao.UserDao
	roleDao            dao.RoleDao
	apiDao             dao.ApiDao
	deviceDao          dao.DeviceDao
	deviceDataModelDao dao.DeviceDataModelDao
	deviceTypeDao      dao.DeviceTypeDao
	userDeviceDao      dao.UserDeviceDao
}

// NewManagerServer create a server
func NewManagerServer() managerV1.ManagerServer {
	s := &manager{
		userDao:            dao.NewUserDao(model.GetDB(), cache.NewUserCache(model.GetCacheType())),
		roleDao:            dao.NewRoleDao(model.GetDB(), cache.NewRoleCache(model.GetCacheType())),
		apiDao:             dao.NewApiDao(model.GetDB(), cache.NewApiCache(model.GetCacheType())),
		deviceDao:          dao.NewDeviceDao(model.GetDB(), cache.NewDeviceCache(model.GetCacheType())),
		deviceDataModelDao: dao.NewDeviceDataModelDao(model.GetDB(), cache.NewDeviceDataModelCache(model.GetCacheType())),
		deviceTypeDao:      dao.NewDeviceTypeDao(model.GetDB(), cache.NewDeviceTypeCache(model.GetCacheType())),
		userDeviceDao:      dao.NewUserDeviceDao(model.GetDB(), cache.NewUserDeviceCache(model.GetCacheType())),
	}
	go s.NewTick()
	return s
}

func (s *manager) NewTick() {
	ch := time.Tick(time.Minute)
	select {
	case <-ch:
		logger.Info("开始执行本轮的定时任务")
		s.checkDeviceState(context.Background())
		logger.Info("本轮的定时任务执行结束")
	}
}

var (
	operationSuccess = &managerV1.OperationRes{
		Status: common.OperationStateSuccess.Uint32(),
		Msg:    common.OperationStateSuccess.String(),
	}

	operationFailed = &managerV1.OperationRes{
		Status: common.OperationStateFailed.Uint32(),
		Msg:    common.OperationStateFailed.String(),
	}
)

// log
//
//	@Description: 日志打印
//	@Author zzh
//	@param msg
//	@param rpcErr
//	@param err
//	@param field
//	@param ctx
//	@return *errcode.RPCStatus
func log(ctx context.Context, msg string, rpcErr *errcode.RPCStatus, err error, key string, val interface{}) error {

	logger.Info(msg, logger.Err(err), logger.Any(key, val), interceptor.ServerCtxRequestIDField(ctx))

	return rpcErr.Err()
}

// Login 登录
func (s *manager) Login(ctx context.Context, req *managerV1.LoginReq) (*managerV1.LoginRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	userCondition := &query.Conditions{Columns: []query.Column{
		{
			Name:  "account",
			Value: req.Account,
		},
	}}
	userInfo, err := s.userDao.GetByCondition(context.Background(), userCondition)
	if err != nil {
		return nil, log(ctx, "login error", ecode.StatusInvalidParams, err, "req", req)
	}

	//密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.Password+userInfo.Salt)); err != nil {
		return nil, log(ctx, "check password err", ecode.StatusInvalidParams, err, "detail", "account or password is invalid")
	}

	//生成token
	jwtToken := common.NewJwtToken()
	token, err := jwtToken.CreateToken(*userInfo)
	if err != nil {
		return nil, log(ctx, "generate token err", ecode.StatusInternalServerError, err, "detail", err.Error())
	}

	var roleId uint64
	roleIds, _ := common.IdToArray(userInfo.RoleId)
	if len(roleIds)!=0{
		roleId = roleIds[0]
	}

	return &managerV1.LoginRes{
		Token: token,
		User: &managerV1.User{
			Id:        userInfo.ID,
			Name:      userInfo.Name,
			Account:   userInfo.Account,
			AvatarUrl: userInfo.AvatarUrl,
			Sex:       userInfo.Sex,
			Phone:     userInfo.Phone,
			RoleId:    roleId,
		},
	}, nil
}

// UserRegister 用户注册
func (s *manager) UserRegister(ctx context.Context, req *managerV1.UserRegisterReq) (*managerV1.OperationRes, error) {

	// fill in the business logic code here
	// example:
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//参数校验 todo

	//检查要注册的账号
	userCondition := &query.Conditions{Columns: []query.Column{
		{
			Name:  "account",
			Value: req.Account,
		},
	}}
	_, err = s.userDao.GetByCondition(context.Background(), userCondition)
	if !errors.Is(err, model.ErrRecordNotFound) {
		return nil, log(ctx, "register err", ecode.StatusInvalidParams, nil, "account", req.Account)
	}

	//生成密码
	//根据时间戳生成随机数，防止恶意伪造
	random := rand.New(rand.NewSource(time.Now().Unix()))
	//生成盐
	salt := strconv.FormatInt(random.Int63(), 10)
	//密码加密加盐
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password+salt), bcrypt.DefaultCost)
	if err != nil {
		return nil, log(ctx, "register err", ecode.StatusInvalidParams, nil, "req", req)
	}

	var roleId []byte
	if len(req.RoleIds) != 0 {
		//绑定用户角色
		roleId, err = json.Marshal(req.RoleIds)
		if err != nil {
			return nil, log(ctx, "register err", ecode.StatusInvalidParams, nil, "req", req)
		}
	}

	userInfo := &model.User{
		Account:   req.Account,
		Name:      req.Name,
		Password:  string(encryptedPass),
		Salt:      salt,
		IsValid:   common.Enable.Int(),
		RoleId:    string(roleId),
		AvatarUrl: req.AvatarUrl,
		Sex:       req.Sex,
		Phone:     req.Phone,
	}

	//插入数据
	err = s.userDao.Create(context.Background(), userInfo)
	if err != nil {
		return nil, log(ctx, "register sql err", ecode.StatusInternalServerError, err, "req", req)
	}

	return &managerV1.OperationRes{
		Status: common.OperationStateSuccess.Uint32(),
		Msg:    common.OperationStateSuccess.String(),
	}, nil
}

// UserGet 用户查询
func (s *manager) UserGet(ctx context.Context, req *managerV1.UserGetReq) (*managerV1.UserGetRes, error) {
	//panic("implement me")

	// fill in the business logic code here
	// example:

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	userParams := s.buildUserDaoParams(req)
	userInfos, count, err := s.userDao.GetByColumns(ctx, userParams)
	if err != nil {
		logger.Error("manager.userDao.GetByColumns error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	return &managerV1.UserGetRes{
		User:  s.buildUserGetRes(userInfos),
		Count: uint64(count),
	}, nil
}

// UserUpdate 用户修改
func (s *manager) UserUpdate(ctx context.Context, req *managerV1.UserUpdateReq) (*managerV1.OperationRes, error) {
	//panic("implement me")

	// fill in the business logic code here
	// example:
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//参数校验 todo

	//查询用户是否存在
	userInfo, err := s.userDao.GetByID(context.Background(), req.Id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			return operationFailed, log(ctx, "manager.UserUpdate.userDao.GetByID error", ecode.StatusInvalidParams, err, "req", req)
		}
		return operationFailed, log(ctx, "manager.UserUpdate.userDao.GetByID error", ecode.StatusInternalServerError, err, "req", req)
	}

	//修改
	//名称
	userInfo.Name = req.Name
	//头像地址
	userInfo.AvatarUrl = req.AvatarUrl
	//性别
	userInfo.Sex = req.Sex
	//电话号码
	userInfo.Phone = req.Phone
	//角色
	roleId, _ := common.ArrayToId(req.RoleId)
	userInfo.RoleId = roleId

	err = s.userDao.UpdateByID(ctx, userInfo)
	if err != nil {
		return operationFailed, log(ctx, "manager.UserUpdate.userDao.GetByID error", ecode.StatusInternalServerError, err, "req", req)
	}

	return operationSuccess, nil
}

// ResetPassword 密码重置
func (s *manager) ResetPassword(ctx context.Context, req *managerV1.ResetPasswordReq) (*managerV1.OperationRes, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	// 	ctx = interceptor.WrapServerCtx(ctx)
	//
	// 	reply, err := s.iDao.ResetPassword(ctx, &model.manager{
	//     	Id: req.Id,
	//     	Name: req.Name,
	//     	Account: req.Account,
	//     	AvatarUrl: req.AvatarUrl,
	//     	IsValid: req.IsValid,
	//     	Sex: req.Sex,
	//     	Phone: req.Phone,
	//     	RoleId: req.RoleId,
	//     })
	// 	if err != nil {
	//			logger.Warn("ResetPassword error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &managerV1.OperationRes{
	//     	Status: reply.Status,
	//     	Msg: reply.Msg,
	//     }, nil
}

// UserDelete 用户注销
func (s *manager) UserDelete(ctx context.Context, req *managerV1.UserDeleteReq) (*managerV1.OperationRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	userInfo, err := s.userDao.GetByID(context.Background(), req.Id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			return operationFailed, log(ctx, "manager.UserDelete.userDao.GetByID error", ecode.StatusInvalidParams, err, "req", req)
		}
		return operationFailed, log(ctx, "manager.UserDelete.userDao.GetByID error", ecode.StatusInternalServerError, err, "req", req)
	}

	//假删除
	userInfo.IsValid = common.Disable.Int()
	//更新数据库
	err = s.userDao.UpdateByID(context.Background(), userInfo)
	if err != nil {
		return operationFailed, log(ctx, "manager.UserDelete.userDao.UpdateByID error", ecode.StatusInternalServerError, err, "req", req)
	}

	return operationSuccess, nil
}

// RoleCreate 角色创建
func (s *manager) RoleCreate(ctx context.Context, req *managerV1.RoleCreateReq) (*managerV1.OperationRes, error) {
	//panic("implement me")

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//查询当前角色是否存在
	roleCondition := &query.Conditions{Columns: []query.Column{
		{
			Name:  "name",
			Value: req.Name,
		},
	}}
	_, err = s.roleDao.GetByCondition(context.Background(), roleCondition)
	if !errors.Is(err, model.ErrRecordNotFound) {
		return operationFailed, log(ctx, "manager.RoleCreate.roleDao.GetByCondition error", ecode.StatusInvalidParams, err, "req", req)
	}

	api, _ := common.ArrayToId(req.Apis)

	roleInfo := &model.Role{
		Name: req.Name,
		Api:  api,
		Desc: req.Desc,
	}
	//插入数据
	err = s.roleDao.Create(context.Background(), roleInfo)
	if err != nil {
		return operationFailed, log(ctx, "manager.RoleCreate.roleDao.Create error", ecode.StatusInvalidParams, err, "req", req)
	}

	return operationSuccess, nil
}

// RoleGet 角色查询
func (s *manager) RoleGet(ctx context.Context, req *managerV1.RoleGetReq) (*managerV1.RoleGetRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//构建查询参数
	roleParams := s.buildRoleDaoParams(req)
	roleInfo, count, err := s.roleDao.GetByColumns(context.Background(), roleParams)
	if err != nil {
		return nil, log(ctx, "manager.RoleGet.roleDao.GetByColumns error", ecode.StatusInvalidParams, err, "req", req)
	}

	return &managerV1.RoleGetRes{
		Role:  s.buildRoleGetRes(roleInfo),
		Count: uint64(count),
	}, nil
}

// RoleUpdate 角色修改
func (s *manager) RoleUpdate(ctx context.Context, req *managerV1.RoleUpdateReq) (*managerV1.OperationRes, error) {
	//panic("implement me")

	// fill in the business logic code here
	// example:
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	roleInfo, err := s.roleDao.GetByID(context.Background(), req.Id)
	if err != nil {
		return operationFailed, log(ctx, "manager.RoleUpdate.roleDao.GetByID error", ecode.StatusInvalidParams, err, "req", req)
	}

	if len(req.Name) != 0 {
		roleCondition := &query.Conditions{Columns: []query.Column{
			{
				Name:  "name",
				Value: req.Name,
			},
		}}
		_, err = s.roleDao.GetByCondition(context.Background(), roleCondition)
		if err != nil {
			return operationFailed, log(ctx, "manager.RoleUpdate.roleDao.GetByCondition error", ecode.StatusInvalidParams, err, "req", req)
		}
	}

	//名称
	roleInfo.Name = req.Name
	//描述
	roleInfo.Desc = req.Desc
	//接口
	api, _ := common.ArrayToId(req.Apis)
	roleInfo.Api = api

	err = s.roleDao.UpdateByID(context.Background(), roleInfo)
	if err != nil {
		return nil, log(ctx, "manager.RoleGet.roleDao.UpdateByID error", ecode.StatusInternalServerError, err, "req", req)
	}

	return operationSuccess, nil
}

// RoleDelete 角色删除
func (s *manager) RoleDelete(ctx context.Context, req *managerV1.RoleDeleteReq) (*managerV1.OperationRes, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//查询角色
	roleInfo, err := s.roleDao.GetByID(context.Background(), req.Id)
	if err != nil {
		return operationFailed, log(ctx, "manager.RoleDelete.roleDao.GetByID error", ecode.StatusInvalidParams, err, "req", req)
	}

	//查询当前角色下的所有用户
	userParams := &query.Params{Columns: []query.Column{
		{
			Name:  "roleId",
			Value: roleInfo.ID, //字符串包含匹配，todo
		},
	}}
	userInfo, _, err := s.userDao.GetByColumns(context.Background(), userParams)
	if err != nil && !errors.Is(err, model.ErrRecordNotFound) {
		return operationFailed, log(ctx, "manager.RoleDelete.userDao.GetByColumns error", ecode.StatusInvalidParams, err, "req", req)
	}

	//开始事务
	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for _, user := range userInfo {
		roleIds, _ := common.IdToArray(user.RoleId)
		for i := 0; i < len(roleIds); i++ {
			if roleIds[i] == roleInfo.ID {
				roleIds = append(roleIds[:i], roleIds[i+1:]...)
				var roleId string
				roleId, err = common.ArrayToId(roleIds)
				if err != nil {
					return operationFailed, log(ctx, "manager.RoleDelete.UserUpdate to json error", ecode.StatusInvalidParams, err, "req", req)
				}
				user.RoleId = roleId
				err = s.userDao.UpdateByTx(context.Background(), tx, user)
				if err != nil {
					return operationFailed, log(ctx, "manager.RoleDelete.UserUpdate error", ecode.StatusInvalidParams, err, "req", req)
				}
			}
		}
	}

	//删除角色
	err = s.roleDao.DeleteByTx(context.Background(), tx, roleInfo.ID)
	if err != nil {
		return operationFailed, log(ctx, "manager.RoleDelete.roleDao.DeleteByTx error", ecode.StatusInternalServerError, err, "req", req)
	}
	//提交事务
	err = tx.Commit().Error
	if err != nil {
		return operationFailed, log(ctx, "manager.ApiDelete.Commit error", ecode.StatusInternalServerError, err, "req", req)
	}

	return operationSuccess, nil
}

// ApiCreate 接口创建
func (s *manager) ApiCreate(ctx context.Context, req *managerV1.ApiCreateReq) (*managerV1.OperationRes, error) {

	// fill in the business logic code here
	// example:
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//检查此api是否存在
	apiCondition := &query.Conditions{Columns: []query.Column{
		{
			Name:  "name",
			Value: req.Name,
			Logic: "or",
		},
		{
			Name:  "url",
			Value: req.Url,
			Logic: "and",
		},
		{
			Name:  "method",
			Value: req.Method,
			Logic: "and",
		},
	}}
	_, err = s.apiDao.GetByCondition(context.Background(), apiCondition)
	if !errors.Is(err, model.ErrRecordNotFound) {
		return operationFailed, log(ctx, "manager.ApiCreate.apiDao.GetByCondition error", ecode.StatusInvalidParams, err, "", nil)
	}

	api := &model.Api{
		Name:   req.Name,
		Url:    req.Url,
		Method: req.Method,
		Desc:   req.Desc,
	}
	err = s.apiDao.Create(context.Background(), api)
	if err != nil {
		return operationFailed, log(ctx, "manager.ApiCreate.apiDao.Create error", ecode.StatusInternalServerError, err, "", nil)
	}

	return operationSuccess, nil
}

// ApiGet 接口查询
func (s *manager) ApiGet(ctx context.Context, req *managerV1.ApiGetReq) (*managerV1.ApiGetRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	apiParams := s.buildApiDaoParams(req)
	apiInfo, count, err := s.apiDao.GetByColumns(context.Background(), apiParams)
	if err != nil && !errors.Is(err, model.ErrRecordNotFound) {
		return nil, log(ctx, "manager.ApiCreate.apiDao.Create error", ecode.StatusInternalServerError, err, "req", req)
	}

	return &managerV1.ApiGetRes{
		Api:   s.buildApiGetRes(apiInfo),
		Count: uint64(count),
	}, nil
}

// ApiUpdate 接口修改
func (s *manager) ApiUpdate(ctx context.Context, req *managerV1.ApiUpdateReq) (*managerV1.OperationRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//查询当前接口是否存在
	apiInfo, err := s.apiDao.GetByID(context.Background(), req.Id)
	if err != nil {
		return operationFailed, log(ctx, "manager.ApiUpdate.apiDao.GetByID error", ecode.StatusInvalidParams, err, "req", req)
	}

	//校验修改后的接口是否与其他接口冲突
	db := model.GetDB()
	var apiInfos = []*model.Api{}
	err = db.Table("api").Where("name = ? or (url = ? and method = ?)", req.Name, req.Url, req.Method).Find(&apiInfos).Error
	if err != nil && !errors.Is(err, model.ErrRecordNotFound) {
		return operationFailed, log(ctx, "manager.ApiUpdate.apiDao.GetByCondition error", ecode.StatusInvalidParams, err, "req", req)
	} else {
		if len(apiInfos) > 1 {
			return operationFailed, log(ctx, "manager.ApiUpdate.apiDao.GetByCondition error", ecode.StatusInvalidParams, err, "req", req)
		}
	}

	apiInfo.Name = req.Name
	apiInfo.Url = req.Url
	apiInfo.Method = req.Method
	apiInfo.Desc = req.Desc

	//更新接口
	err = s.apiDao.UpdateByID(context.Background(), apiInfo)
	if err != nil {
		return operationFailed, log(ctx, "manager.ApiUpdate.apiDao.UpdateByID error", ecode.StatusInternalServerError, err, "req", req)
	}

	return operationSuccess, nil
}

// ApiDelete 接口删除
func (s *manager) ApiDelete(ctx context.Context, req *managerV1.ApiDeleteReq) (*managerV1.OperationRes, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//查询当前接口是否存在
	apiInfo, err := s.apiDao.GetByID(context.Background(), req.Id)
	if err != nil {
		return operationFailed, log(ctx, "manager.ApiUpdate.apiDao.GetByID error", ecode.StatusInvalidParams, err, "req", req)
	}

	//查询关联的当前接口的角色
	roleParams := &query.Params{Columns: []query.Column{
		{
			Name:  "api",
			Value: apiInfo.ID, //字符串包含匹配，todo
			Exp:   "like",
		},
	}}
	roleInfo, _, err := s.roleDao.GetByColumns(context.Background(), roleParams)
	if err != nil && !errors.Is(err, model.ErrRecordNotFound) {
		return operationFailed, log(ctx, "manager.ApiUpdate.roleDao.GetByColumns error", ecode.StatusInvalidParams, err, "apiInfo", apiInfo)
	}

	//更新接口
	//开启事务
	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for _, role := range roleInfo {
		apiIds, _ := common.IdToArray(role.Api)
		for i := 0; i < len(apiIds); i++ {
			if apiIds[i] == apiInfo.ID {
				apiIds = append(apiIds[:i], apiIds[i+1:]...)
				var apiId string
				apiId, err = common.ArrayToId(apiIds)
				if err != nil {
					return operationFailed, log(ctx, "manager.ApiDelete.RoleUpdate to json error", ecode.StatusInvalidParams, err, "req", req)
				}
				role.Api = apiId
				err = s.roleDao.UpdateByTx(ctx, tx, role)
				if err != nil {
					return operationFailed, log(ctx, "manager.ApiDelete.RoleUpdate error", ecode.StatusInvalidParams, err, "req", req)
				}
			}
		}
	}

	//删除接口
	err = s.apiDao.DeleteByID(context.Background(), apiInfo.ID)
	if err != nil {
		return operationFailed, log(ctx, "manager.ApiDelete.DeleteByID error", ecode.StatusInternalServerError, err, "req", req)
	}
	//提交事务
	err = tx.Commit().Error
	if err != nil {
		return operationFailed, log(ctx, "manager.ApiDelete.Commit error", ecode.StatusInternalServerError, err, "req", req)
	}

	return operationSuccess, nil
}

// DeviceCreate 设备创建
func (s *manager) DeviceCreate(ctx context.Context, req *managerV1.DeviceCreateReq) (*managerV1.OperationRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//检查设备类型是否存在
	_, err = s.deviceTypeDao.GetByID(context.Background(), req.DeviceTypeId)
	if err != nil {
		return operationFailed, log(ctx, "deviceType is not exist", ecode.StatusInvalidParams, err, "req", req)
	}

	//检查设备名和编号是否重复
	deviceCondition := &query.Conditions{Columns: []query.Column{
		{
			Name:  "name",
			Value: req.Name,
			Logic: "or",
		},
		{
			Name:  "code",
			Value: req.Code,
			Logic: "or",
		},
	}}
	if _, err = s.deviceDao.GetByCondition(context.Background(), deviceCondition); err == nil {
		return operationFailed, log(ctx, "device is exist", ecode.StatusInvalidParams, err, "req", req)
	}

	db := model.GetDB()
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	device := &model.Device{
		Name:         req.Name,
		Code:         req.Code,
		CheckTime:    float64(req.CheckTime),
		DeviceTypeID: req.DeviceTypeId,
		State:        common.Running.Int(),
	}
	_, err = s.deviceDao.CreateByTx(context.Background(), tx, device)
	if err != nil {
		logger.Warn("DeviceCreate error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, log(ctx, "device create error", ecode.StatusInvalidParams, err, "req", req)
	}

	//创建对应的设备数据表
	err = s.deviceDataModelDao.CreateDeviceDataTable(context.Background(), tx, req.Code)
	if err != nil {
		logger.Warn("DeviceCreate error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, log(ctx, "manager.deviceCreate common.CreateDeviceDataTable error", ecode.StatusInvalidParams, err, "req", req)
	}

	err = tx.Commit().Error
	if err != nil {
		logger.Warn("tx.Commit error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, log(ctx, "tx.Commit error", ecode.StatusInvalidParams, err, "req", req)
	}

	return operationSuccess, nil
}

// DeviceGet 设备查询
func (s *manager) DeviceGet(ctx context.Context, req *managerV1.DeviceGetReq) (*managerV1.DeviceGetRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//构建查询参数
	params, err := s.buildDeviceDaoParams(context.Background(), req)
	if err != nil {
		return nil, log(ctx, "deviceType error", ecode.StatusInvalidParams, err, "", nil)
	}

	deviceInfo, count, err := s.deviceDao.GetByColumns(context.Background(), params)
	if err != nil {
		logger.Warn("DeviceGet error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	device, err := s.buildDeviceGetRes(ctx, deviceInfo)
	if err != nil {
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &managerV1.DeviceGetRes{
		Device: device,
		Count:  uint64(count),
	}, nil
}

// DeviceUpdate 设备修改
func (s *manager) DeviceUpdate(ctx context.Context, req *managerV1.DeviceUpdateReq) (*managerV1.OperationRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//查询设备是否存在
	deviceInfo, err := s.deviceDao.GetByID(context.Background(), req.Id)
	if err != nil {
		return operationFailed, log(ctx, "manager.deviceDao.GetByID error", ecode.StatusInvalidParams, err, "", nil)
	}

	//deviceInfo.DeviceTypeID=req.DeviceTypeName
	if req.DeviceTypeId != deviceInfo.DeviceTypeID {
		_, err = s.deviceTypeDao.GetByID(context.Background(), req.DeviceTypeId)
		if err != nil {
			return operationFailed, log(ctx, "manager.deviceTypeDao.GetByID error", ecode.StatusInvalidParams, err, "", nil)
		}
	}
	//校验
	deviceCondition := &query.Conditions{Columns: []query.Column{
		{
			Name:  "name",
			Value: req.Name,
			Logic: "||",
		},
		{
			Name:  "code",
			Value: req.Name,
			Logic: req.Code,
		},
	}}
	_, err = s.deviceDao.GetByCondition(context.Background(), deviceCondition)
	if errors.Is(err, model.ErrRecordNotFound) {
		return operationFailed, log(ctx, "manager.deviceDao.GetByCondition error", ecode.StatusInvalidParams, err, "", nil)
	}
	deviceInfo.Name = req.Name
	deviceInfo.Code = req.Code
	deviceInfo.DeviceTypeID = req.DeviceTypeId
	deviceInfo.CheckTime = float64(req.CheckTime)
	err = s.deviceDao.UpdateByID(context.Background(), deviceInfo)
	if err != nil {
		logger.Warn("DeviceUpdate error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, log(ctx, "manager.deviceDao.UpdateByID error", ecode.StatusInvalidParams, err, "", nil)
	}

	return operationSuccess, nil
}

// DeviceDelete 设备删除
func (s *manager) DeviceDelete(ctx context.Context, req *managerV1.DeviceDeleteReq) (*managerV1.OperationRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//查询设备是否存在
	deviceInfo, err := s.deviceDao.GetByID(context.Background(), req.Id)
	if err != nil {
		return operationFailed, log(ctx, "manager.deviceDao.GetByID error", ecode.StatusInvalidParams, err, "", nil)
	}

	err = s.deviceDao.DeleteByID(context.Background(), req.Id)
	if err != nil {
		logger.Warn("DeviceDelete error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, ecode.StatusInternalServerError.Err()
	}
	//删除对应设备数据表
	err = s.deviceDataModelDao.DropDeviceTable(context.Background(), model.GetDB(), deviceInfo.Code)
	if err != nil {
		logger.Warn("DeviceDelete error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, ecode.StatusInternalServerError.Err()
	}

	return operationSuccess, nil
}

// DeviceTypeCreate 设备类型创建
func (s *manager) DeviceTypeCreate(ctx context.Context, req *managerV1.DeviceTypeCreateReq) (*managerV1.OperationRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//检查是否存在
	deviceTypeCondition := &query.Conditions{Columns: []query.Column{
		{
			Name:  "name",
			Value: req.Name,
		},
	}}
	if _, err = s.deviceTypeDao.GetByCondition(context.Background(), deviceTypeCondition); err == nil {
		return operationFailed, log(ctx, "manager.deviceDao.GetByID error", ecode.StatusInvalidParams, err, "", nil)
	}

	err = s.deviceTypeDao.Create(context.Background(), &model.DeviceType{
		Name:     req.Name,
		Describe: req.Describe,
	})
	if err != nil {
		logger.Warn("DeviceTypeCreate error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, ecode.StatusInternalServerError.Err()
	}

	return operationSuccess, nil
}

// DeviceTypeGet 设备类型查询
func (s *manager) DeviceTypeGet(ctx context.Context, req *managerV1.DeviceTypeGetReq) (*managerV1.DeviceTypeGetRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	params := s.buildDeviceTypeDaoParams(req)

	deviceTypeInfo, count, err := s.deviceTypeDao.GetByColumns(context.Background(), params)
	if err != nil {
		logger.Warn("DeviceTypeGet error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.Err()
	}

	return &managerV1.DeviceTypeGetRes{
		DeviceType: s.buildDeviceTypeGetRes(deviceTypeInfo),
		Count:      uint64(count),
	}, nil
}

// DeviceTypeUpdate 设备类型修改
func (s *manager) DeviceTypeUpdate(ctx context.Context, req *managerV1.DeviceTypeUpdateReq) (*managerV1.OperationRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//查询是否存在
	deviceTypeInfo, err := s.deviceTypeDao.GetByID(context.Background(), req.Id)
	if err != nil {
		return operationFailed, log(ctx, "manager.deviceTypeDao.GetByID error", ecode.StatusInvalidParams, err, "", nil)
	}
	if deviceTypeInfo.Name != req.Name {
		//校验，todo
	}

	//db := model.GetDB()
	//tx := db.Begin()
	//defer func() {
	//	if err != nil {
	//		tx.Rollback()
	//	}
	//}()
	//if deviceTypeInfo.Protocol != req.Protocol {
	//	//查询属于当前类型的非自定义协议的设备
	//	deviceParams := &query.Params{Columns: []query.Column{
	//		{
	//			Name:  "device_type_id",
	//			Value: deviceTypeInfo.ID,
	//		},
	//		{
	//			Name:  "is_custom",
	//			Value: common.CustomProtocol.Int(),
	//		},
	//	}}
	//	var deviceInfo []*model.Device
	//	deviceInfo, _, err = s.deviceDao.GetByColumns(ctx, deviceParams)
	//	if err != nil {
	//		return operationFailed, log(ctx, "manager.deviceDao.GetByColumns error", ecode.StatusInvalidParams, err, "", nil)
	//	}
	//
	//	for _, device := range deviceInfo {
	//		//device.Protocol = req.Protocol
	//		err = s.deviceDao.UpdateByTx(ctx, tx, device)
	//		if err != nil {
	//			return operationFailed, log(ctx, "manager.deviceDao.UpdateByTx error", ecode.StatusInvalidParams, err, "", nil)
	//		}
	//	}
	//
	//	deviceTypeInfo.Protocol = req.Protocol
	//}

	deviceTypeInfo.Name = req.Name
	deviceTypeInfo.Describe = req.Describe

	err = s.deviceTypeDao.UpdateByID(context.Background(), deviceTypeInfo)
	if err != nil {
		logger.Warn("DeviceTypeUpdate error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, log(ctx, "manager.tx.Commit() error", ecode.StatusInvalidParams, err, "", nil)
	}

	//err = tx.Commit().Error
	//if err != nil {
	//	return nil, log(ctx, "manager.tx.Commit() error", ecode.StatusInvalidParams, err, "", nil)
	//}

	return operationSuccess, nil
}

// DeviceTypeDelete 设备类型删除
func (s *manager) DeviceTypeDelete(ctx context.Context, req *managerV1.DeviceTypeDeleteReq) (*managerV1.OperationRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//检查设备类型是否存在
	deviceTypeInfo, err := s.deviceTypeDao.GetByID(ctx, req.Id)
	if err != nil {
		return operationFailed, log(ctx, "manager.deviceTypeDao.GetByID() error", ecode.StatusInvalidParams, err, "", nil)
	}

	err = s.deviceTypeDao.DeleteByID(ctx, deviceTypeInfo.ID)
	if err != nil {
		logger.Warn("DeviceTypeDelete error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, ecode.StatusInternalServerError.Err()
	}

	return operationSuccess, nil
}

// UserDeviceCreate 用户设备信息创建
func (s *manager) UserDeviceCreate(ctx context.Context, req *managerV1.UserDeviceCreateReq) (*managerV1.OperationRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//检查用户是否存在
	userCondition := &query.Conditions{Columns: []query.Column{
		{
			Name:  "account",
			Value: req.Account,
		},
	}}
	userInfo, err := s.userDao.GetByCondition(ctx, userCondition)
	if err != nil {
		return operationFailed, log(ctx, "manager.UserDeviceCreate.userDao.GetByCondition err", ecode.StatusInvalidParams, err, "", nil)
	}

	//检查设备是否存在
	deviceCondition := &query.Conditions{Columns: []query.Column{
		{
			Name:  "code",
			Value: req.DeviceCode,
		},
	},
	}
	deviceInfo, err := s.deviceDao.GetByCondition(ctx, deviceCondition)
	if err != nil {
		return operationFailed, log(ctx, "manager.UserDeviceCreate.deviceDao.GetByCondition err", ecode.StatusInvalidParams, err, "", nil)
	}

	//检查当前记录是否已经添加
	userDeviceCondition := &query.Conditions{Columns: []query.Column{
		{
			Name:  "user_id",
			Value: userInfo.ID,
		},
		{
			Name:  "device_id",
			Value: deviceInfo.ID,
		},
	}}
	_, err = s.userDeviceDao.GetByCondition(ctx, userDeviceCondition)
	if !errors.Is(err, model.ErrRecordNotFound) {
		return operationFailed, log(ctx, "manager.UserDeviceCreate.deviceDao.GetByCondition err", ecode.StatusInvalidParams, errors.New("record is exist"), "", nil)
	}

	err = s.userDeviceDao.Create(ctx, &model.UserDevice{
		UserID:     userInfo.ID,
		UserName:   userInfo.Name,
		DeviceID:   deviceInfo.ID,
		DeviceCode: deviceInfo.Code,
		DeviceName: deviceInfo.Name,
	})
	if err != nil {
		logger.Warn("UserDeviceCreate error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, log(ctx, "manager.UserDeviceCreate.userDeviceDao.Create err", ecode.StatusInvalidParams, err, "", nil)
	}

	return operationSuccess, nil
}

// UserDeviceGet 用户设备信息查询
func (s *manager) UserDeviceGet(ctx context.Context, req *managerV1.UserDeviceGetReq) (*managerV1.UserDeviceGetRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	params := s.buildUserDeviceDaoParams(req)
	userDeviceInfo, count, err := s.userDeviceDao.GetByColumns(context.Background(), params)
	if err != nil {
		logger.Warn("UserDeviceGet error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, log(ctx, "manager.UserDeviceGet.userDeviceDao.GetByColumns err", ecode.StatusInvalidParams, err, "", nil)
	}

	return &managerV1.UserDeviceGetRes{
		UserDevice: s.buildUserDeviceGetRes(userDeviceInfo),
		Count:      uint64(count),
	}, nil
}

// UserDeviceDelete 用户设备信息删除
func (s *manager) UserDeviceDelete(ctx context.Context, req *managerV1.UserDeviceDeleteReq) (*managerV1.OperationRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//校验
	_, err = s.userDeviceDao.GetByID(ctx, req.Id)
	if err != nil {
		logger.Warn("UserDeviceDelete error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, log(ctx, "manager.UserDeviceDelete.userDeviceDao.GetByID err", ecode.StatusInvalidParams, err, "", nil)
	}

	err = s.userDeviceDao.DeleteByID(ctx, req.Id)
	if err != nil {
		return operationFailed, log(ctx, "manager.UserDeviceDelete.userDeviceDao.DeleteByID err", ecode.StatusInvalidParams, err, "", nil)
	}

	return operationSuccess, nil
}

// DeviceDataGet 设备数据查询
func (s *manager) DeviceDataGet(ctx context.Context, req *managerV1.DeviceDataGetReq) (*managerV1.DeviceDataGetRes, error) {
	panic("implement me")

	// fill in the business logic code here
	// example:
	//	    err := req.Validate()
	//	    if err != nil {
	//		    logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
	//		    return nil, ecode.StatusInvalidParams.Err()
	//	    }
	// 	ctx = interceptor.WrapServerCtx(ctx)
	//
	// 	reply, err := s.iDao.DeviceDataGet(ctx, &model.Manager{
	//     	Id: req.Id,
	//     	Name: req.Name,
	//     	Code: req.Code,
	//     })
	// 	if err != nil {
	//			logger.Warn("DeviceDataGet error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
	//			return nil, ecode.StatusInternalServerError.Err()
	//		}
	//
	//     return &managerV1.DeviceDataGetRes{
	//     	Data: reply.Data,
	//     }, nil
}

// DeviceDataReceive 设备数据接收
func (s *manager) DeviceDataReceive(ctx context.Context, req *managerV1.DeviceDataReceiveReq) (*managerV1.OperationRes, error) {

	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	//设备数据解析
	parseData, err := s.deviceDataParse(ctx, req.Data)
	if err != nil {
		logger.Warn("DeviceDataReceive error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, ecode.StatusInternalServerError.Err()
	}

	//存储设备解析后的数据
	err = s.saveDeviceData(ctx, parseData)
	if err != nil {
		logger.Warn("DeviceDataReceive error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return operationFailed, ecode.StatusInternalServerError.Err()
	}

	return operationSuccess, nil
}
