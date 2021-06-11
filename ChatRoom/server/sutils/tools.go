package sutils

import (
	"fmt"
	"net"
)

func SeProcess(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 1024)

	n, err := conn.Read(b)
	if n != 4 || err != nil {
		fmt.Println("conn.Read 数据读取错误 err= ", err)
	} else {
		fmt.Println("读取数据为：", b[:n])
	}
}
