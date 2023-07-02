package main

import (
	"fmt"
	"time"
)

func main() {

	// If statements are used to execute code based on a condition.
	x := 10
	y := 20
	if x < y {
		fmt.Println("x is less than y")
	} else if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("x is equal to y")
	}

	//Golang - if statement initialization
	// If statements can also be used to initialize variables.

	if x := 10; x < y {
		fmt.Println("x is less than y")
	}

	// Golang - switch statement
	// Switch statements are used to execute different code based on different conditions.
	today := time.Now()
	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
	// Golang - switch multiple cases Statement

	var dayOfWeek = "Monday"
	switch dayOfWeek {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Golang.")
	default:
		fmt.Println("It's weekend, time to rest!")
	}

}
