package main

import "fmt"

const (
	b = iota
	c
)
const (
	a2 = iota
)

func main() {

	//fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a2)
}
