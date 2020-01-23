package main

import "fmt"

func main() {
	var arrv[10] int = [...] int {10, 9, 12, 6, 8, 1, 90, 18, 20, 28}
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
	fmt.Println(arrv)
}
