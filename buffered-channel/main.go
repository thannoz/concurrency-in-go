package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Printf("Got %d from channel\n", i)

		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Printf("sending %d to channel...\n", i)
		ch <- i
		fmt.Printf("sent %d to channel!\n", i)
	}

	fmt.Println("Done!")
	close(ch)
}
