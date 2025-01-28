package main

import (
	"fmt"
)

func main() {
	// Entry point of the program
	println("Starting Loops Example")

	// Demonstrating a classic `for` loop
	classicLoop()

	// Demonstrating `range` loop for arrays
	rangeKeyLoopForArray()

	// Demonstrating `range` loop for maps
	rangeKeyLoopForMap()

	// Demonstrating a `for` loop with custom termination logic
	justForLoop()
}

// classicLoop demonstrates a basic `for` loop in Go
func classicLoop() {
	// Loop that runs from 0 to 1 (exclusive of 2)
	for x := 0; x < 2; x++ {
		fmt.Println(x)
	}
}

// rangeKeyLoopForArray demonstrates how to use `range` with an array
func rangeKeyLoopForArray() {
	// Create a slice (dynamic array) of fruits
	var fruits = []string{"apple", "banana", "Grapes"}

	// Loop through the array using `range`
	// `index` holds the position, and `value` holds the actual item
	for index, value := range fruits {
		fmt.Printf("Array Index is %d and Value is %s \n", index, value)
	}
}

// rangeKeyLoopForMap demonstrates how to use `range` with a map
func rangeKeyLoopForMap() {
	// Create a map where keys are names and values are ages
	var staff = make(map[string]int)
	staff["John"] = 25
	staff["Darren"] = 28
	staff["Cannon"] = 24
	staff["Shankar"] = 30

	// Print the entire map (raw structure)
	fmt.Println(staff)

	// Loop through the map using `range` to print each key-value pair
	for key, value := range staff {
		fmt.Printf("Staff name %s and age is %d \n", key, value)
	}

	// Loop through the map to print only the keys
	for name := range staff {
		fmt.Println(name)
	}
}

// justForLoop demonstrates how to use a `for` loop with custom logic
func justForLoop() {
	// Create a map of countries and their national animals
	animals := map[string]string{
		"India":          "Bengal Tiger",
		"United States":  "Bald Eagle",
		"Australia":      "Kangaroo",
		"Canada":         "Beaver",
		"China":          "Giant Panda",
		"United Kingdom": "Lion",
	}

	// Extract all keys from the map into a slice
	keys := make([]string, 0, len(animals))
	for key := range animals {
		keys = append(keys, key)
	}

	// Get the total number of entries in the map
	countOfAnimals := len(animals)

	// Example 1: Using a `for` loop with manual index increment and termination
	x := 0
	for {
		// Print the key and value using the current index
		fmt.Printf("National Animal of %s is %s\n", keys[x], animals[keys[x]])

		// Terminate the loop when the last element is reached
		if x == countOfAnimals-1 {
			break
		}
		x++
	}

	// Example 2: Using a `for` loop with a condition
	fmt.Println("Another implementation...")
	y := 0
	for y < countOfAnimals-1 {
		fmt.Printf("National Animal of %s is %s\n", keys[y], animals[keys[y]])
		y++
	}

	// Example 3: Using a `for` loop with initialization, condition, and increment
	fmt.Println("Another implementation...")
	for z := 0; z < countOfAnimals-1; z++ {
		fmt.Printf("National Animal of %s is %s\n", keys[z], animals[keys[z]])
	}
}
