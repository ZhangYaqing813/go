package main

import (
	"fmt"
	"go/oopexec/untils"
)

func main() {
	acc := tools.Newacc("111111111", "999999", 50)
	fmt.Println((*acc))

	acc.Getmeny("999999", 20)

	fmt.Println((*acc))
}
