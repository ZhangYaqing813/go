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
	//json 转struct
	uloginmsg := public.LoginMsg{}
	fmt.Println(msg)
	json.Unmarshal([]byte(msg.Data), &uloginmsg)

	//登陆判断
	//var loginCode public.LoginReMsg

	if msg.Type == public.LoginMsgType {
		fmt.Println("准备校验用户名密码")
		if uloginmsg.UserID == 100 && uloginmsg.UserPwd == "abc" {
			fmt.Println("login secuss")
			//loginCode.Code = 200
			//err = public.Send(conn, loginCode)
			//if err !=nil{
			//	fmt.Println("sending code err = ",err)
			//	return
			//}
		}

	}

}
