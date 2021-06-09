package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//s := "Yes我爱慕课网!"
	//fmt.Println(len(s))
	//
	//for _, b := range []byte(s) {
	//	fmt.Printf("%x ", b)
	//}
	//fmt.Println()
	//
	//for i, ch := range s {
	//	fmt.Printf("(%d，%x)", i, ch)
	//}
	//fmt.Println()
	//
	//bytes := []byte(s)
	//
	//for len(bytes) > 0 {
	//	ch, size := utf8.DecodeRune(bytes)
	//	bytes = bytes[size:]
	//	fmt.Printf("%c", ch)
	//}
	//fmt.Println()
	//
	//for i, ch := range []rune(s) {
	//	fmt.Printf("(%d,%c)", i, ch)
	//}
	//fmt.Println()
	//stc := "hello "
	//
	//fmt.Println(strings.ToUpper(stc))

	DB, _ := sql.Open("mysql", "worm8pJs_u:9s%d7&F$1bA5Nk@tcp(46.51.233.134:3789)/dididu_exchange")

	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}
