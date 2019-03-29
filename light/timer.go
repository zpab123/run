package main

import (
	"fmt"
	"time"
)

func runTimer() {
	for {
		select {
		case <-time.After(10 * time.Second):
			sessionMap.Range(checkHeart)
		}
	}
}

// 遍历函数
func checkHeart(id interface{}, ses interface{}) bool {
	var ok bool
	var session *Session

	_, ok = id.(uint32)
	if !ok {
		return false
	}

	session, ok = ses.(*Session)
	if !ok {
		return false
	}

	fmt.Println(session.id)

	return true
}
