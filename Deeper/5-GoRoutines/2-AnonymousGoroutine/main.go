// Go program to illustrate how
// to create an anonymous Goroutine
package main

import (
	"fmt"
	"time"
)

// Main function
func main() {

	fmt.Println("Welcome!! to Main function")

	// Creating Anonymous Goroutine
	go func() {

		fmt.Println("Welcome!! to anonymous Goroutine")
	}()

	time.Sleep(1 * time.Second) // without this line, the program will end before the goroutine
	fmt.Println("GoodBye!! to Main function")
}
