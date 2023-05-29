package main

import "fmt"

const (
	a = 2017 + iota
	b = 2017 + iota
)

func main() {
	fmt.Println("\n", a, "\n", b)
}
