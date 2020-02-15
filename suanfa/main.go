package main

import (
	"fmt"
	"go/suanfa/untils"
)

func main() {
	var arrv[10] int = [...] int {10, 9, 12, 6, 8, 1, 90, 18, 20, 28}
	untils.Maopao(&arrv)
	fmt.Println()
	fmt.Println(arrv)
	fmt.Println()
	untils.Binary(arrv,20)

}