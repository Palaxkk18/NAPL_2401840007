package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks float64
}

func bonus(s *Student, grace_marks float32) {
	s.Marks += float64(grace_marks)
}

func main() {
	Asha := Student{
		Name:  "Asha",
		Age:   22,
		Marks: 67.0,
	}
	fmt.Printf("Marks before adding bonus: %g\n", Asha.Marks)
	bonus(&Asha, 13.5)
	fmt.Printf("Marks after adding bonus: %g\n", Asha.Marks)
}
