package sutils

import (
	"encoding/json"
	"fmt"
	"go/ChatRoom/public"
	"net"
)

func SeProcess(conn net.Conn) {
	defer conn.Close()
	fmt.Println("开始接收数据")
	msg, err := public.RedMsg(conn)
	if err != nil {
		fmt.Println("数据返回错误")
	}

	fmt.Println(msg)
	switch msg.Type {
	case public.LoginMsgType:
		fmt.Println("用户登陆校验")
		err = LoginPrec(conn, msg.Data)
		if err != nil {
			fmt.Println("回复信息失败")
		}
	default:
		fmt.Println("暂时无法处理")
	}
}

func LoginPrec(conn net.Conn, msg string) (err error) {
	var userMsg public.LoginMsg
	var logmsg public.LoginReMsg
	var remsg public.Messages

	json.Unmarshal([]byte(msg), &userMsg)
	remsg.Type = public.LoginReMsgType
	if userMsg.UserID == 100 && userMsg.UserPwd == "abc" {
		logmsg.Code = 200
	} else {
		logmsg.Code = 500
		logmsg.Error = "用户名密码错误"
	}
	data, err := json.Marshal(logmsg)
	remsg.Data = string(data)

	err = public.Response(conn, remsg)
	if err != nil {
		fmt.Println("消息回复失败")
	}
	return
}
