package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go/ChatRoom/utils"
	"net"
)

func Login(userid int, userpwd string) (err error) {
	/*
		1.链接server
	*/
	conn, err := net.Dial("tcp", "0.0.0.0:10086")
	if err != nil {
		fmt.Println("连接服务器失败 err=", err)
		return
	}

	var msg utils.Messages
	msg.Type = utils.LoginMsgType
	var userloginmsg utils.LoginMsg
	userloginmsg.UserID = userid
	userloginmsg.UserPwd = userpwd

	data, err := json.Marshal(userloginmsg)
	if err != nil {
		fmt.Println("json 序列化失败 err= ", err)
		return
	}

	msg.Data = string(data)

	// 对msg 进行序列化

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json 序列化失败 err= ", err)
		return
	}

	msgLen := uint32(len(data))
	var b = make([]byte, 4)

	binary.BigEndian.PutUint32(b[0:4], msgLen)

	n, err := conn.Write(b)

	if n != 4 || err != nil {
		fmt.Println("发送失败")
	}

	return
}
