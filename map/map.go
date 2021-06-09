package main

import "fmt"

func main() {

	//map 定义和初始化 方法一
	//studentsAge := make(map[string]int)
	//studentsAge["xiaoming"] = 15
	//studentsAge["xiaohong"] = 14
	//studentsAge["xiaohei"] = 15
	//studentsAge["dawa"] = 20
	//studentsAge["erwa"] = 1000
	//
	//for key, value := range studentsAge {
	//	fmt.Printf("kye = %s, value  = %d\n", key, value)
	//}

	//fmt.Println("==============")
	//
	////map 定义和初始化 方法2
	//studentsAge1 := map[string]int{}
	//fmt.Println("赋值前的map", studentsAge1)
	//studentsAge1["john"] = 1
	//
	//fmt.Printf("姓名：john, 年龄：%d\n", studentsAge1["john"])
	//fmt.Printf("studentsAge1 的内存地址是：%p\n", &studentsAge1)
	//
	//fmt.Println("==============")
	//
	////map 定义和初始化 方法3
	//
	//var studentsAge2 map[string]int
	//fmt.Printf("studentsAge2 的内存地址是：%p\n", &studentsAge2)
	////studentsAge2["erha"] = 1
	//studentsAge2 = make(map[string]int)
	//fmt.Printf("studentsAge2 的内存地址是：%p\n", &studentsAge2)
	//studentsAge2["erha"] = 1
	//studentsAge2["jinmao"] = 3
	//fmt.Printf("studentsAge2 的内存地址是：%p\n", &studentsAge2)
	//fmt.Printf("姓名：erha, 年龄：%d\n", studentsAge2["erha"])

	//studentsAge3 := map[string]string{}
	//studentsAge3["erha"] = "big"
	//
	//if _, ok := studentsAge3["erha"]; ok { //先判断是否存在
	//	fmt.Println("studentsAge hava it")
	//	if studentsAge3["erha"] == "big" { //比较值
	//		fmt.Println(" studentsAge3[\"erha\"] = big")
	//	} else {
	//		fmt.Println(" studentsAge3[\"erha\"] != big")
	//	}
	//
	//} else {
	//	fmt.Println(ok)
	//}
	//
	//studentsAge3["erha"] = "small" //修改一下
	//
	//if _, ok := studentsAge3["erha"]; ok { //先判断是否存在
	//	fmt.Println("studentsAge hava it")
	//	if studentsAge3["erha"] == "big" { //比较值
	//		fmt.Println(" studentsAge3[\"erha\"] = big")
	//	} else {
	//		fmt.Println(" studentsAge3[\"erha\"] = ", studentsAge3["erha"])
	//	}
	//} else {
	//	fmt.Println(ok)
	//}

	fmt.Println(fib(6))
}

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		return (fib(n-1) + fib(n-2))
	}
}
