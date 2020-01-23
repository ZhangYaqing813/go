package main

import "fmt"

func def (a int, b int ) int {
	res := 0
	defer fmt.Printf("第一次压入： %d\n", a)
	defer  fmt.Printf("第二次压入： %d\n", b)
	res = a + b
	return  res
}

func main() {
	var rest = def(10,20)
	fmt.Printf("第一次打印： %d\n", rest)
}
