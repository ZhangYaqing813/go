package arrv

import (
	"fmt"
	"math/rand"
	"time"
)

func RandNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
	
}

func Arrex1(arr *[10]int )  {
	for index,value := range arr{
		fmt.Printf("index：%d, value:%d " ,index, value)
	}
	fmt.Println()
	lenth := len(arr)-1
	sum := 0
	for k := lenth; k>=0; k-- {
		fmt.Printf("%v ", arr[k])
		sum += arr[k]

	}
	fmt.Println()
	fmt.Printf("平均值为：%v\n", sum/(lenth+1) )

	var max int = arr[0]
	var maxindex int = 0
	var min int = arr[0]
	var minindex int = 0

	for j :=1; j<lenth;j++ {
		if max < arr[j+1] {
			max = arr[j+1]
			maxindex = j + 1
		}

		if min > arr[j] {
			min = arr[j]
			minindex = j
		}
	}
	for j :=0; j<lenth;j++{
		if arr[0] == 55 || arr[j] == 55{
			fmt.Println("数组中有 55")
		}
	}
	fmt.Println()
	fmt.Printf("最大值为：%d, 下标为：%d", max, maxindex)
	fmt.Println()
	fmt.Printf("最小值为：%d, 下标为：%d", min, minindex)
}

func InsertNumber(arr *[5]int, number int  ) {
	const  Lenth = len(arr)+1
	var newarr [Lenth] int
	for k :=0; k<len(arr); k++{
		newarr[k] = arr[k]
	}
	if newarr[Lenth-2] < number{
		newarr[Lenth-1] = number
	} else {
		newarr[Lenth-1] = number
	}

	for i:=0;i<Lenth-1; i++{
		for j :=Lenth-1; j>i; j--{
			if newarr[j] < newarr[j-1]{
				newarr[j], newarr[j-1] = newarr[j-1],newarr[j]
			}
		}
	}
	fmt.Println(newarr)
}

func Towarr(){
	var arr[3][4]int
	for i :=0; i< len(arr); i++{
		for j := 0; j< len(arr[i]);j++{
			fmt.Printf("请输入第%d行,第%d列的数字：" ,i, j)
			fmt.Scanln(&arr[i][j])
		}
	}

	fmt.Println(arr)
	for k,v := range arr{
		for k1,_ :=range v{
			if k == 0|| k1 == 0{
				arr[k][k1] = 0
			}
			if k == len(arr)-1 || k1 == len(arr[k]) - 1{
				arr[k][k1] = 0
			}
		}
	}
	fmt.Println()
	fmt.Println(arr)
}
