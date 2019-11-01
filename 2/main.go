package main

import "fmt"

// TeamWarranty ...
func TeamWarranty() (string, string) {
	return "WARRANTY", "D2C"
}
func main() {
	mnemonic, group := TeamWarranty()

	fmt.Println("Mnemonic:", mnemonic)
	fmt.Println("Group:", group)
}
