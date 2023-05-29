package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("Outer loop :", i)
		for j := 0; j <= 5; j++ {
			fmt.Print("inner loop:", i)
		}
	}
}
