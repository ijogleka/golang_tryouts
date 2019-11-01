package main

import (
	"fmt"
	"time"
)

// input argument is a channel
// NOTE: Since the first letter of the function is smallcase
// this function is not accessible outside this file. :ok_hand
func printCount(c chan int) {
	num := 0

	// Loop till the value of the variable `num` is >= 0
	for num >= 0 {
		// Block the thread until the channel `c` receives a value
		num = <-c
		fmt.Print(num, " ")
	}
}

// Function prints the slice data until a value < 0 is encountered
func main() {
	fmt.Println("Buckle up son")

	// Create a channel for thread to thread data passing
	c := make(chan int)

	// Just a simple initialized slice (a.k.a. C++ vector)
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}

	// Start a application thread and pass in the channel for the thread
	// to use, to consume data or pass data back
	go printCount(c)

	// For loop for iterating over all the elements of the slice
	// Format: `for <index>, <object> := range <slice> {}`
	for _, v := range a {
		// Our main thread, will be passing data to the `printCount` application
		// thread, via the channel `c`
		// I wonder if the values being passed to the channel need to be of native types
		c <- v
	}

	// Sleep, the crude way, for the application threads to die
	time.Sleep(time.Millisecond * 1)

	// We done
	fmt.Println("We done")
}
