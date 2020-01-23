package tools

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Years(year int, mon int) {

	month := strconv.Itoa(mon)
	monthList := map[string]int {
		"1" : 31,
		"2" : 28,
		"3" : 31,
		"4" : 30,
		"5" : 31,
		"6" : 30,
		"7" : 31,
		"8" : 30,
		"9" : 30,
		"10": 31,
		"11": 30,
		"12": 31,
	}
	if year %4 ==0 && year %100 !=0 || year %400 == 0{
		fmt.Println("")
		if month == "2"{
			fmt.Println("当前月有29天")
		}
	}else if value, ok := monthList[month]; ok{
		fmt.Printf("%s月份天数为：%d", month,value)
	}
}

func rondNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(100)
}

func Guset(){
	var sourceNum int = rondNumber()
	var guestNum int
	fmt.Println(sourceNum)
	for i := 1; i <=10; i++ {

		fmt.Println("请输入你猜的数字")
		fmt.Scanln(&guestNum)
		if guestNum == sourceNum{

			if i == 1 {
				fmt.Println("你是一个天才")
				break
			}
			if i <=3 && i >= 2{
				fmt.Println("跟我一样聪明")
				break
			}
			if i <=9 && i >= 4{
				fmt.Println("一般般了")
				break
			}
			if i > 10{
				fmt.Println("不知道说什么了。。。。")
			}
		}else {
			fmt.Printf("继续猜。。。。，还有%d", 10-i)
			continue
		}

	}

}

func Prnumber ()  {
	sum := 0
	var count int = 0
	for i := 2; i <=100; i++{
		var is_pri bool = true
		for j := 2; j <=int(i/2) +1; j++{
			if i % j == 0{
				is_pri =false
				break
			}
		}
		if is_pri == true {
			sum += i
			count++
			fmt.Printf("%d ", i)
			if count%5 == 0 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
	fmt.Println("和为：", sum)
}

func Fish(newY int, newM int, newD int ){
	// 定义了一个月份有多少天的map,
	monthList := map[string]int {
		"1" : 31,
		"2" : 28,
		"3" : 31,
		"4" : 30,
		"5" : 31,
		"6" : 30,
		"7" : 31,
		"8" : 30,
		"9" : 30,
		"10": 31,
		"11": 30,
		"12": 31,
	}
	var years int = 1990
	var month int = 1
	var day int =0  // %5 的余数，用于判断是打鱼还是晒网
	var days int = 0 // 总天数
	for i := years; i < newY; i++{
		// 如果输入的年份和定义好的年份相同的话，则不需要处理，直接退出循环
		if i == newY{
			break
		}
		// 判断某一年是否位闰年，如果是的话则需要＋366 天
		if i %4 ==0 && i %100 !=0 || i %400 == 0 {
			days += 366
		}else {
			days += 365
		}
	}
	//如果输入的月份是1月，不需要＋整个月份的天数，只需要计算给出的天数即可
	for ;month < newM; month++{
		if newM == 1{
			break
		}
		days += monthList[strconv.Itoa(month)]
	}
	days += newD
	fmt.Println("newD= ",newD)
	day = days % 5
	fmt.Println("day = ", day)
	if day == 0 || day >= 4{
		fmt.Println("晒网")
	}else {
		fmt.Println("打鱼")
	}




}