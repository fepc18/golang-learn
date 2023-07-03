package main

import (
	"fmt"

	math "github.com/fepc18/golang-learn/Basic/3-Packages/Math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
