package main

import "fmt"

func main() {
	var name = "Juan perez" // Type declaration is optional here.
	fmt.Printf("Variable 'name' is of type %T\n", name)

	// Multiple variable declarations with inferred types
	var firstName, lastName, age, salary = "John", "Maxwell", 28, 50000.0

	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n",
		firstName, lastName, age, salary)

	// Short variable declaration syntax
	nameShort := "Carlos Gurierrez"
	age, salary, isProgrammer := 35, 50000.0, true

	fmt.Println(nameShort, age, salary, isProgrammer)
}
