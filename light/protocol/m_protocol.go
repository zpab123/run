package protocol

// 主消息ID
const (
	C_MID_HANDSHAKE     uint16 = iota + 1 // 握手消息
	C_MID_HANDSHAKE_ACK                   // 客户端->服务器握手ACK
	C_MID_HEARTBEAT                       // 心跳消息
	C_MID_KICK                            // 服务器主动断开
	C_MID_DATA                            // 数据
)

// 子消息ID
const (
//C_SID_ uint16 = iota + 1 // 握手消息
)
