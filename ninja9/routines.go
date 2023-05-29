package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin gs", runtime.NumGoroutine())
	fmt.Println("begin cpu", runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello There 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello There 2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("exit")
}
