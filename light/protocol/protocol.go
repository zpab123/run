package protocol

// 客户端 -> 服务器 登录请求
type LoginReq struct {
	Account  string // 帐号
	Password string // 密码
	Email    string // 邮箱
}

// 服务器 -> 客户端 登录结果
type LoginRes struct {
	Code  uint32
	Email string // 邮箱
}

type Data struct {
	Zp string `json:"zp"`
	Cd string `json:"cd"`
	Ef string `json:"ef"`
}
