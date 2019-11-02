package main

import (
	"fmt"
	"sync" // This is the go package for synchronization
	"time"
)

var counter = 0 // Variable shared across threads. For the case of this example

var lock sync.Mutex // Plain old mutex
// sync provides sync.RWMutex as well

func main() {
	// Start 20 application threads
	for i := 0; i < 20; i++ {
		go incr()
	}

	// Wait for the threads to finish before exiting
	time.Sleep(time.Millisecond * 10)
}

func incr() {
	// Get a lock on the mutex. Block until we succeed
	lock.Lock()

	// Make sure to unlock the mutex when the thread ends.
	defer lock.Unlock()

	// Increment the counter variable value
	counter++

	// Print the current state of the counter
	fmt.Println(counter)
}

/*
  The mutex use inside the thread essentially lets the threads access the
  counter variable, one after the other, in place of parallely

  Without the mutex lock in the thread, the counter can get printed
  in any order and may not increment up to 20

  With the mutex lock in the thread, the counter will get printed
  in increasing order and will always increment up to 20
*/
