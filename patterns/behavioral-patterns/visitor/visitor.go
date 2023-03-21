package main

import "fmt"

// Shape is Element
type Shape interface {
	accept(Visitor)
}

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

type Rectangle struct {
	l, r int
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}


// Concrete Visitor
type areaCalculator struct {
}

func (a *areaCalculator) visitForSquare(s *Square) {
	fmt.Printf("area is %d\n", s.side * s.side)
}

func (a *areaCalculator) visitForCircle(c *Circle) {
	fmt.Printf("area is %.2f\n", float64(c.radius) * float64(c.radius) * 3.14)
}

func (a *areaCalculator) visitForRectangle(r *Rectangle) {
	fmt.Printf("area is %d\n", r.l * r.r)
}

func main() {
	v := &areaCalculator{}
	sq := &Square{3}
	cir := &Circle{3}
	rect := &Rectangle{3, 4}
	sq.accept(v)
	cir.accept(v)
	rect.accept(v)
}

// output:
//area is 9
//area is 28.26
//area is 12