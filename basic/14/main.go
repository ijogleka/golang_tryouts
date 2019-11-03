package main

import "fmt"

// Sum is a first class function type
// that takes in two integers and returs a integer
type Sum func(a int, b int) int

// @param s [First class function of type `Sum`]
// @param left [int]
// @param right [int]
func process(s Sum, left int, right int) int {
	return s(left, right)
}

func main() {
	// Functions are first class types. They can be stored as variables
	summer := func(l int, r int) int { return l + r }

	// Functions can be passed as arguments to other functions,
	// since they are first-class types in golang
	outputSum := process(summer, 12, 13)

	// Test
	fmt.Println("Output Sum:", outputSum)
}
