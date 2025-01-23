package main

import (
	"fmt"
	"math"
)

func main() {

	/*
		const myConst
		var myVar

		bool
		float32 float64
		int int16 int32 int64
		rune
		string
		uint uint8 uint16 uint32 uint64
	*/

	// bool
	var isRaining bool
	fmt.Println(isRaining) // default to false

	isCompleted := true // shorthand declarations
	fmt.Println(isCompleted)

	//uint
	// shorthand declarations , auto type detection
	x := 1
	fmt.Println(x)

	// declarations using var with type
	var y uint = 2
	fmt.Println(y)

	// single line multiple declarations
	var p, q, r uint = 1, 2, 3
	fmt.Println(p, q, r)

	//uint8 limit 255
	/*
		# go-examples/examples/01_variables_constants
		./main.go:40:16: cannot use 256 (untyped int constant) as uint8 value in variable declaration (overflows)
	*/
	var a uint8 = 255 // make it 256 to learn limit
	fmt.Println(a)

	/*
		^ is the bitwise NOT operator.
		uint8(0) is 0 as an 8-bit unsigned integer (binary 00000000), and ^uint8(0) flips the bits to 11111111 (decimal 255).
		The same logic applies to uint16, uint32, and uint64.
	*/

	maxUint8 := ^uint8(0)
	maxUint16 := ^uint16(0)
	maxUint32 := ^uint32(0)
	maxUint64 := ^uint64(0)

	fmt.Println("Max of uint8  :", maxUint8)
	fmt.Println("Max of uint16 :", maxUint16)
	fmt.Println("Max of uint32 :", maxUint32)
	fmt.Println("Max of uint64 :", maxUint64)

	/*
		^uint8(0) - This results in 0xFF (255 in decimal).
		>> 1 - Right shift by 1 gives us 0x7F (127 in decimal for int8), which is the maximum value of int8.
		Negative values: In two's complement, the minimum value of a signed integer is -max - 1. Hence, we compute it as -maxInt8 - 1.
		so for example -
		Imagine you have an 8-bit number, which means it can hold values using 8 tiny switches, each being either 0 (off) or 1 (on).
		When you write uint8(0), it means all switches are off, looking like 00000000 in binary.
		Now, when you use the ^ symbol, it flips all these switches to 1, turning it into 11111111, which is the biggest number you can store in 8 bits, equal to 255 in decimal (or 0xFF in hexadecimal).
		Next, when we shift this number one step to the right using >> 1, all the bits move one position, making it 01111111.
		This new number represents 127, which is the highest value for an 8-bit signed number (int8) because the first bit is used to show if the number is positive or negative.
		In a system called "two's complement," which computers use to store negative numbers, the smallest value is always one more negative than the biggest positive value.
		So, if the biggest number is 127, the smallest number becomes -128, calculated as -127 - 1.
		This is why in an 8-bit signed integer, the range goes from -128 to 127.
	*/

	maxInt8 := int8(^uint8(0) >> 1)
	maxInt16 := int16(^uint16(0) >> 1)
	maxInt32 := int32(^uint32(0) >> 1)
	maxInt64 := int64(^uint64(0) >> 1)

	minInt8 := -maxInt8 - 1
	minInt16 := -maxInt16 - 1
	minInt32 := -maxInt32 - 1
	minInt64 := -maxInt64 - 1

	fmt.Println("Max of int8  :", maxInt8)
	fmt.Println("Max of int16 :", maxInt16)
	fmt.Println("Max of int32 :", maxInt32)
	fmt.Println("Max of int64 :", maxInt64)

	fmt.Println("Min of int8  :", minInt8)
	fmt.Println("Min of int16 :", minInt16)
	fmt.Println("Min of int32 :", minInt32)
	fmt.Println("Min of int64 :", minInt64)

	// Maximum values for float32 and float64
	maxFloat32 := math.MaxFloat32
	maxFloat64 := math.MaxFloat64

	// Minimum positive values (closest to zero)
	minPosFloat32 := math.SmallestNonzeroFloat32
	minPosFloat64 := math.SmallestNonzeroFloat64

	/*
		The number 3.4028235e+38 (scientific notation) for float32 - 3.4028235 × 10³⁸

		float32 can store values approximately in the range:
		Positive range: 1.401298e-45 (smallest positive) to 3.4028235e+38 (largest positive).
		Negative range: -3.4028235e+38 to -1.401298e-45.

		Precision Considerations:
		While float32 can store numbers like 3.5 without issue, it has limited precision (about 7 decimal digits of accuracy).
		If you need higher precision (up to 15-16 digits), you should use float64.
	*/
	fmt.Println("Max of float32:", maxFloat32)
	fmt.Println("Max of float64:", maxFloat64)

	fmt.Println("Min positive of float32:", minPosFloat32)
	fmt.Println("Min positive of float64:", minPosFloat64)
}
