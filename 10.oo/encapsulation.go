package main

import "fmt"

type Encapsulation struct{
}

func(e *Encapsulation) Expose() {
	fmt.Println("i am  exposed")
}

func(e *Encapsulation) hide() {
	fmt.Println("this is a super secret")

}

func (e *Encapsulation) Unhide() {
	e.hide()
	fmt.Println("...jk")
}

func main() {
    e := Encapsulation{}

    e.Expose()     // "AHHHH! I'm exposed!"

    // e.hide()    // if uncommented, causes the following error
                   // ./main.go:10: e.hide undefined (cannot refer
                   // to unexported field or method encapsulation.
                   // (*Encapsulation)."".hide)

    e.Unhide()     // "Shhhh... this is super secret"
                   // "...jk"
}
