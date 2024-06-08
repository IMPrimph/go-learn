package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	// telling wg to close the process
}

func main() {
	// telling wait group that we are adding 1 go routine
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	// telling wg to wait until all the go routines are done
	wg.Wait()
}
