package public

const (
	LoginMsgType = "LoginMsg"
	RegMsgType   = "RegMsgType"
)

//登录是向server 提交的用户信息
type LoginMsg struct {
	// 用户Id
	UserID int `json:"userid"`
	//用户密码
	UserPwd string `json:"userpwd"`
	//用户名
	UserName string `json:"username"`
}

// 用户登录状态信息
type LoginReMsg struct {
	//登录状态码
	Code int `json:"code"`
	// 登录错误时的错误信息
	Error string `json:"Error"`
}

//要发送的信息
type Messages struct {
	//发送的消息类型
	Type string `json:"type"`
	// 发送信息的内容
	Data string `json:"data"`
}

type RegMsg struct { // 用户注册消息的结构体
	// 用户Id
	UserID int `json:"userid"`
	//用户密码
	UserPwd string `json:"userpwd"`
	//用户名
	UserName string `json:"username"`
}
