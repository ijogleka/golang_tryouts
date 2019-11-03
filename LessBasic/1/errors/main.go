package main

import (
	"errors"
	"fmt"
)

// Divide ...
func Divide(Num int, Den int) (int, bool) {
	if Den == 0 {
		return 0, false
	}

	return Num / Den, true
}

// BasicTest ...
func BasicTest() {
	// Below describes the error handling pattern thats generally used in golang
	// The last returned value from a function is signified as a error object
	// In below example we sinply use a boolean value to indicate the successful
	// division. This can be changed to using a more standard error object
	output, success := Divide(10, 0)
	fmt.Println("Returned values:", output, " ", success)

	output, success = Divide(10, 2)
	fmt.Println("Returned values:", output, " ", success)
}

// DivideV2 ...
func DivideV2(Num int, Den int) (int, error) {
	if Den == 0 {
		return 0, errors.New("The denominator cant be 0")
	}

	return Num / Den, nil
}

// ErrorObjectTest ...
func ErrorObjectTest() {
	// Below describes the error handling pattern thats generally used in golang
	// The last returned value from a function is signified as a error object
	// In below example we sinply use a boolean value to indicate the successful
	// division. This can be changed to using a more standard error object
	output, errorObject := DivideV2(10, 0)
	fmt.Println("Returned values:", output, " ", errorObject)

	output, errorObject = DivideV2(10, 2)
	fmt.Println("Returned values:", output, " ", errorObject)

	// NOTE: Below is a general pattern used in golang for testing the errorCodes
	// returned by functions
	if _, err := DivideV2(19, 0); err != nil {
		fmt.Println("Error!")
	} else { // In golang, else block needs to start on the same line as the ending parentheses of the if block
		fmt.Println("Success!")
	}
}

// RandomizerError is a custom error type example
// Golang allows us to add more data to the custom errors
// How to create a custom error type?
//   1. Create a custom error struct
//   2. Tie the standard error interface by attaching the function `Error`
//      to the struct
type RandomizerError struct {
	Reason       string
	Distribution string
}

func (re *RandomizerError) Error() string {
	return "CustomError:" + re.Reason + " " + re.Distribution + "."
}

// Randomizer ...
func Randomizer() (string, error) {
	return "", &RandomizerError{"Test reason", "Uniform distribution"}
}

// CustomErrorTest ...
func CustomErrorTest() {
	output, err := Randomizer()
	fmt.Println("Output:", output, " ", "Error:", err)
}

func main() {
	BasicTest()
	ErrorObjectTest()
	CustomErrorTest()
}
