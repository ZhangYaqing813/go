package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Printf("连接失败 err = %v\n", err)
		return
	}

	fmt.Println("请输入信息:")
	//从标准输入读取数据
	reader := bufio.NewReader(os.Stdin)
	//终止字符
	str, err := reader.ReadString('\n')
	//发送数据信息
	conn.Write([]byte(str))
	//关闭链接
	defer conn.Close()
}
