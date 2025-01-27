package main

import "fmt"

func main() {
	println("Starting Loops Example")
	classicLoop()
}

func classicLoop() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}
