package protocol

// 消息类型
const (
	C_TYPE_HANDSHAKE     uint16 = iota + 1 // 握手消息
	C_TYPE_HANDSHAKE_ACK                   // 客户端->服务器握手ACK
	C_TYPE_HEARTBEAT                       // 心跳消息
	C_TYPE_DATA                            // 数据
	C_TYPE_KICK                            // 服务器主动断开
)

// 主协议常量
const (
	C_MTYPE_DATA     uint16 = iota // 主协议1
	C_MTYPE_LOGIN                  // 主协议2
	C_MTYPE_REGISTER               // 主协议3
)

// 自协议常量
const (
	C_STYPE_DATA  uint16 = iota // 子协议1
	C_STYPEDATA1                // 子协议2
	C_STYPE_DATA2               // 子协议3
)
