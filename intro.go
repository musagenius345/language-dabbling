// Learn Golang in 300 Lines - For JavaScript Developers

package main

import (
	"fmt"
	"math"
)

// Function - Just like JavaScript
func greet(name string) string {
	return "Hello, " + name + "!"
}

// Variables - Strongly typed, similar to `let` in JS
func variables() {
	var num int = 42
	str := "Golang"
	fmt.Println(num, str)
}

// Conditional Statements - if-else, similar to JavaScript
func conditionals() {
	num := 10
	if num > 5 {
		fmt.Println("Number is greater than 5")
	} else {
		fmt.Println("Number is 5 or less")
	}
}

// Loops - for loop, similar to JavaScript
func loops() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

// Arrays and Slices - Arrays have fixed size, Slices are dynamic
func arraysAndSlices() {
	arr := [3]int{1, 2, 3}
	slice := []int{4, 5, 6}
	fmt.Println(arr[0], slice[1])
}

// Maps - Similar to JavaScript objects
func maps() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Println(ages["Alice"])
}

// Structs - Custom data types, like JavaScript objects
type Person struct {
	Name string
	Age  int
}

func structs() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Println(p.Name)
}

// Functions as First-Class Citizens - Can be passed as arguments
func higherOrderFunc(add func(int, int) int) int {
	return add(3, 5)
}

func main() {
	fmt.Println(greet("Gopher"))
	variables()
	conditionals()
	loops()
	arraysAndSlices()
	maps()
	structs()

	sum := higherOrderFunc(func(a, b int) int {
		return a + b
	})
	fmt.Println("Sum:", sum)

	// Using Math package - similar to JavaScript Math object
	sqrt := math.Sqrt(16)
	fmt.Println("Square root:", sqrt)
}
