package main

import (
	"fmt"
	"slices"
)

func main() {

	// Implementing slice
	s := []int{}
	fmt.Println("empty slice:", s)
	s = append(s, 10, 20, 30, 40, 50)
	fmt.Println("slice after append:", s)
	fmt.Println("original slice:", s)

	// Remove element at index 1
	s = slices.Delete(s, 1, 2)
	fmt.Println("slice after deletion:", s)

// Update element at index 0
s[0] = 50
fmt.Println("slice after updating:", s)
fmt.Println("final slice:", s)

// Implementing maps

subjects := map[string]int{}

// Inserting values in map
subjects["maths"] = 90
subjects["Coa"] = 80
subjects["os"] = 70

fmt.Println("map:", subjects)

// Deleting value from map
delete(subjects, "maths")

fmt.Println("map after deletion:", subjects)

// Lookup value in map
for k, v := range subjects {
	fmt.Println(k, v)

}