package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	DB, _ := sql.Open("mysql", "owenE0Bx7D:Ks^d7&L$8sLN0k@tcp(46.51.233.134:3789)/dididu_exchange")

	if err := DB.Ping(); err != nil {
		fmt.Println(err)
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}
