package main

import "log"

func main() {
	grades := make(map[string]float32)

	grades["user1"] = 32
	grades["user2"] = 90

	log.Println(grades)
	log.Println(grades["user1"])

	// to delete a key from a map
	delete(grades, "user2")
	log.Println(grades)
}
