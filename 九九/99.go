package main

import "fmt"

func main() {

	sum := 0
	max := 0
	min := 0
	for i :=1; i < 10;i++{

		for j:=1; j<=i;j++{
			fmt.Printf("%d x %d = %d\t",j, i, j*i)
		}
		fmt.Println()
	}


	for x := 1; x<=100; x ++{
		sum +=x
	}
	fmt.Println(sum)

	for x := 1; x <=100; x++{
		if x % 9 == 0{
			max ++
			min += x
		}
	}
	fmt.Printf("倍数和： %d\n 倍数个数： %d\n", min, max)
}
