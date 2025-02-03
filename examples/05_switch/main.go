package main

import (
	"errors"
	"fmt"
	"log"
)

// main function: Entry point of the program
func main() {
	log.Println("Starting the program...")

	// Calling dayOfTheWeek function with 5
	day, err := dayOfTheWeek(5)
	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Println("Day of the week:", day)
	}

	// Calling grading function with different marks
	fmt.Println("Grade:", grading(75))
	fmt.Println("Grade:", grading(0))

	log.Println("Program completed successfully")
}

// dayOfTheWeek returns the name of the week based on a given number (1-7)
// Uses an expression-based switch
func dayOfTheWeek(dayOfTheMonth int) (string, error) {
	var day string
	switch dayOfTheMonth {
	case 1:
		day = "Sunday"
	case 2:
		day = "Monday"
	case 3:
		day = "Tuesday"
	case 4:
		day = "Wednesday"
	case 5:
		day = "Thursday"
	case 6:
		day = "Friday"
	case 7:
		day = "Saturday"
	default:
		return "", errors.New("invalid input: day should be between 1 and 7")
	}
	return day, nil
}

// grading function returns the grade based on total marks
// Uses a condition-based switch (without an expression)
func grading(totalMarks int) string {
	switch {
	case totalMarks < 0 || totalMarks > 100:
		return "Invalid marks: should be between 0 and 100"
	case totalMarks <= 35:
		return "Failed"
	case totalMarks <= 40:
		return "Third Division"
	case totalMarks <= 60:
		return "Second Division"
	case totalMarks <= 80:
		return "First Division"
	case totalMarks <= 100:
		return "Distinction"
	}
	return "Unknown grade" // This will never execute, but it's a safe fallback
}
