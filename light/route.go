package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/gorilla/websocket" // websocket 库
	"github.com/zpab123/syncutil"  // 同步变量
	"github.com/zpab123/zaplog"    // log
)

var (
	// 创建 websocket
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许跨域访问
		},
	}

	// 连接对象保存
	sessionMap = &sync.Map{}

	// sessionID
	sessionID syncutil.AtomicUint32
)

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
		id:         uint32(sessionID),
	}

	sessionID.Add(1)

	zaplog.Debugf("收到新的 websocket 连接: %s", conn.RemoteAddr())

	// 保存连接
	sessionMap.Store(ses.id, ses)

	ses.Recv()
}
