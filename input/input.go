package main

import (

	"fmt"
	"zyq.cn/input/tools"
)

func main() {

	var num int = 0
	var ch int = 0

	fmt.Printf("source： num= %d\n ch= %d\n", num, ch)
	fmt.Scanf("%d %d\n", &num, &ch)
	fmt.Println()
	fmt.Printf("请输入要转换的十进制数字 num： %d\n 选项： %d\n", num, ch)

	switch ch {
	case 1 :
		fmt.Println(tools.ToBy(num))
	case 2 :
		fmt.Println(tools.ToOX(num))
	case 3:
		fmt.Println(tools.BaniyToDec(num))
	default:
		fmt.Println("请选则正确的选项")

	}

}
