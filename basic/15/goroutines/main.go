package main

import (
	"fmt"
	"time"
)

// Counter ...
func Counter() {
	// Just print the index 5 times, over roughly 5 milliseconds
	for i := 0; i < 5; i++ {
		// Print something indicating the progress of the application thread
		fmt.Println("Index:", i)

		// Sleep for 1 millisecond at a time
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	// Start off the application thread a.k.a. goroutine
	go Counter()

	// Wait for 3 milliseconds in the main thread
	time.Sleep(time.Millisecond * 3)

	// Print something to indicate the execution of the main thread
	fmt.Println("Inside the main thread")

	// Wait for 3 more milliseconds to allow for the application thread
	// to finish off, before exiting the function
	time.Sleep(time.Millisecond * 2)

	// Note: If the main thread closes before the application thread,
	// the app threads continue running in the background as daemons..
	// basically
}
