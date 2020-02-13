package main

import (
	"fmt"
	"math"
)

type  getall interface {
	area() float64
	perim() float64
}

type rect struct {
	width, heigh float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64  {
	return r.width * r.heigh
}

func (r rect) perim() float64 {
	return (r.heigh + r.width) * 2
}

func (c circle) area() float64  {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim()  float64 {
	return 2 * c.radius * math.Pi
}

func show(s getall) {

	fmt.Printf("s是：%v\n", s)
	fmt.Printf("面积：%v\n", s.area())
	fmt.Printf("周长：%v\n", s.perim())
}

func main() {
	r := rect{10,10}
	c := circle{10}

	show(r)
	show(c)
}
