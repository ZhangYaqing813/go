package main

import (
	"fmt"
	"github.com/ZhangYaqing813/go/build/utils"
)

var Re_add = map[string]string{
	"cncc-admin": "ssh://git@git.cmes.io:4422/cncc-eshop/cncc-admin.git",
}

func main() {
	cmesadmin := utils.Fat("cmes-admin", Re_add["cncc-admin"], "target", "1000000")
	fmt.Println(Re_add["cncc-admin"])
	fmt.Println(cmesadmin.Repository)

}
