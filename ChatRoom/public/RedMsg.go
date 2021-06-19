package public

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func RedMsg(conn net.Conn) (msg Messages, err error) {
	b := make([]byte, 1024)
	// 第一次读取客户端 发送的 将要发送数据的长度（默认为4）
	_, err = conn.Read(b)
	if err != nil {
		fmt.Println("conn.Read 数据读取错误 err= ", err)
	}
	//反序列化
	var msgLen uint32
	msgLen = binary.BigEndian.Uint32(b[:4])
	fmt.Println("msgLen = ", msgLen)
	if err != nil {
		fmt.Println("binary.BigEndian.Uint32 err= ", err)
		return
	}
	fmt.Println("debug 11111111111")
	// 读取客户端发送的数据（真实数据）
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ", err)
	}
	fmt.Println("debug 22222222222")

	err = json.Unmarshal(buf[:msgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
	}
	fmt.Println("debug 3333333")
	return msg, err

}
