package main

import (
	"github.com/gorilla/websocket" // websocket 库
)

type robot struct {
}

// 登录
func (this *robot) login() {
	//ws, err = websocket.Dialer.Dial("127.0.0.1:8888/ws")
}

// 发送心跳
func (this *robot) sendHearBeat() {

}
