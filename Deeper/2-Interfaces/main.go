package main

import (
	"fmt"
	"math"

	"github.com/fepc18/golang-learn/Deeper/2-Interfaces/interfaces"
)

// examples for used Sequence

func main() {

	var abser interfaces.Abser
	//var f MyFloat = -math.Sqrt2
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs()) // Implements Abser interface (method Abs()) for MyFloat type

	abser = f // a MyFloat implements Abser
	fmt.Println(abser.Abs())

	fmt.Println(f.Double()) // Implements Multi interface (method Double()) for MyFloat type

	if f, ok := abser.(MyFloat); ok { // a MyFloat implements Abser
		fmt.Println(f.Double())
	}

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) Double() float64 {
	return float64(f) * 2
}
