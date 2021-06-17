package cutils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go/ChatRoom/public"
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

	var msg public.Messages
	msg.Type = public.LoginMsgType
	var userloginmsg public.LoginMsg
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
	fmt.Printf("data = %v", data)

	msgLen := uint32(len(data))
	var b = make([]byte, 4)
	binary.BigEndian.PutUint32(b[0:4], msgLen)
	n, err := conn.Write(b)
	if n != 4 || err != nil {
		fmt.Println("发送失败")
		return
	}

	n, err = conn.Write(data)

	if err != nil {
		fmt.Println("发送失败")
		return
	}
	//time.Sleep(10*time.Second )

	resmsg, err := public.RedMsg(conn)
	if err != nil {
		fmt.Println("接受返回状态码错误 err = ", err)
		return
	}
	fmt.Println("********resmg = ", resmsg)
	var logcod public.LoginReMsg
	err = json.Unmarshal([]byte(resmsg.Data), &logcod)
	if resmsg.Type == public.LoginReMsgType {
		if err != nil {
			fmt.Println("解析返回状态码失败 err = ", err)
		}
		if logcod.Code == 200 {
			fmt.Println("服务验证登陆成功")
		}
	}
	//time.Sleep(10*time.Second )
	return
}
