package sutils

import (
	"encoding/json"
	"fmt"
	"go/ChatRoom/public"
	"net"
)

type Session struct {
	Conn net.Conn
}

type userPrc struct {
	conn net.Conn
	msg  string
}

func (S *Session) SeProcess() {
	/*
		1、会话处理函数
		1.1 根据不同的message 中的message.Tpye 类型进行不同的业务处理
	*/
	// 延时关闭会话

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭链接失败")
		}
	}(S.Conn)
	// 初始化一个 Transfer 实例
	tf := &public.Transfer{
		Conn: S.Conn,
	}

	fmt.Println("开始接收数据")
	// 调用ReMsg 函数获取client 发送的信息

	var msg, err = tf.RedMsg()
	if err != nil {
		fmt.Println("数据返回错误")
	}

	// 初始化一个 UserPrc 实例
	up := userPrc{
		conn: S.Conn,
		msg:  msg.Data,
	}
	// 根据 MSG.Type 进行不同的业务处理
	switch msg.Type {
	// 处理用户登录的信息
	case public.LoginMsgType:
		fmt.Println("用户登陆校验")
		err = up.loginPrec()
		if err != nil {
			fmt.Println("回复信息失败")
		}
	// 处理用户注册的业务请求
	case public.RegMsgType:
		fmt.Println("新用户注册=========")
		err = up.regPrec()

	default:
		fmt.Println("暂时无法处理")
	}
}

func (U *userPrc) loginPrec() (err error) {

	// userMsg 接受client 发送过来的用户信息
	var userMsg public.LoginMsg

	// logmsg 用于接受用户登录结果的信息
	var logmsg public.LoginReMsg

	// remsg 回复client 信息
	var remsg public.Messages
	//解析json ,付给userMsg,获取userID和userPwd

	tf := &public.Transfer{
		Conn: U.conn,
	}

	json.Unmarshal([]byte(U.msg), &userMsg)

	//设置 message 类型为client 登录类型的回复消息
	remsg.Type = "LoginReMsgType"
	//调用ReisGet 函数，判断client 发送的UserID与userPwd 匹配,拿到返回值code，组装进回复的消息体内
	logmsg.Code = public.RedisGet(userMsg.UserID, userMsg.UserPwd)
	//序列话消息体
	data, err := json.Marshal(logmsg)
	if err != nil {
		fmt.Println("回复client 的loninResMsg 消息体序列化错误 err = ", err)
	}
	// loninResMsg 组装到回复的消息体
	remsg.Data = string(data)
	//调用 senMsg 函数进行信息回复
	err = tf.SendMsg(remsg)
	if err != nil {
		fmt.Println("消息回复失败")
	}
	return
}

// 用户注册消息的处理函数

func (U *userPrc) regPrec() (err error) {
	// client 发送的注册信息
	var userRegMsg public.RegMsg
	// 回复client 的注册结果的总信息
	var umsg public.Messages
	// umsg 的有效载荷（存放注册结果的信息），
	var userRRmsg public.LoginReMsg
	//设置umsg 的消息类型为回复注册消息的结果类型
	umsg.Type = "RegMsgResType"
	tf := public.Transfer{
		Conn: U.conn,
	}

	// 解析client 发送的msg 信息，
	err = json.Unmarshal([]byte(U.msg), &userRegMsg)
	if err != nil {
		return err
	}

	// 用户信息存入redis 中
	code := public.RedisPush(userRegMsg.UserID, userRegMsg.UserPwd)
	//判断是否成功写入redis
	// 200 表示成功
	//600 表示此UID已经被注册
	if code != 200 && code == 600 {
		userRRmsg.Code = code
		userRRmsg.Error = "UID 已注册"
	} else {
		userRRmsg.Code = code
	}

	data, err := json.Marshal(userRRmsg)
	if err != nil {
		fmt.Println("注册回复信息序列化失败err = ", err)
	}

	// 组装到umsg 中
	umsg.Data = string(data)
	// 发送注册结果到client
	err = tf.SendMsg(umsg)
	if err != nil {
		fmt.Println("消息回复失败,err", err)
	}
	return
}
