package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	// create a channel unidirectional
	sendch := make(chan<- int)
	go sendData(sendch)
	//  fmt.Println(<-sendch)
	//invalid operation: cannot receive from
	//send-only channel sendch (variable of type chan<- int)

	//Create a channel bidirectional
	ch := make(chan int)
	go sendData(ch)
	fmt.Println(<-ch)

}
