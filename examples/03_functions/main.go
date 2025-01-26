package main

import (
	"errors"
	"fmt"
)

// main function - Entry point of the Go application
// The execution of the program starts from here.
func main() {
	fmt.Println("Starting the Go program...")

	// Calling a no-argument function
	goDev()

	// Calling a function with a single argument
	hello("Gopher")

	// Calling a function with multiple arguments
	helloWithName("Dear", "Gopher")

	// Demonstrating function with single return value
	numerator := 25
	denominator := 5
	fmt.Println("\nPerforming division without remainder:")
	fmt.Printf("Numerator = %d\n", numerator)
	fmt.Printf("Denominator = %d\n", denominator)

	result, err := divideInt(numerator, denominator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result of division: %d\n", result)
	}

	// Handling division by zero scenario
	numerator = 25
	denominator = 0
	fmt.Println("\nAttempting division with zero denominator:")
	fmt.Printf("Numerator = %d\n", numerator)
	fmt.Printf("Denominator = %d\n", denominator)

	res, er := divideInt(numerator, denominator)
	if er != nil {
		fmt.Println("Error:", er)
	} else {
		fmt.Printf("Result of division: %d\n", res)
	}

	// Demonstrating function with multiple return values (quotient and remainder)
	numerator = 29
	denominator = 4
	fmt.Println("\nPerforming division with remainder capture:")
	fmt.Printf("Numerator = %d\n", numerator)
	fmt.Printf("Denominator = %d\n", denominator)

	result, remainder, err := divideIntCaptureRemainder(numerator, denominator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result of division: %d, Remainder: %d\n", result, remainder)
	}
}

// goDev is a no-argument function that prints information about Go language creators.
func goDev() {
	fmt.Println("\nGo was designed by Robert Griesemer, Rob Pike, and Ken Thompson.")
}

// hello is a single-argument function that greets the provided address.
func hello(address string) {
	fmt.Printf("\nHello, %v!\n", address)
}

// helloWithName is a multi-argument function that prints a personalized greeting.
func helloWithName(address string, name string) {
	fmt.Printf("\nHello %v, learning to become %v!\n", address, name)
}

// divideInt performs integer division and returns the quotient or an error if division by zero occurs.
func divideInt(numerator int, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("error: cannot divide by zero")
	}
	return numerator / denominator, nil
}

// divideIntCaptureRemainder performs integer division and captures both quotient and remainder.
func divideIntCaptureRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("error: denominator is zero")
	}
	return numerator / denominator, numerator % denominator, nil
}
