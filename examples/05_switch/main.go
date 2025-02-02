package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(dayOfTheWeek(5))
}

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
		day = "Thrusday"
	case 6:
		day = "Friday"
	case 7:
		day = "Saturday"
	default:
		return day, errors.New("incorrect day of the month")
	}
	return day, nil
}
