package main

import (
	"io"
	"net/http"

	"golang.org/x/net/websocket" // websocket 库
)

// 根路由
func rootReq(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "你好！")
	io.WriteString(w, "你好！")
}

// websocket 路由
func wsReq(conn *websocket.Conn) {

}
