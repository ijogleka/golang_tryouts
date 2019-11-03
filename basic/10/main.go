package main

import "fmt"

func main() {
	fmt.Println("This file simply shows different ways of defining varianles")

	// Integers
	var age1 int  // Declared, with a value of 0, and not used
	var age2 = 10 // Declared, with a value of 10, and not used
	age3 := 11    // Declared, with a value of 11, and not used
	fmt.Println("Integers:", age1, " ", age2, " ", age3)

	// Strings
	var name1 string
	var name2 = "Indrajeet"
	name3 := "Joglekar"
	fmt.Println("Strings:", name1, " ", name2, " ", name3)

	// Slices
	scores1 := []int{1, 3, 4}
	scores2 := make([]int, 10)    // Length of 10 and Capacity of 10
	scores3 := make([]int, 0, 10) // Length of 0 and Capacity of 10
	scores4 := scores1[1:2]       // Its just a reference to the scores1 array
	fmt.Println("Slices:", scores1, " ", scores2, " ", scores3, " ", scores4)

	// Maps
	var mapping1 map[string]int
	mapping2 := make(map[string]int, 10) // Create a map with size of 10 keys?
	mapping2["some_key"] = 1111          // Add elements to map
	mapping3 := map[string]int{          // Literal way of defining a map
		"other":   1,
		"another": 2, // Trailing comma is required
	}
	fmt.Println("Maps:", mapping1, " ", mapping2, " ", mapping3)
	for key, value := range mapping3 { // for loop is not sorted in any way
		fmt.Println("  Mapping3: key:", key, " value:", value)
	}

	type composite struct {
		City    string
		Zipcode int
	}
	instance1 := composite{}
	instance2 := composite{"NYC", 10022}
	instance3 := composite{City: "Jersey City", Zipcode: 07310}
	fmt.Println("Structs:", instance1, " ", instance2, " ", instance3)

	// Adds the values of the struct instances to the slice
	sliceOfStructs := []composite{instance1, instance2, instance3}
	fmt.Println("SliceOfStructs:", sliceOfStructs)

	fmt.Println("Thats about it")
}
