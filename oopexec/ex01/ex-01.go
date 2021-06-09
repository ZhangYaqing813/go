package main

import "fmt"

type Point struct {
	x float64
	y float64
}

func (p Point) area(P Point) float64 {

	return p.x * p.y
}
func (p *Point) area01(P Point) float64 {

	return p.x * p.y
}

func main() {
	p1 := Point{3, 4}
	p2 := Point{5, 6}
	fmt.Println(p1.area(p1))
	fmt.Println(p2.area01(p2))
}
