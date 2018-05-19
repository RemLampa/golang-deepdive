package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	sq := square{
		sideLength: 4.0,
	}
	printArea(sq)

	tr := triangle{
		height: 4.0,
		base:   2.0,
	}
	printArea(tr)
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(sh shape) {
	fmt.Println("Area:", sh.getArea())
}
