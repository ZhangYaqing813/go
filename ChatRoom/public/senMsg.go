package public

import (
	"encoding/json"
	"fmt"
	"net"
)

func Send(conn net.Conn, msg LoginReMsg) (err error) {
	defer conn.Close()
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("sending msg json.Marshal err = ", err)
	}
	_, err = conn.Write(data)

	return
}
