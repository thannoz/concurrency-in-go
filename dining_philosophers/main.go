package main

import (
	"fmt"
	"sync"
	"time"
)

// Philosopher is a struct which stores information about a philosopher
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosopher = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger int = 3
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("----------------------------")
	fmt.Println("The table is empty")

	// start the meal
	dine()

	fmt.Println("The table is empty")

}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosopher))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosopher))

	// forks is a map of all 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosopher); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal by iterating through our slice of Philosophers.
	for i := 0; i < len(philosopher); i++ {
		// fire off a goroutine for the current philosopher
		go diningProblem(philosopher[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

}
