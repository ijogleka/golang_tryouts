package main

import "fmt"

type car struct {
	Model string
	Year  int
}

func valueHandler(v car) { // `v` is a copy of the passed in object
	fmt.Println("Value:", v)
}

func pointerHandler(p *car) { // `p` is a copy of the passed in `pointer to the object`
	fmt.Println("Pointer:", p)
	p.Year = 2001
}

func main() {
	value1 := car{"Honda", 1970}      // Store the value of the object in the variable
	pointer1 := &car{"Ferrari", 1971} // Store the pointer to the value of the object in the variable

	fmt.Println("[BEFORE] Value:", value1, " ", "Pointer:", pointer1)
	// This prints: `[BEFORE] Value: {Honda 1970}   Pointer: &{Ferrari 1971}`

	valueHandler(value1)     // Passing a object by value to a function
	pointerHandler(pointer1) // Passing a pointer to a object, by value to a function

	fmt.Println("[AFTER] Value:", value1, " ", "Pointer:", pointer1)
	// This prints: `[AFTER] Value: {Honda 1970}   Pointer: &{Ferrari 2001}`
	// The `pointerHandler` function modifies the object `Year` value
}
