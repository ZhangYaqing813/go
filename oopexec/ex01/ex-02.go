package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) Info() {
	fmt.Printf("姓名：%s\n年龄：%d\n", h.name, h.age)
}

func (s *Student) Info() {
	fmt.Printf("姓名：%s\n年龄：%d\n学校：%s", s.name, s.age, s.school)
}
func main() {
	student01 := Student{
		Human{
			"zhangsan",
			18,
		},
		"清华大学",
	}
	student01.Info()

}
