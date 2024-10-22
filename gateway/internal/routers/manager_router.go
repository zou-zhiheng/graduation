// Code generated by https://github.com/zhufuyi/sponge

package routers

import (
	"context"

	gatewayV1 "gateway/api/gateway/v1"
	"gateway/internal/service"

	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/logger"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func init() {
	allMiddlewareFns = append(allMiddlewareFns, func(c *middlewareConfig) {
		managerMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			managerRouter(r, groupPathMiddlewares, singlePathMiddlewares, service.NewManagerClient())
		})
}

func managerRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService gatewayV1.ManagerLogicer) {
	ctxFn := func(c *gin.Context) context.Context {
		md := metadata.New(map[string]string{
			// set metadata to be passed from http to rpc
			middleware.ContextRequestIDKey: middleware.GCtxRequestID(c), // request_id
			//middleware.HeaderAuthorizationKey: c.GetHeader(middleware.HeaderAuthorizationKey),  // authorization
		})
		return metadata.NewOutgoingContext(c.Request.Context(), md)
	}

	gatewayV1.RegisterManagerRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		gatewayV1.WithManagerRPCResponse(),
		gatewayV1.WithManagerLogger(logger.Get()),
		gatewayV1.WithManagerRPCStatusToHTTPCode(
		// Set some error codes to standard http return codes,
		// by default there is already ecode.StatusInternalServerError and ecode.StatusServiceUnavailable
		// example:
		// 	ecode.StatusUnimplemented, ecode.StatusAborted,
		),
		gatewayV1.WithManagerWrapCtx(ctxFn),
	)
}

// you can set the middleware of a route group, or set the middleware of a single route,
// or you can mix them, pay attention to the duplication of middleware when mixing them,
// it is recommended to set the middleware of a single route in preference
func managerMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/manager/:id  will take effect
	// c.setGroupPath("/api/v1/manager", middleware.Auth())

	c.setGroupPath("/api/v1/user", jwtToken.JWTAuth(), authority.ApiAuth())
	c.setGroupPath("/api/v1/api", jwtToken.JWTAuth(), authority.ApiAuth())
	c.setGroupPath("/api/v1/role", jwtToken.JWTAuth(), authority.ApiAuth())
	c.setGroupPath("/api/v1/device", jwtToken.JWTAuth(), authority.ApiAuth())
	c.setGroupPath("/api/v1/deviceType", jwtToken.JWTAuth(), authority.ApiAuth())
	c.setGroupPath("/api/v1/userDevice", jwtToken.JWTAuth(), authority.ApiAuth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("POST", "/api/v1/login", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/user/register", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/user/get", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/user/update", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/user/resetPassword", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/user/delete", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/role/create", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/role/get", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/role/update", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/role/delete", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/api/create", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/api/get", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/api/update", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/api/delete", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/userDevice/create", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/userDevice/get", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/userDevice/delete", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/device/create", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/device/get", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/device/update", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/device/delete", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/device/dateReceive", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/deviceType/create", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/deviceType/get", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/deviceType/update", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/deviceType/delete", middleware.Auth())
	//c.setSinglePath("GET", "/api/v1/device/dataPush", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/device/dataPush", middleware.Auth())
}
