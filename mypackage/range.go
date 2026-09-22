package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["palak"] = 19
	ages["john"] = 25
	ages["jane"] = 30

	for k, v := range ages {
		fmt.Println(k, v)
	}
}
