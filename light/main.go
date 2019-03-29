package main

import (
	"fmt"
	"net/http"
)

func main() {
	//go runTimer()

	// 注册路由
	http.HandleFunc("/ws", wsReq)
	http.HandleFunc("/", rootReq)

	// 开启服务器
	addr := "127.0.0.1:8888"

	fmt.Println("开启服务器", addr)
	err := http.ListenAndServe(addr, nil)

	if err != err {
		fmt.Println("服务器监听错误", err)
	}
}
