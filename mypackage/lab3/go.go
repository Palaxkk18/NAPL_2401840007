package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	go sayHello()

	// Keep the main goroutine running
	fmt.Scanln()
}
