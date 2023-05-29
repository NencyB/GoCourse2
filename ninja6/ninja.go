package main

import (
	"fmt"
	"math"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	defer print(foo(ii...))

	iii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	print(boo(iii))

	print(bar())

	persons()

	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}

	info(circ)
	info(squa)
}

func foo(ii ...int) int {
	var total int
	for _, i := range ii {
		total += i
	}
	return total
}

func boo(iii []int) int {
	var totaliii int
	for _, i := range iii {
		totaliii += i
	}
	return totaliii
}

func bar() (int, string) {
	return 78, "int and It's a String"
}

func print(value ...interface{}) {
	fmt.Println(value)
}

// _____________________________________________

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(
		"Hi, here it is", p.first, p.last, p.age,
	)
}

func persons() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p1.speak()
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * math.Pi * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
