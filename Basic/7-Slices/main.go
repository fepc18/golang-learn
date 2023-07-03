package main

import "fmt"

//Slice examples
func main() {
	// Slices
	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// a[low : high]
	// Example:
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4] // Note the explicit address-of operator
	// array is passed implicitly as an address, not a copy
	fmt.Println(s)
	// Example:
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// Slice literals
	// A slice literal is like an array literal without the length.
	// Example:
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	// Slice defaults
	// When slicing, you may omit the high or low bounds to use their defaults instead.
	// The default is zero for the low bound and the length of the slice for the high bound.
	// For the array
	// var a [10]int
	// these slice expressions are equivalent:
	// a[0:10]
	// a[:10]
	// a[0:]
	// a[:]
	// Example:
	r := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(r)
	r = r[1:4]
	fmt.Println(r)
	r = r[:2]
	fmt.Println(r)
	r = r[1:]
	fmt.Println(r)
	// Slice length and capacity
	// A slice has both a length and a capacity.
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
	// Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.
	// Example:
	t := []int{2, 3, 5, 7, 11, 13}
	printSlice(t)
	// Slice the slice to give it zero length.
	t = t[:0]
	printSlice(t)
	// Extend its length.
	t = t[:4]
	printSlice(t)
	// Drop its first two values.
	t = t[2:]
	printSlice(t)
	// Nil slices
	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array.
	// Example:
	var u []int
	fmt.Println(u, len(u), cap(u))
	if u == nil {
		fmt.Println("nil!")
	}
	// Creating a slice with make
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	// a := make([]int, 5)  // len(a)=5
	// To specify a capacity, pass a third argument to make:
	// b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	// b = b[:cap(b)] // len(b)=5, cap(b)=5
	// b = b[1:]      // len(b)=4, cap(b)=4
	// Example:
	v := make([]int, 5)
	printSlice2("v", v)
	w := make([]int, 0, 5)
	printSlice2("w", w)
	x := w[:2]
	printSlice2("x", x)
	y := x[2:5]
	printSlice2("y", y)
	// Slices of slices
	// Slices can contain any type, including other slices.
	// Example:
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"}}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"

	board[1][2] = "X"
	board[1][0] = "O"

	board[0][2] = "X"
	printSlice3(board)
	// Appending to a slice
	// It is common to append new elements to a slice, and so Go provides a built-in append function.
	// The documentation of the built-in package describes append.
	// func append(s []T, vs ...T) []T
	// The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
	// The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
	// If the backing array of s is too small to fit all the given values a bigger array will be allocated.
	// The returned slice will point to the newly allocated array.
	// Example:
	var z []int
	printSlice2("z", z)
	// append works on nil slices.
	z = append(z, 0)
	printSlice2("z", z)
	// The slice grows as needed.
	z = append(z, 1)
	printSlice2("z", z)
	// We can add more than one element at a time.
	z = append(z, 2, 3, 4)
	printSlice2("z", z)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func printSlice3(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", s[i])
	}
}
