package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p *Person) ReadData() {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter Name: ")
	fmt.Scanln(&p.Name)
	reader.ReadLine()
	fmt.Print("Enter Age: ")
	fmt.Scanln(&p.Age)
	reader.ReadLine()
	fmt.Print("Enter Job: ")
	fmt.Scanln(&p.Job)
	reader.ReadLine()
	fmt.Print("Enter Salary: ")
	fmt.Scanln(&p.Salary)
	reader.ReadLine()
}

func (p *Person) Show() {
	fmt.Println("\nPerson Details:")
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Job:", p.Job)
	fmt.Printf("Salary: %.2f\n", p.Salary)
}

func main() {

	var p1 Person
	var p2 Person

	fmt.Println("Enter details for Person 1:")
	p1.ReadData()
	p1.Show()

	fmt.Println("\nEnter details for Person 2:")
	p2.ReadData()
	p2.Show()
}
