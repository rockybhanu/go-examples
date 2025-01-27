package main

import "fmt"

func main() {
	println("Starting Loops Example")
	classicLoop()
	rangeKeyLoop()
}

func classicLoop() {
	for x := 0; x < 2; x++ {
		fmt.Println(x)
	}
}

func rangeKeyLoop() {
	// create array
	var fruits = []string{"apple", "banana", "Grapes"}
	for index, value := range fruits {
		fmt.Printf("Array Index is %d and Value is %s \n", index, value)
	}
}
