package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{0.5, 2.0}
	sq := square{0.5}

	fmt.Println("Area of the triangle is ->")
	printArea(t)
	fmt.Println("Area of the square is ->")
	printArea(sq)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
