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
	for {
		untils.Menus()
		fmt.Scanln(&choose)
		if choose >= 0 && choose <= 4 {
			switch {
			case choose == 1:
				//fmt.Printf("%v\n", details)
				untils.ShowDa(&details)
			case choose == 2:
				untils.Icoming(&balance, &income, &note, &details)
			case choose == 3:
				untils.Favor(&balance, &income, &note, &details)
			case choose == 4:
				fmt.Println("显示余额")
			case choose == 0:
				for flag == true {
					fmt.Println("确定要退出本系统吗(Y/N, Y)?")
					key := ""
					fmt.Scanln(&key)
					if key == "Y" || key == "y" {
						flag = false
						break
					} else if key == "N" || key == "n" {

						break
					} else {
						fmt.Println("请按提示输入")
					}
				}

			default:
				fmt.Println("请输入正确的选项")
			}
		}
		// 退出时进行提示
		//fmt.Println(flag)
		if !flag {
			fmt.Println("退出本系统......")
			break
		}
	}
}
