package main

import "fmt"

// Non-exportable struct type `car`
type car struct {
	Model string
	Year  int
}

// Print ...
// This function is attached to the struct `car` in the following way
// We say that the `car, receives the function Print`
func (c *car) Print() {
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
}

func main() {
	// Create a instance of a car
	c := car{"Honda", 1970}

	// Print the car metadata using the attached function `Print`
	c.Print()
}
