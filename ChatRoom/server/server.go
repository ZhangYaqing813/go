package main

import (
	"fmt"
	"go/ChatRoom/utils"
	"net"
)

func main() {
	/*
		1、打开监听 net.listen,参数为 tcp协议，监听地址:端口  0.0.0.0：10086
		2、开始接受数据 listen.accept,返回两个对象，conn 和err
		3、开启协程处理conn, SeProcess()

	*/

	Listen, err := net.Listen("tcp", "0.0.0.0:10086")
	fmt.Println("打开监听")
	if err != nil {
		fmt.Println("监听错误 err= ", err)
		return
	}
	defer Listen.Close()
	//开始接受数据
	fmt.Println("准备接受数据")
	for {
		conn, err := Listen.Accept()
		if err != nil {
			fmt.Println("数据接收错误 err=", err)
		}
		fmt.Printf("client inf:%v", conn.RemoteAddr())
		go utils.SeProcess(conn)
	}

}
