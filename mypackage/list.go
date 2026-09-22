package main

import "fmt"

func main() {
	s := []int{10, 20, 30}
	var sum int
	for _, v := range s {
		sum += v
	}
	fmt.Println("Sum:", sum)
	avg := sum / 3
	fmt.Println("Average:", avg)
}
