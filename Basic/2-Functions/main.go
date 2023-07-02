package main

import "fmt"

// Function accepting arguments
func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

// Function returning a value
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Function returning a value with named return value

func rectangle(l int, b int) (area int) {
	parameter := 2 * (l + b)
	fmt.Println("Parameter: ", parameter)

	area = l * b
	return // Return statement without specify variable name
}

// Function returning multiple values
func swap(x, y string) (string, string) {
	return y, x
}

// Passing Address to a Function
func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address, donde esta el valor
	*t = *t + " Doe" // defrencing pointer address
}

//Anonymous Functions in Golang
// Anonymous functions are the functions that don’t have any name.

func main() {
	// Function accepting arguments
	add(10, 20)
	// Function returning a value
	res := plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	// Return statement without specify variable name
	fmt.Println("Area: ", rectangle(10, 20))

	// Function returning multiple values
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)

	// Passing Address to a Function
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)
	update(&age, &text) // Passing address to a function
	fmt.Println("After :", text, age)

	//Anonymous Functions in Golang
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)

}
