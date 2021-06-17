package public

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func Response(conn net.Conn, msg Messages) (err error) {

	data, err := json.Marshal(msg)

	msgLen := uint32(len(data))
	var b = make([]byte, 4)
	binary.BigEndian.PutUint32(b[0:4], msgLen)
	n, err := conn.Write(b)
	fmt.Printf("=======n : %d, msgLen = %d ", n, int(msgLen))
	if n != 4 || err != nil {
		fmt.Println("数据长度发送失败")
	}
	fmt.Println("data, err := json.Marshal(msg) data :", data)
	n, err = conn.Write(data)
	if n != int(msgLen) || err != nil {
		fmt.Println("数据信息发送失败")
	}
	return
}
