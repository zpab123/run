package protocol

import (
	"fmt"
	"io"
	"net/http"
)

// 消息头
type MessageHead struct {
	Mt  int // 主协议
	Sub int // 子协议
}
