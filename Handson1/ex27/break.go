package main

import (
	"fmt"
)

func main() {
	i := 10
	for {

		if i <= 0 {
			break
		}
		fmt.Println(i)
		i--
	}
}
