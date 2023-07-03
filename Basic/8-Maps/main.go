package main

import "fmt"

func main() {
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println(employee)

	// Map Length
	fmt.Println(len(employee))

	// Accessing Items
	fmt.Println(employee["Mark"])

	// Adding Items
	employee["John"] = 30
	fmt.Println(employee)

	// Removing Items
	delete(employee, "Mark")
	fmt.Println(employee)

	// Updating Items
	employee["Sandy"] = 40
	fmt.Println(employee)

	// Iterating Over Items
	for key, value := range employee {
		fmt.Println(key, value)
	}

	// truncating a map
	employee = make(map[string]int)
	fmt.Println(employee)

	// sorting a map
	// Must used a slice to sort a map
	// Example:
	// var keys []string
	// for key := range employee {
	// 	keys = append(keys, key)
	// }
	// sort.Strings(keys)

	// Merge two maps
	// Example:
	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}

	for k, v := range second {
		first[k] = v
	}

	fmt.Println(first)

}
