package main

import (
	"fmt"
	"go/ChatRoom/client/cutils"
	"os"
)

var userID int
var userPwd string

func main() {
	//循环控制条件
	// 选项就收
	var key int
	//菜单的循环
	user := &cutils.UserProcess{}
	for true {
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
			fmt.Println("请输入用户ID：")
			fmt.Scanf("%d\n", &userID)

			fmt.Println("请输入用户PWD：")
			fmt.Scanf("%s\n", &userPwd)
			err := user.Register(userID, userPwd)
			if err != nil {
				fmt.Println("用户注册失败err = ", err)
			}
		case 2:

			fmt.Println("用户登录")
			fmt.Println("请输入用户ID：")
			fmt.Scanf("%d\n", &userID)

			fmt.Println("请输入用户PWD：")
			fmt.Scanf("%s\n", &userPwd)

			err := user.Login(userID, userPwd)
			if err != nil {
				fmt.Println("登录失败 err= ", err)
				os.Exit(0)
			}
		case 3:
			fmt.Println("即将退出系统")

		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
	// 子菜单

}
