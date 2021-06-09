package main

import (
	"fmt"
	"os"
)

func main() {
	srcFile := "d:\\code\\go\\src\\go\\file\\123.jpg"
	decFlile := "d:\\code\\go\\src\\go\\file\\456.jpg"
	tmpFlile := "d:\\code\\go\\src\\go\\file\\tmp"

	bs := make([]byte, 1024, 124)

	src_f, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	dec_f, err := os.OpenFile(decFlile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	tmp_f, err := os.OpenFile(tmpFlile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

}
