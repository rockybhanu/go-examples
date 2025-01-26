package main

import "fmt"

// main function - Entry point of the Go application
func main() {
	fmt.Println("Starting the Go program to demonstrate if-else control flow...\n")

	// Example 1: Simple if-else
	checkEvenOdd(10)
	checkEvenOdd(7)

	// Example 2: if-else if-else for grading system
	checkGrade(85)
	checkGrade(65)
	checkGrade(45)
	checkGrade(20)

	// Example 3: Nested if-else example
	checkVotingEligibility(20, true)
	checkVotingEligibility(15, false)

	fmt.Println("\nProgram execution completed.")
}

// checkEvenOdd demonstrates simple if-else control flow by checking if a number is even or odd.
func checkEvenOdd(num int) {
	fmt.Printf("Checking if %d is even or odd...\n", num)
	if num%2 == 0 {
		fmt.Println(num, "is an even number.")
	} else {
		fmt.Println(num, "is an odd number.")
	}
	fmt.Println()
}

// checkGrade evaluates a student's grade based on score using if-else
func checkGrade(score int) {
	fmt.Printf("Evaluating grade for score: %d\n", score)
	if score >= 90 {
		fmt.Println("Grade: A (Excellent)")
	} else if score >= 75 {
		fmt.Println("Grade: B (Good)")
	} else if score >= 50 {
		fmt.Println("Grade: C (Average)")
	} else {
		fmt.Println("Grade: F (Fail)")
	}
	fmt.Println()
}

// checkVotingEligibility checks if a person is eligible to vote based on age and citizenship status
func checkVotingEligibility(age int, isCitizen bool) {
	fmt.Printf("Checking voting eligibility for age %d and citizenship status %v\n", age, isCitizen)
	if age >= 18 {
		if isCitizen {
			fmt.Println("You are eligible to vote.")
		} else {
			fmt.Println("You meet the age requirement, but you must be a citizen to vote.")
		}
	} else {
		fmt.Println("You are not eligible to vote due to age restrictions.")
	}
	fmt.Println()
}
