package untils

import "fmt"

func Menus() {
	// to print a menu ,just like :
	// 1 、 describe
	// 2 、 describe

	fmt.Println("------------家庭账簿------------\n")
	fmt.Println("\t\t1、记录明细")
	fmt.Println("\t\t2、收入明细")
	fmt.Println("\t\t3、支出明细")
	fmt.Println("\t\t4、余额详情")
	fmt.Println("\t\t0、退出系统\n")
	fmt.Println("-------------请选择-------------\n")
	fmt.Println()

}

func Icoming(balance *float64, income *float64, note *string, details *string) {
	// 收入记录
	*balance += *income
	*details += fmt.Sprintf("%v\t\t%v\t\t%v\t\t%v\n", "+", *balance, *income, *note)

}

func Favor(balance *float64, back *float64, note *string, details *string) {
	// 支出记录
	*balance -= *back
	*details += fmt.Sprintf("%v\t\t%v\t\t%v\t\t%v\n", "-", *balance, *back, *note)
}

func ShowDa(details *string) {
	fmt.Sprintf("%v\n", *details)
}
