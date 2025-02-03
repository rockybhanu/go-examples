package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	log.Println("Starting maps illustration...")

	// 1️⃣ Declaring an empty map using make()
	var countryCapital = make(map[string]string)
	countryCapital["India"] = "New Delhi"
	countryCapital["UK"] = "London"
	log.Println("Country-Capital Map:", countryCapital)
	// Output: map[India:New Delhi UK:London]

	// 2️⃣ Using shorthand declaration (map literal)
	tShirtSize := map[string]uint8{
		"S":   27,
		"M":   35,
		"L":   40,
		"XL":  43,
		"XXL": 46,
	}
	log.Printf("T-shirt size for XL is %d\n", tShirtSize["XL"])
	// Output: T-shirt size for XL is 43

	// 3️⃣ Using make() function with shorthand notation
	priceList := make(map[string]int)
	priceList["S"] = 140
	priceList["M"] = 240
	priceList["L"] = 300
	priceList["XL"] = 340
	priceList["XXL"] = 400
	log.Println("Price List Map:", priceList)
	// Output: map[S:140 M:240 L:300 XL:340 XXL:400]

	// 4️⃣ Declaring an empty map without make() (Will be nil)
	var emptyMap map[string]int
	log.Println("Empty map (nil):", emptyMap)
	// Output: map[]

	// ⚠️ emptyMap["test"] = 1 // ❌ This will panic because it's a nil map!

	// 5️⃣ Checking if a key exists in a map
	ageMap := map[string]int{"Alice": 25, "Bob": 30}
	age, exists := ageMap["Charlie"]
	if exists {
		fmt.Println("Charlie's age:", age)
	} else {
		fmt.Println("Charlie not found") // Output: Charlie not found
	}

	// 6️⃣ Deleting a key from a map
	delete(ageMap, "Bob")
	log.Println("Updated Age Map after deleting 'Bob':", ageMap)
	// Output: map[Alice:25]

	// 7️⃣ Nested Map (map inside a map)
	users := map[string]map[string]string{
		"alice": {"email": "alice@example.com", "city": "New York"},
		"bob":   {"email": "bob@example.com", "city": "Los Angeles"},
	}
	log.Println("Nested Map (users):", users)
	// Output: map[alice:map[email:alice@example.com city:New York] bob:map[email:bob@example.com city:Los Angeles]]

	// Accessing a nested map value
	fmt.Println("Alice's email:", users["alice"]["email"])
	// Output: Alice's email: alice@example.com

	// 8️⃣ Using a map inside a struct
	type Employee struct {
		Name string
		Age  int
	}

	employees := map[int]Employee{
		101: {"Alice", 25},
		102: {"Bob", 30},
	}
	fmt.Println("Employee ID 101:", employees[101])
	// Output: Employee ID 101: {Alice 25}

	// 9️⃣ Using a map with slices (one key has multiple values)
	studentSubjects := map[string][]string{
		"Alice": {"Math", "Science"},
		"Bob":   {"History", "English"},
	}
	log.Println("Student Subjects Map:", studentSubjects)
	// Output: map[Alice:[Math Science] Bob:[History English]]

	// 1️⃣0️⃣ Encoding a map to JSON (useful for APIs)
	person := map[string]string{"name": "Alice", "age": "25"}
	jsonData, _ := json.Marshal(person)
	fmt.Println("JSON representation of person:", string(jsonData))
	// Output: {"name":"Alice","age":"25"}

	log.Println("Maps illustration completed successfully.")
}
