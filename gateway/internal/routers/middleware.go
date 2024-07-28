package routers

import "gateway/internal/middleware"

var (
	websocket = middleware.NewWebSocket()
	authority = middleware.NewAuthority()
	jwtToken  = middleware.NewJwtToken()
	cors      = middleware.Cors()
)
