package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket" // websocket 库
)

// 连接创建
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域访问
	},
}

// 根路由
func rootReq(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "你好！")
	io.WriteString(w, "你好！")
}

// websocket 路由
func wsReq(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if nil != err {
		fmt.Println("创建 websocket.conn 错误：", err)

		return
	}

	ses := &Session{
		Connection: conn,
		Md5:        "",
	}

	ses.Recv()
}
