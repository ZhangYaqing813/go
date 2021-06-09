package main

import "fmt"

type info struct {
	ID   int
	Role string
}

type Persion struct {
	Name    string
	Sex     string
	Address string
	info
}

func factory(name string, Sex string, address string) *Persion {
	return &Persion{
		Name:    name,
		Sex:     Sex,
		Address: address,
	}
}

type mathd struct {
	lenth int
	weide int
}

func (m mathd) area() int {
	return m.lenth * m.weide
}

func (m *mathd) crycle() int {
	return m.weide*2 + m.lenth*2
}

func main() {
	XiaoM := Persion{
		Name:    "XiaoMing",
		Sex:     "Man",
		Address: "BeiJing",
	}

	var XH Persion
	XH.Name = "xiaohong"
	XH.Sex = "Fm"
	XH.Address = "Hb"

	P := &XiaoM
	P.Sex = "Fm"

	var Dawa = factory("huluwa", "man", "dashan")
	Dawa.ID = 001
	Dawa.Role = "First"
	fmt.Println(*Dawa)
	fmt.Println(XiaoM)
	fmt.Println(XH)

	var x = mathd{2, 5}

	fmt.Println("x area = ", x.area())
	fmt.Println("x crycle = ", x.crycle())
}
