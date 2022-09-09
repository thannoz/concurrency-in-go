package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var balance sync.Mutex

type Income struct {
	Source string
	Amount float32
}

func main() {
	// variable for bank balance
	var bankBalance float32

	// print out starting values
	fmt.Printf("Initial account balance: %f\n", bankBalance)

	// define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))
	// loop throug 52 weeks and print out how much is made; keep a running total
	for x, income := range incomes {
		go func(x int, income Income) {
			defer wg.Done()

			for week := 0; week < 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d, you earned $%f from %s\n", week, income.Amount, income.Source)
			}
		}(x, income)
	}

	// print out final balance
	wg.Wait()
	fmt.Printf("Final balance: %f\n", bankBalance)

}
