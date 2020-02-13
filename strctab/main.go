package main

import "fmt"

type person struct {
	name string
	age int
	add string
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return (*r).width * (*r).height
}

func (r *rect) perim() int {
	return ((*r).height+(*r).width) * 2
}


func (s *person) print()  {
	fmt.Println("姓名：", (*s).name)
}

func main() {

	p1 := person{"Wangwu",15,"hebei"}
	var p person
	p.name = "zhangsan"
	p.age = 15
	p.add = "Beijing"

	fmt.Println(p)
	ptr := &p
	(*ptr).name = "LiSi"

	fmt.Printf("姓名：%v\n", (*ptr).name)

	p1.print()

	fmt.Println("分割========")
	//初始化 rect 结构体
	r := rect{10,20}
	fmt.Printf("rect 面积：%v\n", r.area())
	fmt.Printf("rect 周长：%v\n", r.perim())


	rptr := &r
	rptr.width = 30
	fmt.Printf("rect 面积：%v\n", r.area())
	fmt.Printf("rect 周长：%v\n", r.perim())

}
