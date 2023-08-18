package main

import "fmt"

func main() {
	// Arrays: Fixed-size, contiguous memory
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println("Array:", arr)

	// Slices: Dynamic arrays, can grow and shrink
	slice := []int{3, 4, 5}
	slice = append(slice, 6)
	fmt.Println("Slice:", slice)

	// Maps: Key-value pairs
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	ages["Charlie"] = 28
	fmt.Println("Ages:", ages)

	// Structs: Custom data types
	type Person struct {
		Name   string
		Age    int
		Gender string
	}
	person1 := Person{Name: "Alice", Age: 25, Gender: "Female"}
	person2 := Person{Name: "Bob", Age: 30, Gender: "Male"}
	fmt.Println("Person 1:", person1)
	fmt.Println("Person 2:", person2)
}
