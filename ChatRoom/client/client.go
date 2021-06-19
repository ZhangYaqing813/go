package main

import (
	"fmt"
	"go/ChatRoom/client/cutils"
)

var userID int
var userPwd string

func main() {
	//循环控制条件
	var loop = true
	// 选项就收
	var key int
	//菜单的循环
	for loop {
		fmt.Println("\t=======欢迎进入聊天室=======")
		fmt.Println("\t\t\t1、新用户注册")
		fmt.Println("\t\t\t2、用户登录")
		fmt.Println("\t\t\t3、退出系统")
		fmt.Println("\t--------------------------")

		fmt.Println("请输入选择：")

		fmt.Scanf("%d\n", &key)
		//根据不同的选择进入不同的子菜单
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
	// 子菜单
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
