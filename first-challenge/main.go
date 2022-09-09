package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, pluto!", &wg)
	wg.Wait()
	printMessage()

	// wg.Add(1)
	// go func(msg string) {
	// 	defer wg.Done()
	// 	fmt.Println(msg)

	// }("Hello, Pluto!")

	// wg.Add(1)
	// go func(msg string) {
	// 	defer wg.Done()
	// 	fmt.Println(msg)

	// }("Hello, Cosmos!")

	// wg.Add(1)
	// go func(msg string) {
	// 	defer wg.Done()
	// 	fmt.Println(msg)
	// }("Hello, Universe!")

	// wg.Wait()

}
