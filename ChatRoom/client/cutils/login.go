package cutils

import (
	"encoding/json"
	"fmt"
	"go/ChatRoom/public"
	"net"
	"os"
)

type UserProcess struct {
}

//用来处理用户登录的函数，需要传入用户ID，用户密码
//同时返回一个int 类型的状态码和一个error
func (U *UserProcess) Login(userid int, userpwd string) (err error) {

	// 创建链接
	conn, err := net.Dial("tcp", "0.0.0.0:10086")
	if err != nil {
		fmt.Println("连接服务器失败 err=", err)
		return
	}
	//创建一个结构体，用于调收发信息的处理方法
	tf := &public.Transfer{
		Conn: conn,
	}
	// 定义msg 的类型为public.Messages ,用于封装发送的消息
	var msg public.Messages
	// 定义消息类型为 LoginMsgType
	msg.Type = public.LoginMsgType
	//定义userloginmsg 封装用户登录信息
	var userloginmsg public.LoginMsg
	//用户ID
	userloginmsg.UserID = userid
	// 用户密码
	userloginmsg.UserPwd = userpwd
	// 序列化用户登录信息
	data, err := json.Marshal(userloginmsg)
	if err != nil {
		fmt.Println("json 序列化失败 err= ", err)
		return
	}
	//封装用户信息准备发送
	msg.Data = string(data)
	// 调用方法，发送登录信息
	err = tf.SendMsg(msg)
	if err != nil {
		fmt.Println("登录信息发送失败 ，err= ", err)
	}
	// 接受服务端返回的信息状态码
	resmsg, err := tf.RedMsg()
	if err != nil {
		fmt.Println("\n接受返回状态码错误 err = ", err)
		return
	}
	//定义 logcod 的类型为 public.LoginReMsg
	var logcod public.LoginReMsg
	//反序列化，获取返回状态码
	err = json.Unmarshal([]byte(resmsg.Data), &logcod)
	//判断返回的数据类型是否为登录状态返回的消息类型
	menus := &Menus{}
	if resmsg.Type == "LoginReMsgType" {
		if logcod.Code == 200 {
			go keepSession(conn)

			menus.TwoLevelMenu()
		} else {
			os.Exit(0)
		}
	}
	return
}

//用于处理新用户注册的方法，需要传入用户ID，用户密码
//同时返回一个int 类型的状态码和一个error

func (U *UserProcess) Register(userID int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println("注册服务打开链接失败 err = ", err)
	}
	tf := public.Transfer{
		Conn: conn,
	}
	var msg public.Messages      // 定义一个消息类型为 public.Messages 用于封装注册信息
	var userRegMsg public.RegMsg //定义一个消息类型为 public.RegMsg 用于封装用户注册时的消息
	var userRegresMsg public.LoginReMsg
	msg.Type = public.RegMsgType
	userRegMsg.UserID = userID
	userRegMsg.UserPwd = userPwd
	data, err := json.Marshal(userRegMsg)
	msg.Data = string(data)

	err = tf.SendMsg(msg)
	if err != nil {
		fmt.Println("发送注册信息失败")
		return
	}

	msg, err = tf.RedMsg()
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

func keepSession(conn net.Conn) {
	fmt.Println("keepsession 等待读取信息")
	tf := &public.Transfer{
		Conn: conn,
	}
	msg, err := tf.RedMsg()
	if err != nil {
		fmt.Println("KeepSession ResMsg err = ", err)
	}
	fmt.Println("keepsession read msg = ", msg)
}
