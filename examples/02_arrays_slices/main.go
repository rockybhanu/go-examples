package main

import "fmt"

func main() {

	/*
		Fixed Length
		Same Type
		Indexable
		Contiguous in Memory
	*/

	// declaration of array with shorthand
	fruits := [3]string{"apple", "orange", "banana"}
	fmt.Println(fruits)

	// declaration of array with var
	var intArr [3]int16
	// Contiguous in Memory address printed
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	intArr[0] = 1
	intArr[1] = 2
	intArr[2] = 3

	/*
	   Breakdown of 0:3 (colon syntax):

	   The colon : separates the start index and the end index.
	   0: This is the starting index, meaning the slice will start from index 0 (inclusive).
	   3: This is the ending index, meaning the slice will end at index 3 (exclusive), meaning it includes elements at indices 0, 1, and 2 but not index 3.

	   Inclusive Start, Exclusive End: The slice will include elements starting from the start index up to, but not including, the end index.
	*/

	// print range from 1st to 3rd element
	fmt.Println(intArr[0:3])

	// print range from 2nd to 3rd element
	fmt.Println(intArr[1:3])

	// print range from 1st to 3rd element
	fmt.Println(intArr[2:3])

	// with defined length
	var intArr2 = [3]int8{4, 5, 6}
	fmt.Println(intArr2)

	// with length infrared
	intArr3 := [...]int16{1, 2, 3, 4, 5, 6}
	fmt.Println(intArr3)

	// slices

	var intSlice = []int32{1, 2, 3, 4, 5, 6}
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Printf("\nAfter addeding 7 - The length is %v with capacity %v", len(intSlice), cap(intSlice))

	fmt.Printf("\nGetting value of 6th index - %v \n", intSlice[6])
	// however capacity is 12, but we won't be able to access 7th index position since length is 7 that means up till 6th index position value exists
	// uncomment below line to check
	// fmt.Printf("\nGetting value of 7th index - %v", intSlice[7]) // throw error panic: runtime error: index out of range [7] with length 7

	// joining two array
	sports := []string{"cricket", "hockey", "football"}
	fmt.Println(sports)
	aquaSports := []string{"Swimming", "Scuba Diving"}
	fmt.Println(append(sports, aquaSports...))
	mountainSports := []string{"Paragliding", "Mountaineering"}
	fmt.Println(append(sports, mountainSports...))

	// however while adding individual elements you can append as many as you want
	allSports := append(sports, "Gymnastics", "Boxing", "Badminton", "Archery", "Snowboarding")
	fmt.Println(allSports)

	// using make keyword
	var colors = make([]string, 5)
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "black"
	colors[3] = "white"
	colors[4] = "pink"

	fmt.Println(colors)

	// Create a slice of integers with length 5 and capacity 10
	numbers := make([]int, 5, 10)

	// Assign values to the slice
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	fmt.Println(numbers)      // Output: [1 2 3 0 0]
	fmt.Println(len(numbers)) // Output: 5 (length)
	fmt.Println(cap(numbers)) // Output: 10 (capacity)

	// Appending to slice
	numbers = append(numbers, 4, 5, 6)
	fmt.Println(numbers) // Output: [1 2 3 0 0 4 5 6]
}
