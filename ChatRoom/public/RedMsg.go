package public

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [1024]byte
}

func (T *Transfer) RedMsg() (msg Messages, err error) {
	//b := make([]byte, 1024)
	// 第一次读取客户端 发送的 将要发送数据的长度（默认为4）
	_, err = T.Conn.Read(T.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read 数据读取错误 err= ", err)
	}
	//反序列化
	var msgLen uint32
	msgLen = binary.BigEndian.Uint32(T.Buf[:4])
	fmt.Println("msgLen = ", msgLen)
	if err != nil {
		fmt.Println("binary.BigEndian.Uint32 err= ", err)
		return
	}
	fmt.Println("debug 11111111111")
	// 读取客户端发送的数据（真实数据）
	//buf := make([]byte, 1024)
	_, err = T.Conn.Read(T.Buf[:msgLen])
	if err != nil {
		fmt.Println("conn.Read err = ", err)
	}
	fmt.Println("debug 22222222222")
	fmt.Println("buff = ", T.Buf[:msgLen])
	err = json.Unmarshal(T.Buf[:msgLen], &msg)
	fmt.Println("debug msg= ", msg)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
	}

	fmt.Println("debug 3333333")
	return msg, err

}

func (T *Transfer) SendMsg(msg Messages) (err error) {
	//var b = make([]byte, 4)

	data, err := json.Marshal(msg)
	msgLen := uint32(len(data))
	binary.BigEndian.PutUint32(T.Buf[0:4], msgLen)

	n, err := T.Conn.Write(T.Buf[:4])

	fmt.Printf("=======n : %d, msgLen = %d ", n, int(msgLen))
	if n != 4 || err != nil {
		fmt.Println("数据长度发送失败")
	}
	fmt.Println("data, err := json.Marshal(msg) data :", data)

	n, err = T.Conn.Write(data)
	if n != int(msgLen) || err != nil {
		fmt.Println("数据信息发送失败")
	}
	return
}
