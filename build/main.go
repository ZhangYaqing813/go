package main

import (
	"fmt"
	"time"
)

var Re_add = map[string]string{
	"cncc-admin": "ssh://git@git.cmes.io:4422/cncc-eshop/cncc-admin.git",
}

func main() {
	fmt.Println(time.Now().Year(), time.Now().Day())

}
