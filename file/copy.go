package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hhh")
	srcFile := "d:\\code\\go\\src\\go\\file\\123.jpg"
	decFile := "d:\\code\\go\\src\\go\\file\\456.jpg"
	src_f, err := os.Open(srcFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	src_f_w, err := os.OpenFile(decFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := io.Copy(src_f_w, src_f)
	fmt.Println(n)
	fmt.Println(err)

	//for {
	//	if err != io.EOF {
	//		n,err := image_f.Read(bs)
	//		if err != nil{
	//			fmt.Println(err)
	//			break
	//		}else {
	//			image_f_total = image_f_total + n
	//		}
	//
	//		n2,err := image_f2.Write(bs[:])
	//
	//		if err != nil{
	//			fmt.Println(err)
	//			break
	//		}else {
	//			image_f2_total = image_f2_total + n2
	//		}
	//	}

}
