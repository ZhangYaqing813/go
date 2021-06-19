package main

import (
	"fmt"
	"go/ChatRoom/client/cutils"
)

var userID int
var userPwd string

func main() {
	var loop = true
	var key int

	for loop {
		fmt.Println("\t=======欢迎进入聊天室=======")
		fmt.Println("\t\t\t1、新用户注册")
		fmt.Println("\t\t\t2、用户登录")
		fmt.Println("\t\t\t3、退出系统")
		fmt.Println("\t------------------------")

		fmt.Println("请输入选择：")
		fmt.Scanf("%d\n", &key)
		//tmp := bufio.NewReader(os.Stdin)
		//key ,_ = tmp.ReadString('\n')

		switch key {
		case 1:
			fmt.Println("新用户注册")
			loop = false
		case 2:
			//fmt.Println("用户登录")
			loop = false
		case 3:
			fmt.Println("即将退出系统")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

	if key == 1 {
		fmt.Println("新用户注册")
		fmt.Println("请输入用户ID：")
		fmt.Scanf("%d\n", &userID)

		fmt.Println("请输入用户PWD：")
		fmt.Scanf("%s\n", &userPwd)
		err := cutils.Register(userID, userPwd)
		if err != nil {
			fmt.Println("用户注册失败err = ", err)
		}
	} else if key == 2 {
		fmt.Println("用户登录")
		fmt.Println("请输入用户ID：")
		fmt.Scanf("%d\n", &userID)

		fmt.Println("请输入用户PWD：")
		fmt.Scanf("%s\n", &userPwd)
		err := cutils.Login(userID, userPwd)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("即将退出系统")
	}

}
