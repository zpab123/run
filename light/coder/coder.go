package coder

import (
	"encoding/json"
	"fmt"
)

type Packet struct {
	Type   uint8    // 类型
	Length uint16   // 长度
	Msg    *Message // 消息
}

type Message struct {
	Mtype uint16 // 主命令
	Stype uint16 // 子命令
	Data  []byte // 消息
}

// 编码成 Message
func EncodeMsg(mtype uint16, stype uint16, data interface{}) *Message {
	b, err := json.Marshal(data)

	if nil != err {
		fmt.Println(err)
		return
	}

	msg := &Message{
		Mtype: mtype,
		Stype: stype,
		Data:  b,
	}

	return msg
}

// 编码成 Message
func EncodePacket(dataType uint8, msg *Message) {

}
