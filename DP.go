package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"apple", "banana", "cherry"}

	// Using map to transform each element in the list
	uppercased := Map(words, strings.ToUpper)
	fmt.Println(uppercased) // [APPLE BANANA CHERRY]

	// Using filter to select elements that meet a condition
	filtered := Filter(words, func(s string) bool {
		return strings.HasPrefix(s, "b")
	})
	fmt.Println(filtered) // [banana]
}

// Map applies a function to each element in a slice
func Map(slice []string, fn func(string) string) []string {
	result := make([]string, len(slice))
	for i, val := range slice {
		result[i] = fn(val)
	}
	return result
}

// Filter selects elements from a slice that satisfy a condition
func Filter(slice []string, predicate func(string) bool) []string {
	result := []string{}
	for _, val := range slice {
		if predicate(val) {
			result = append(result, val)
		}
	}
	return result
}
