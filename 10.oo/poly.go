package main

import "fmt"

type Human interface {
    myStereotype() string
}

type Man struct {
}

func (m Man) myStereotype() string {
    return "I'm going fishing."
}

type Woman struct {
}

func (m Woman) myStereotype() string {
    return "I'm going shopping."
}
func main() {
    m := new (Man)
    w := new (Woman)

    //an array of Humans - we don’t know whether Man or Woman
    hArr := [...]Human{m, w} //array of 2 Humans. One is the type Man, one is the type Woman.
    for n, _ := range (hArr) {

        fmt.Println("I'm a human, and my stereotype is: ", hArr[n].myStereotype())   //appears as human type, but behavior changes depending on actual instance
        
    }

}