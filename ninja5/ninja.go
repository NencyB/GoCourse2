package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	// loopingslice()
	// exthree()
	anony()
}

func loopingslice() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	fmt.Println("\n", p1.first, p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println("\n", p2.first, p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("______________________")
	}
}

// // _____________________________
// Create a new type: vehicle.
// ○ The underlying type is a struct.
// ○ The fields:
// ■ doors
// ■ color
// ● Create two new types: truck & sedan.
// ○ The underlying type of each of these new types is a struct.
// ○ Embed the “vehicle” type in both truck & sedan.
// ○ Give truck the field “fourWheel” which will be set to bool.
// ○ Give sedan the field “luxury” which will be set to bool. solution
// ● Using the vehicle, truck, and sedan structs:
// ○ using a composite literal, create a value of type truck and assign values to the
// fields;
// ○ using a composite literal, create a value of type sedan and assign values to the
// fields.
// ● Print out each of these values.
// ● Print out a single field from each of these values.
// ________________________________________________________//

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheeler bool
}

type sedan struct {
	vehicle
	luxury bool
}

func exthree() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Black",
		},
		fourwheeler: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Blue",
		},
		luxury: false,
	}

	fmt.Print(t, "\n", s)
}

func anony() {

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
