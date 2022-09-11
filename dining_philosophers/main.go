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

var orderMutex sync.Mutex  // a mutex for the slice orderFinished
var orderFinished []string // the order in which philosophers finish dining and leave

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("----------------------------")
	fmt.Println("The table is empty")

	time.Sleep(sleepTime)

	// start the meal
	dine()

	fmt.Println("The table is empty")

}

func dine() {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

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

	// Wait for the philosopher to finish. This blocks until the wait group counter is 0.
	wg.Wait()
}

// diningProblem is the function fired off as a goroutine for each of our philosophers. It takes one
// philosopher, waitgroup to determine when everyone is done, a map containing the mutexes for every
// fork on the table, and waitgroup used to pause execution of every instance of this goroutine
// until everyone is seated at the table.
func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table\n", philosopher.name)
	seated.Done()

	seated.Wait()

	// eat three time
	for i := 0; i < hunger; i++ {
		// get a lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(eatTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the fork\n", philosopher.name)

	}

	fmt.Println(philosopher.name, "is satisified")
	fmt.Println(philosopher.name, "left the table")

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()

}
