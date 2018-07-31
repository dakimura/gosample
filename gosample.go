package main

import "fmt"

// Add returns the sum of argument values
func Add(i ...int) int {
	sum := 0
	fmt.Println(i)
	for _, val := range i {
		sum += val
	}
	return sum
}

func main() {
	fmt.Printf("Hello world\n")
	fmt.Println(Add(1, 2))
}
