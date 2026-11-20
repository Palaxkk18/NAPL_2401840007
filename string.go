package main

import (
	"fmt"
	"strings"
)

// Reverse reverses a string
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// CountVowels counts vowels
func CountVowels(s string) int {
	count := 0
	for _, ch := range strings.ToLower(s) {
		if strings.ContainsRune("aeiou", ch) {
			count++
		}
	}
	return count
}

// Factorial calculates factorial
func Factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Power calculates power
func Power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func main() {
	str := "Hello World"

	fmt.Println("Original String:", str)
	fmt.Println("Reversed String:", Reverse(str))
	fmt.Println("Vowel Count:", CountVowels(str))
	fmt.Println("Factorial of 5:", Factorial(5))
	fmt.Println("Power of 2^3:", Power(2, 3))
}
