package main

import (

	"fmt"
)

func main() {
	var test [26] byte

	for i := 0; i <26; i++{
		test[i] = 'A' + byte(i)
	}

	for _,value :=  range test{
		fmt.Printf("%c ", value)
	}

	var dg[6]int = [...]int {10,32,21,1,50,23}
	var maxVal int = 0
	var maxIdex int = 0

	for i :=0; i < 6; i++{
		if maxVal < dg[i]{
			maxVal = dg[i]
			maxIdex = i
		}
	}
	fmt.Println()
	fmt.Printf("最大数：%d, 索引是：%d", maxVal, maxIdex)


}
