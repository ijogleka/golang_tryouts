package main

import (
	"errors"
	"fmt"
)

func main() {
	// NOTE: `defer` function gets called at the end of the scope that
	// its defined in. This provides programmers a good way of not forgetting
	// to cleanup certain things (that they should cleanup at the end of a task)
	defer func() {
		// NOTE: `recover` is a golang function provided to assist with recovering
		// from a `panic`. Like a `catch` block in most languages.
		// Stack unwinding is stopped right here, thanks to this function
		if err := recover(); err != nil {
			fmt.Println("Recovering from a panic caused by, ", err)
		}
	}()

	// NOTE: Golang way of throwing errors. `panic`
	// The function takes in a object of any type in as a argument.
	// Many golang programmers use a `error` object here as a standard,
	// but its not mandatory
	// NOTE: `panic` from a function causes stack unwinding
	panic(errors.New("This is a test error that we are trying to panic about"))
}

// NOTE:
/*
  If panics are thrown inside goroutines (application threads)
  they should be receovered inside the application threads.
  If they are not recovered there, the application threads dont have a
  way to propagate the error back to the main thread (for recovery there for example)
  and the program may crash
*/
