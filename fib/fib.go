package main

import "fmt"

func fib(n int) int {
	if n <= 2 {
		return n
	} else {

		tmp := fib(n-2) + fib(n-1)
		return tmp
	}

}

func taozi(t int) int {
	if t == 10 {
		return 1
	} else {
		return (taozi(t+1) + 1) * 2
	}

}

func area(x int) func() int {
	var h int
	h = x
	return func() int {
		return h * h
	}
}

func main() {
	//	f := fib(6)
	//fmt.Println(f)
	f := area(10)
	fmt.Println(f())
	//
	//a,b := 0, 1
	//
	//for i := 0; i<=5; i++{
	//	fmt.Println(a)
	//	a, b = b, a+b
	//}

	//fmt.Println(taozi(1))
}
