package main

import (
	"fmt"
	"net"
)

func proc(conn net.Conn) {
	b := make([]byte, 1024)
	n, err := conn.Read(b)
	if err != nil {
		fmt.Println("conn read fail err = ", err)
	} else {
		fmt.Println(string(b[:n]))
	}
	defer conn.Close()

}

func main() {

	conn, err := net.Listen("tcp", "0.0.0.0:12345")
	defer conn.Close()
	if err != nil {
		return
	}

	for {
		fmt.Println("开始接受信息")
		conn, err := conn.Accept()
		if err != nil {
			fmt.Println("接受信息错误 err = %s\n", err)
		} else {
			fmt.Printf("accept client info = %v\n", conn.RemoteAddr().String())
			go proc(conn)
		}
	}
}
