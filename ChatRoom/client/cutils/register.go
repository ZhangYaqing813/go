package cutils

import (
	"encoding/json"
	"fmt"
	"go/ChatRoom/public"
	"net"
)

func Register(userID int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println("注册服务打开链接失败 err = ", err)
	}

	var msg public.Messages      // 定义一个消息类型为 public.Messages 用于封装注册信息
	var userRegMsg public.RegMsg //定义一个消息类型为 public.RegMsg 用于封装用户注册时的消息
	var userRegresMsg public.LoginReMsg
	msg.Type = public.RegMsgType
	userRegMsg.UserID = userID
	userRegMsg.UserPwd = userPwd
	data, err := json.Marshal(userRegMsg)
	msg.Data = string(data)

	err = public.SendMsg(conn, msg)
	if err != nil {
		fmt.Println("发送注册信息失败")
		return
	}

	msg, err = public.RedMsg(conn)
	if err != nil {
		fmt.Println("获取注册返回信息失败err ", err)
		return
	}
	fmt.Println("register debug")
	err = json.Unmarshal([]byte(msg.Data), &userRegresMsg)
	if err != nil {
		fmt.Println("解析注册返回码失败err", err)
		return
	}
	if userRegresMsg.Code == 200 {
		fmt.Println("UserID 注册成功")
	} else if userRegresMsg.Code == 600 {
		fmt.Println("该用户已注册")
	}
	return
}
