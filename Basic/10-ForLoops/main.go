package main

import "fmt"

//examples for loops
func main() {
	//for loop traditional for Statement
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for loop with condition
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	//for loop with range
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	//for loop with range and ignore index
	for _, v := range s {
		fmt.Println(v)
	}

	//Infinite loop
	i := 5
	for {
		fmt.Println("Hello")
		if i == 10 {
			break
		}
		i++
	}

}
