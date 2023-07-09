package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	// loop will continue until the channel is closed with for
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received with loop", v, ok)
	}

	// loop will continue until the channel is closed with range
	ch2 := make(chan int)
	go producer(ch2)
	for v := range ch2 {
		fmt.Println("Received with Range ", v)
	}
}
