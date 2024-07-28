package common

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func GetWsCoon(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
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
	return UpGrader.Upgrade(w, r, responseHeader)
}
