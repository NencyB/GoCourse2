package main

import "fmt"

func main() {
	avalue := 78
	fmt.Println(&avalue)

	p1 := person{
		name: "Nency",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

}

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Miss Money"
	(*p).name = "Batada"
}
