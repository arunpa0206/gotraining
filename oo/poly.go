package main

import (
	"fmt"
)

// This is the interface. Types that wish to confirm to Vehicle must implement
// this interface. Implementing this interface is done by implementing the
// methods defined here. `description()` and `medium()`.
type Vehicle interface {
	description() string
	medium() string
}

// This is a simple struct type.
type Car struct {
	doors int
}

// This is a simple struct type.
type Boat struct {
	rudders int
	masts   int
}

// By implementing this interface method, Car is starting to conform to the
// Vehicle interface.
func (c Car) description() string {
	return fmt.Sprintf("I am a car with %v door(s). I move on the %v.", c.doors, c.medium())
}

// By implementing this interface method, Car now completely conforms to the
// Vehicle interface.
func (c Car) medium() string {
	return "road"
}

// By implementing this interface method, Boat is starting to conform to the
// Vehicle interface.
func (b Boat) description() string {
	return fmt.Sprintf("I am a boat with %v rudder(s) and %v mast(s). I ride on %v.", b.rudders, b.masts, b.medium())
}

// By implementing this interface method, Boat now completely conforms to the
// Vehicle interface.
func (b Boat) medium() string {
	return "water"
}

// Note that an interface can either be a value-type or a reference-type, so
// we use `Vehicle` here, not `*Vehicle`.
// See http://openmymind.net/Things-I-Wish-Someone-Had-Told-Me-About-Go/
func ride(vehicle Vehicle) {
	fmt.Println(fmt.Sprintf("Literal Vehicle: Value is %v. Type is %T.", vehicle, vehicle))
	fmt.Println(vehicle.description())
	fmt.Println(fmt.Sprintf("Did you notice that the vehicle rides on %v?", vehicle.medium()))
}

// Type checking for a vehicle.
func checkType(vehicle Vehicle) {
	test1, ok1 := vehicle.(*Car)

	fmt.Println(fmt.Sprintf("Is %v a *Car? %v.", test1, ok1))

	test2, ok2 := vehicle.(Car)

	fmt.Println(fmt.Sprintf("Is %v a Car? %v.", test2, ok2))

	test3, ok1 := vehicle.(*Boat)

	fmt.Println(fmt.Sprintf("Is %v a *Boat? %v.", test3, ok1))

	test4, ok2 := vehicle.(Boat)

	fmt.Println(fmt.Sprintf("Is %v a Boat? %v.", test4, ok2))

	test5, ok3 := vehicle.(Vehicle)

	fmt.Println(fmt.Sprintf("Is %v a Vehicle? %v.", test5, ok3))
}

func main() {
	// Using & returns a pointer to the struct value.
	car := &Car{
		doors: 4,
	}

	// Not using & returns the struct value.
	boat := Boat{
		rudders: 1,
		masts:   3,
	}

	fmt.Println("Ride in the car!")
	ride(car)
	fmt.Println()

	fmt.Println("Ride in the boat!")
	ride(boat)
	fmt.Println()

	fmt.Println("Check the type of the car.")
	checkType(car)
	fmt.Println()

	fmt.Println("Check the type of the boat.")
	checkType(boat)
	fmt.Println()
}
