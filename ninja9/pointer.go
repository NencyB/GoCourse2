package main

import (
	"fmt"
)

type person struct {
	name string
}

type human interface {
	speak()
}

func saysomething(h human) {
	h.speak()
}

func (p *person) speak() {
	fmt.Println("Hello Person here")
}

func main() {
	p1 := person{
		name: "Nency",
	}
	p1.speak()

	saysomething(&p1)
	// saysomething(p1)

}
