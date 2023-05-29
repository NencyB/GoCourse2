package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(2)

		switch x {
		case 0:
			fmt.Println("This is case 0")
		case 1:
			fmt.Println("This is Case 1")
		}
	}
}
