package main

import (
	"fmt"
)

func main() {
	//var test [26] byte
	//
	//for i := 0; i <26; i++{
	//	test[i] = 'A' + byte(i)
	//}
	//
	//for _,value :=  range test{
	//	fmt.Printf("%c ", value)
	//}
	//
	//var dg[6]int = [...]int {10,32,21,1,50,23}
	//var maxVal int = 0
	//var maxIdex int = 0
	//
	//for i :=0; i < 6; i++{
	//	if maxVal < dg[i]{
	//		maxVal = dg[i]
	//		maxIdex = i
	//	}
	//}
	//fmt.Println()
	//fmt.Printf("最大数：%d, 索引是：%d", maxVal, maxIdex)

	/*
		var arr [10] int
		for i := 0; i<10; i++{
			time.Sleep(10)
			(&arr)[i] = arrv.RandNumber()
		}
		arrv.Arrex1(&arr)
		fmt.Println("\n=======")

		var arr2[5] int = [5]int {1,3,5,7,9}
		arrv.InsertNumber(&arr2,4)
	*/
	//arrv.Towarr()

	q := [4]int{1, 2, 3, 4}
	fmt.Println("q=", q)

	for i := 0; i <= len(q)-1; i++ {
		fmt.Println(q[i])
	}

	for index := 0; index < len(q); index++ {
		value := q[index]
		fmt.Printf("index= %d vaule= %d\n", index, value)
	}

	var qq [4][4]int
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			qq[x][y] = x*y + 1
		}
		fmt.Println("qq[x]= ", qq[x])
	}
	fmt.Println("\n qq=", qq)

}
