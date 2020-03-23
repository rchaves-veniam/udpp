package udpp

import "fmt"

// Println wrapper
func Println(args ...string) {
	fmt.Println(args)
}

// Print wrapper
func Print(args ...string) {
	fmt.Print(args)
}
