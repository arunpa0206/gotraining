package main

import "fmt"

type purchase interface {
	sell()
}

type display interface {
	show()
}

// We use the two previous
// interfaces to form
// The following composite
// interface through embedding
type salesman interface {
	purchase
	display
}

type game struct {
	name, price    string
	gameCollection []string
}

// We use the game struct to
// implement the interfaces
func (t game) sell() {
	fmt.Println("--------------------------------------")
	fmt.Println("Name:", t.name)
	fmt.Println("Price:", t.price)
	fmt.Println("--------------------------------------")
}
func (t game) show() {
	fmt.Println("The Games available are: ")
	for _, name := range t.gameCollection {
		fmt.Println(name)
	}
	fmt.Println("--------------------------------------")
}

// This method takes the composite
// interface as a parameter
// Since the interface is composed
// of purchase and display
// Hence the child methods of those
// interfaces can be accessed here
func shop(s salesman) {
	fmt.Println(s)
	s.sell()
	s.show()
}

func main() {

	collection := []string{"XYZ",
		"Trial by Code", "Sea of Rubies"}
	game1 := game{"ABC", "$125", collection}
	shop(game1)

}
