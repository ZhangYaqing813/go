package untils

import (
	"fmt"
)

func Maopao( arrv *[10]int ) {

	var lenth = len(arrv) - 1
	for i := 0; i <lenth ; i++{
		for j :=lenth; j>i; j--{
			if arrv[j] < arrv[j-1]{
				//tmp := arrv[j]
				//arrv[j] = arrv[j-1]
				//arrv[j-1] = tmp
				arrv[j], arrv[j-1] = arrv[j-1], arrv[j]
			}
		}
	}
	fmt.Println(*arrv)
}

func Binary( arrv2[10]int, num int)  {
	var left int = 0
	var right int = len(arrv2) - 1
	for left <= right{
		var mid = left +(right - left) / 2
		fmt.Println("middle index is ", mid)
		fmt.Println("minddle value is ",arrv2[mid])

		if arrv2[mid] == num {
			fmt.Println("find the num ", arrv2[mid])
			break
		} else if arrv2[mid] > num{
			right = mid -1

			fmt.Println("right index is ", right)
			fmt.Println("right value is ", arrv2[right])
		}else if arrv2[mid] < num {
			left = mid + 1

		}

	}

}