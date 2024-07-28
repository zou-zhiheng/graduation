package middleware

import (
	"encoding/json"
	"fmt"
	"gateway/internal/cache"
	"gateway/internal/dao"
	"gateway/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"net/http"
)

type Authority struct {
}

func NewAuthority() *Authority {
	return &Authority{}
}

func (a *Authority) ApiAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取访问api的url和方法
		method, url := ctx.Request.Method, ctx.Request.URL.Path
		var api *model.Api

		apiCondition := &query.Conditions{Columns: []query.Column{
			{
				Name:  "url",
				Value: url,
			},
			{
				Name:  "method",
				Value: method,
			},
		}}
		api, err := dao.NewApiDao(model.GetDB(), cache.NewApiCache(model.GetCacheType())).GetByCondition(ctx, apiCondition)
		if err != nil {
			logger.Error(fmt.Sprintf("api not found:{ url:%s, methoed:%s}", url, method), interceptor.ServerCtxRequestIDField(ctx))
			ctx.Abort()
			return
		}

		//获取token解析出来的user
		userInterface, _ := ctx.Get("user")
		user := userInterface.(model.User)
		//获取user对应的role
		db := model.GetDB()
		_ = json.Unmarshal([]byte(user.RoleId), &user.RoleIds)
		var i int
		for i = 0; i < len(user.RoleIds); i++ {
			db = db.Or("id = ?", user.RoleIds[i])
		}
		role := []*model.Role{}
		err = db.Limit(len(user.RoleIds)).Find(&role).Error
		if err != nil {
			ctx.JSON(http.StatusOK, "此用户无角色")
			ctx.Abort()
			return
		}

		apiMap := make(map[uint64]bool)
		var ok bool
		for i = 0; i < len(role); i++ {
			_ = json.Unmarshal([]byte(role[i].Api), &role[i].Apis)
			for j := range role[i].Apis {
				if _, ok = apiMap[role[i].Apis[j]]; !ok {
					apiMap[role[i].Apis[j]] = true
					_, ok = apiMap[api.ID] //判断权限是否存在
					if ok {
						ctx.Next()
						return
					}
				}
			}
		}

		ctx.JSON(http.StatusOK, "验证api：此用户无访问此api的权限")
		ctx.Abort()
		return
	}
}
