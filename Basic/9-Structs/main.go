package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20}) // You can initialize a struct with a struct literal without naming the fields, in the order they are defined in the struct

	fmt.Println(person{name: "Alice", age: 30}) // You can name the fields when initializing a struct

	fmt.Println(person{name: "Fred"}) // Omitted fields will be zero-valued

	fmt.Println(&person{name: "Ann", age: 40}) // Structs are passed by reference and can be created with & operator, used when you want to share a struct between different parts of your program

	fmt.Println(newPerson("Jon")) // Structs are mutable and can be created with a function

	s := person{name: "Sean", age: 50} // Structs are passed by value
	fmt.Println(s.name)

	sp := &s // Structs are passed by reference
	fmt.Println(sp.age)

	sp.age = 51 // Structs are mutable
	fmt.Println(s.age)
	fmt.Println(sp.age)

	dog := struct { // Anonymous struct
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
