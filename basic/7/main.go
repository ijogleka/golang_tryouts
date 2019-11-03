package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// configuration ...
// This is how `struct`s are defined in golang
// The fields are not command separated.
// Note: If the fields are uppercase'd, they can be accessed from outside
type configuration struct {
	Key   string
	Value string
}

func main() {
	// The file path search starts off from the `GOPATH`
	file, _ := os.Open("7/conf.json")

	// `defer` is a super cool thing. It gets executed
	// at the end of the current scope.
	// So even if the task `panic`s and fails at some point,
	// `defer` gets executed and makes sure that the file is closed
	defer file.Close()

	// Decode the JSON FILE using a json decoder
	// The methodology of using functions like `New..`
	// for constructing a instance of the class, is golang like.
	// Can think of it like a factory pattern
	decoder := json.NewDecoder(file)

	// Create a empty instance of the `configuration` struct
	conf := configuration{}

	// Write to the empty instance write from the file
	err := decoder.Decode(&conf)

	// Print a error if decode fails
	if err != nil {
		fmt.Println("Something failed with the json decode")
	}

	// Print Key and Value as read from the json file
	fmt.Println("Key:", conf.Key, "  Value: ", conf.Value)
}
