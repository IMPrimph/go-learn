package main

import "fmt"

func say() {
	fmt.Println("start")
	defer fmt.Println("deferred")
	defer fmt.Println("deferred 2")
	fmt.Println("end")
}

func main() {
	say()
}
