package main

import "fmt"

type Fooer interface {
	Foo1()
	Foo2()
	Foo3()
}

type foo struct {
}

func (f foo) Foo1() {
	fmt.Println("Foo1() here")
}

func (f foo) Foo2() {
	fmt.Println("Foo2() here")
}

func (f foo) Foo3() {
	fmt.Println("Foo3() here")
}

func NewFoo() Fooer {
	return &foo{}
}

func main() {
	f := NewFoo()

	f.Foo1()

	f.Foo2()

	f.Foo3()

}
