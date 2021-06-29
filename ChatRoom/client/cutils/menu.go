package cutils

import "fmt"

type Menus struct {
	key int
}

func (M *Menus) TwoLevelMenu() {
	for {

		fmt.Println("1、显示在线用户列表")
		fmt.Println("2、和谁聊天")
		fmt.Println("3、退出本级菜单")
		fmt.Println("请输入选择（1-3）")

		fmt.Scanf("%d\n", &M.key)
		switch M.key {
		case 1:
			fmt.Println("显示在线用户")
		case 2:
			fmt.Println("和谁在聊天")
		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("暂时无法处理")
		}
	}
}
