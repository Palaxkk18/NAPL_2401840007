package main

import "fmt"

func changeValue(num *int) {
	*num = 100
}

type Student struct {
	Name string
	Age  int
}

func main() {

	x := 50

	fmt.Println("Address of x:", &x)

	ptr := &x
	fmt.Println("Value of x using *:", *ptr)

	fmt.Println("Before modification:", x)

	changeValue(&x)

	fmt.Println("After modification:", x)

	// Allocate struct
	student := new(Student)

	// Access and modify fields
	student.Name = "Palak"
	student.Age = 19

	fmt.Println("Student Name:", student.Name)
	fmt.Println("Student Age:", student.Age)
}
