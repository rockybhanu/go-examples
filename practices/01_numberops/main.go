package main

import (
	"errors"
	"fmt"
)

func calculate() {
	var num1, num2 int
	fmt.Println("Enter first number:")

	if _, err := fmt.Scanln(&num1); err != nil {
		fmt.Println("Please enter a valid input")
		return
	}

	fmt.Println("Enter second number:")

	if _, err := fmt.Scanln(&num2); err != nil {
		fmt.Println("Please enter a valid input")
		return
	}

	calc(num1, num2)
	var isContinue string
	fmt.Print("Do you want to perform more calculation - Press y (small letter) and enter to continue Or Press anykey to terminate: ")
	fmt.Scan(&isContinue)
	if isContinue == "y" {
		calculate()
	} else {
		fmt.Println("Thanks for using Number operations Program, See you again!")
	}
}

func main() {
	fmt.Println("Welcome to number operation program !")
	calculate()
}

func calc(num1 int, num2 int) {
	fmt.Println(`Choose an operation:
1: Add
2: Subtract
3: Multiply
4: Divide
5: Modulus`)

	var op int

	fmt.Scan(&op)

	var output int
	var err error

	switch op {
	case 1:
		output = addition(num1, num2)
		fmt.Printf("Final output is %v \n", output)
	case 2:
		output = subtract(num1, num2)
		fmt.Printf("Final output is %v \n", output)
	case 3:
		output = multiply(num1, num2)
		fmt.Printf("Final output is %v \n", output)
	case 4:
		output, err = divide(num1, num2)
		if err != nil {
			fmt.Printf("Final output can not be calculated because %v \n", err.Error())
		} else {
			fmt.Printf("Final output is %v \n", output)
		}
	case 5:
		output = modulus(num1, num2)
		fmt.Printf("Final output is %v \n", output)
	}
}

func addition(num1 int, num2 int) int {
	fmt.Printf("Addition of %v with %v is %v \n", num1, num2, num1+num2)
	return num1 + num2
}

func subtract(num1 int, num2 int) int {
	fmt.Printf("subtraction from %v and %v is %v \n", num1, num2, num1-num2)
	return num1 - num2
}

func multiply(num1 int, num2 int) int {
	fmt.Printf("multiply of %v with %v is %v \n", num1, num2, num1*num2)
	return num1 * num2
}

func divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		fmt.Println("Second number is Zero and that is Denominator in division operations so this can not be processed")
		return 0, errors.New("denominator can not be Zero")
	}
	fmt.Printf("division of %v by %v is %v \n", num1, num2, num1/num2)
	return num1 / num2, nil
}

func modulus(num1 int, num2 int) int {
	fmt.Printf("modulus of %v with %v is %v \n", num1, num2, num1%num2)
	return num1 % num2
}
