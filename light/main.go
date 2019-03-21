package main

import (
	"net/http"

	"golang.org/x/net/websocket" // websocket 库
)

func main() {

	// 注册路由
	http.Handle("/ws", websocket.Handler(wsReq))

	// 开启服务器
	http.ListenAndServe(":8888", nil)
}
