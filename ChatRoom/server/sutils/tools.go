package sutils

import (
	"encoding/json"
	"fmt"
	"go/ChatRoom/public"
	"net"
)

func SeProcess(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭链接失败")
		}
	}(conn)
	fmt.Println("开始接收数据")
	fmt.Printf("public.RedMsg", public.RedMsg)
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
	case public.RegMsgType:
		fmt.Println("新用户注册=========")
		err = RegPrec(conn, msg.Data)

	default:
		fmt.Println("暂时无法处理")
	}
}

func LoginPrec(conn net.Conn, msg string) (err error) {
	var userMsg public.LoginMsg
	var logmsg public.LoginReMsg
	var remsg public.Messages

	json.Unmarshal([]byte(msg), &userMsg)
	remsg.Type = "LoginReMsgType"
	code := public.RedisGet(userMsg.UserID, userMsg.UserPwd)

	logmsg.Code = code

	data, err := json.Marshal(logmsg)
	remsg.Data = string(data)

	err = public.SendMsg(conn, remsg)
	if err != nil {
		fmt.Println("消息回复失败")
	}
	return
}

func RegPrec(conn net.Conn, msg string) (err error) {
	var userRegMsg public.RegMsg
	var umsg public.Messages
	var userRRmsg public.LoginReMsg

	err = json.Unmarshal([]byte(msg), &userRegMsg)
	if err != nil {
		return err
	}

	umsg.Type = "RegMsgResType"
	code := public.RedisPush(userRegMsg.UserID, userRegMsg.UserPwd)
	if code != 200 && code == 600 {
		userRRmsg.Code = code
		userRRmsg.Error = "UID 已注册"
	} else {
		userRRmsg.Code = code

	}

	data, err := json.Marshal(userRRmsg)
	umsg.Data = string(data)

	err = public.SendMsg(conn, umsg)
	if err != nil {
		fmt.Println("消息回复失败,err", err)
	}

	return
}
