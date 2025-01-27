package main

import "fmt"

func main() {
	println("Starting Loops Example")
	classicLoop()
	rangeKeyLoop()
}

func classicLoop() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}
func rangeKeyLoop() {
	// create array
	var fruits = []string{"Apple", "Orange", "Banana", "Grape"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}
}
