package main

import (
	"fmt"
	"time"
)

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Println("Type of a is %T", a)
	}

	b := make(chan int, 3)
	fmt.Println("Type of b is %T", b)

	c := make(chan string, 10)
	fmt.Println("Type of c is %T", c)

	//****Sending and receiving from a channel****

	c <- "Hello"
	fmt.Println(<-c) // Hello is received from the channel

	//*****Channel example with goroutine*****

	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	<-done // This line of code will block the main goroutine until it receives data from the channel
	fmt.Println("Main received data")
}

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
