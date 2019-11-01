package main

import (
	"fmt"
	"time"
)

func count() {
	// No parentheses :thinking_face
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		// Sleep a millisecond at a time
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	// Starts a application thread
	go count()

	// Sleep for a couple of milliseconds
	time.Sleep(time.Millisecond * 2)

	// Customary Hellooooo
	fmt.Println("Hello World")

	// Sleep a bit more for the thread to complete. Kinda crude way.
	time.Sleep(time.Millisecond * 5)
}
