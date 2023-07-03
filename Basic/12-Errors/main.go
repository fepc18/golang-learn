package main

import (
	"errors"
	"fmt"

	err "github.com/fepc18/golang-learn/Basic/12-Errors/errorsTypes"
)

// Constructing Errors
// the following function uses the errors package to return a new error with a static error message

func DoSomething() error {
	return errors.New("something didn't work")
}

var ErrDivideByZero = errors.New("divide by zero")

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {

	doSomethingErr := DoSomething()
	//try to do something with the error
	if doSomethingErr != nil {
		fmt.Println(doSomethingErr)
	}

	// type errors
	number := -1.0
	resultFloat, err := Sqrt(number)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultFloat)

	// Using errors.Is
	// Defining Expected Errors
	a, b := 10, 0
	result, err := Divide(a, b)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivideByZero):
			fmt.Println("divide by zero error")
		default:
			fmt.Printf("unexpected division error: %s\n", err)
		}
		return
	}

	fmt.Printf("%d / %d = %d\n", a, b, result)

}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		var a err.NegativeSqrtError
		a = err.NegativeSqrtError(f)
		return 0, a

		//return 0, errors.New("math: square root of negative number")
	}
	return 42, nil
	// implementation
}
