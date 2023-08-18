package main

import (
	"fmt"
	"math"
)

// Higher-Order Function: Apply a function to each element of a slice
func mapInts(numbers []int, fn func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = fn(num)
	}
	return result
}

// Pure Function: Square a number (no side effects)
func square(x int) int {
	return x * x
}

// Closure: Generate a function that adds a constant
func adder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// Immutability: Calculate the sum of squares using reduce
func reduceInts(numbers []int, fn func(int, int) int, initialValue int) int {
	result := initialValue
	for _, num := range numbers {
		result = fn(result, num)
	}
	return result
}

// Currying: Create a curried function to calculate powers
func power(x int) func(int) int {
	return func(y int) int {
		return int(math.Pow(float64(x), float64(y)))
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Using higher-order function and pure function
	squaredNumbers := mapInts(numbers, square)
	fmt.Println("Squared numbers:", squaredNumbers)

	// Using closure to create custom adder functions
	addTwo := adder(2)
	addFive := adder(5)
	fmt.Println("Add 2 to 10:", addTwo(10))
	fmt.Println("Add 5 to 10:", addFive(10))

	// Using reduce to calculate sum of squares
	sumOfSquares := reduceInts(numbers, func(a, b int) int {
		return a + square(b)
	}, 0)
	fmt.Println("Sum of squares:", sumOfSquares)

	// Using currying to calculate powers
	powerOfTwo := power(2)
	powerOfThree := power(3)
	fmt.Println("2^3:", powerOfTwo(3))
	fmt.Println("3^2:", powerOfThree(2))
}
