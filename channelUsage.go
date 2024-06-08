package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, val int) {
	defer wg.Done()
	c <- val * 5
}

func main() {
	fooChan := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go foo(fooChan, i)
	}
	wg.Wait()
	close(fooChan)

	for item := range fooChan {
		fmt.Println(item)
	}
}
