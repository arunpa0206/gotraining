package main

import "fmt"

// Functions
func add(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Result: ", res)
	return res
}

// Main function
func main() {
	defer fmt.Println("End")
	add(34, 56)
}
