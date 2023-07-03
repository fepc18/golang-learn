package main

import "fmt"

//Range is used to iterate over elements in a variety of data structures
func main() {

	a := []string{"Foo", "Bar"}
	for i, s := range a {
		fmt.Println(i, s)
	}
	//String iteration: runes or bytes
	//For strings, the range loop iterates over Unicode code points.
	for i, ch := range "日本語" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}

	// To loop over individual bytes, simply use a normal for loop and string indexing:
	const s = "日本語"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	// Map iteration: keys and values
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Channel iteration
	// For channels, the iteration values are the successive values sent on the channel until closed.

	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}

	//******Gotchas********
	//1. Unexpected values in range loop
	//For arrays and slices, the range loop generates two values:
	//first the index,
	//then the data at this position.
	//If you omit the second value, you get only the indices.
	//**BAD**
	primes := []int{2, 3, 5, 7}
	for p := range primes {
		fmt.Println(p)
	}
	//**GOOD**
	for _, p := range primes {
		fmt.Println(p)
	}

	//2. Can’t change entries in range loop
	//Why isn’t the slice updated in this example?
	//**BAD**
	slice1 := []int{1, 1, 1}
	for _, n := range slice1 {
		n += 1
	}
	fmt.Println(slice1)

	//The range loop copies the values from the slice to a local variable n; updating n will not affect the slice.
	//**GOOD**
	slice2 := []int{1, 1, 1}
	for i := range slice2 {
		slice2[i] += 1
	}
	fmt.Println(slice2)

}
