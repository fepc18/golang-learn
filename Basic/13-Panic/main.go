package main

import "fmt"

// Examples of Defer and Panic
// Defer
// Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.

func main() {
	// Defer
	// Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
	// defer is often used where e.g. ensure and finally would be used in other languages.
	// defer is even more useful in the context of panic and recover.

	defer fmt.Println("world")
	defer first() // this will be executed first
	fmt.Println("hello")

	// Panic
	f()
	fmt.Println("Returned normally from f.")

}

func first() {
	fmt.Println("1st")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("-->function of panic %v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
