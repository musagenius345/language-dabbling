package main

import (
	"fmt"
	"net/http"
	"time"
	"strings"
	"io/ioutil"
	"encoding/json"
	"os"
	"log"
	"regexp"
	"sort"
)

func main() {
	// Time and Date
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime.Format("2006-01-02 15:04:05"))

	// HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go Standard Library!")
	})
	http.ListenAndServe(":8080", nil)

	// Strings
	s := "Hello, Go!"
	fmt.Println(strings.Contains(s, "Go"))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.TrimSpace("   Hello   "))

	// File I/O
	content := "Hello, File I/O!"
	err := ioutil.WriteFile("file.txt", []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Content:", string(data))

	// JSON Encoding/Decoding
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	person := Person{Name: "Alice", Age: 25}
	jsonData, _ := json.Marshal(person)
	fmt.Println("JSON Data:", string(jsonData))

	// Environment Variables
	envVar := os.Getenv("HOME")
	fmt.Println("Home Directory:", envVar)

	// Regular Expressions
	re := regexp.MustCompile(`\d{3}-\d{2}-\d{4}`)
	fmt.Println("Valid SSN:", re.MatchString("123-45-6789"))

	// Sorting
	numbers := []int{4, 2, 7, 1, 5}
	sort.Ints(numbers)
	fmt.Println("Sorted Numbers:", numbers)
}
