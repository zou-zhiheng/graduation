package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type WebSocket struct {
}

func NewWebSocket() *WebSocket {
	return &WebSocket{}
}

type HandlerFunc func(ctx *gin.Context) (res interface{}, err error)

//func(w *WebSocket) Upgrade() gin.HandlerFunc {
//	return func(context *gin.Context) {
//
//	}
//}

// Upgrade
//
//	@Description: 升级为WebSocket请求
//	@Author zzh
//	@receiver w
//	@param ctx
//	@param handlerFunc
func (w *WebSocket) Upgrade(handlerFunc HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if handlerFunc == nil {
			ctx.Abort()
			return
		}

		//升级为webSocket请求
		ws, err := w.GetWsCoon(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.Abort()
			return
		}

		for {
			var res interface{}
			res, err = handlerFunc(ctx)
			if err != nil {
				break
			}
			var byteData []byte
			byteData, err = json.Marshal(res)
			if err != nil {
				break
			}
			if err = ws.WriteMessage(websocket.TextMessage, byteData); err != nil {
				break
			}

			time.Sleep(time.Minute)
		}

		_ = ws.Close()
		ctx.Abort()
	}
}

func (w *WebSocket) GetWsCoon(writer http.ResponseWriter, req *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	UpGrader := websocket.Upgrader{
		// 读取缓存大小
		ReadBufferSize: 1024,
		// 写入缓存大小
		WriteBufferSize: 1024,
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//升级http请求
	return UpGrader.Upgrade(writer, req, responseHeader)
}
