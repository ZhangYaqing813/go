package main

import (
	"fmt"
	"go/family/untils"
)

func main() {
	var flag bool = true
	var choose int = 0
	var balance float64 = 1000.00
	var note string
	var income float64 = 0.00
	var details string = "收支\t余额\t收支金额\t\t明     细\n"
	fmt.Println("-------欢迎登录家庭账簿系统------\n")
	for flag {
		untils.Menus()
		fmt.Scanln(&choose)

		if choose >= 0 && choose <= 4 {

			switch {
			case choose == 1:
				fmt.Printf("%v\n", details)
				//untils.ShowDa(&details)
			case choose == 2:
				fmt.Println("请输入收入：")
				fmt.Scanln(&income)
				fmt.Println("请输入收入来源：")
				fmt.Scanln(&note)
				untils.Icoming(&balance, &income, &note, &details)
			case choose == 3:
				fmt.Println("请输入收入：")
				fmt.Scanln(&income)
				fmt.Println("请输入收入来源：")
				fmt.Scanln(&note)
				untils.Favor(&balance, &income, &note, &details)
			case choose == 4:
				fmt.Println("显示余额")
			case choose == 0:
				flag = false
			default:

			}
			//if flag {
			//	fmt.Println("即将退出本系统......")
			//	break
			//}
		} else {
			fmt.Println("请输入正确的选项")
		}
	}
}
