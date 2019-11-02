package main

import "fmt"

// Car ...(exportable since its uppercase)
type Car struct {
	Model string
	Year  int
}

// NewCar ...(exportable since its uppercase)
// This is a replacement for a struct constructor
// using somewhat of a factory pattern
// Standard Format for such a function is
// func New<struct type>(input arguments) (return pointer or value) {
//   return <struct type>{input arguments} // Value
//   OR
//   return &<struct type>{input arguments} // Pointer to a value
// }
func NewCar(model string, year int) *Car {
	return &Car{model, year}
}

func main() {
	c := NewCar("Honda", 1980)

	fmt.Println("Car instance: ", c)
	// This will print `Car instance:  &{Honda 1980}`
}
