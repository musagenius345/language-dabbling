package main

import (
	"fmt"
)

// Basic Function
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Function with Return Value
func add(a, b int) int {
	return a + b
}

// Function with Multiple Return Values
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Variadic Function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Anonymous Function (Closure)
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Function as Parameter
func apply(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}

func main() {
	greet("Gopher")

	result := add(3, 5)
	fmt.Println("3 + 5 =", result)

	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", quotient)
	}

	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", total)

	// Anonymous Function Example
	counterFn := counter()
	fmt.Println(counterFn()) // 1
	fmt.Println(counterFn()) // 2

	// Function as Parameter Example
	product := apply(func(a, b int) int {
		return a * b
	}, 4, 3)
	fmt.Println("Product:", product)
}
