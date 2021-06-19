package main

import (
	"fmt"
	"go/ChatRoom/server/sutils"
	"net"
)

func main() {
	/*
		1、打开监听 net.listen,参数为 tcp协议，监听地址:端口  0.0.0.0：10086
		2、开始接受数据 listen.accept,返回两个对象，conn 和err
		3、开启协程处理conn, SeProcess()

	*/
	//监听0.0.0.0 的端口 10086
	Listen, err := net.Listen("tcp", "0.0.0.0:10086")
	if err != nil {
		fmt.Println("监听错误 err= ", err)
		return
	}
	defer func() {
		err = Listen.Close()
		if err != nil {
			fmt.Println("服务监听关闭失败")
		}
	}()
	//开始接受数据
	for {
		conn, err := Listen.Accept()
		if err != nil {
			fmt.Println("数据接收错误 err=", err)
		}
		fmt.Printf("client inf:%v", conn.RemoteAddr())
		go sutils.SeProcess(conn)
	}

}
