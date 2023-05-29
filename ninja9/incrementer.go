package main

import (
	"fmt"
	"runtime"
	"sync"
)

var val int

func main() {
	// incre()
	norace()
}

func incre() {
	var wg sync.WaitGroup
	wg.Add(100)
	val = 0
	for i := 0; i < 100; i++ {
		go func() {
			v := val
			runtime.Gosched()
			v++
			val = v
			fmt.Println(val)
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End Value", val)
}

func norace() {
	var wg sync.WaitGroup
	var m sync.Mutex

	wg.Add(100)
	val = 0
	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			v := val
			v++
			val = v
			fmt.Println(val)
			fmt.Println(v)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("End Value", val)
}
