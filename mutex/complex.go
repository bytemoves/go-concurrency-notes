package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {

	var bankBalance int
	var balance sync.Mutex


	fmt.Printf("initial account balance : $%d.00",bankBalance)
	fmt.Println()


	incomes := []Income {
		{Source: "main job",Amount:500},
		{Source: "gifts",Amount:50},
		{Source: "part time",Amount:60},
		{Source: "investemt",Amount:100},
	}
	wg.Add(len(incomes))
	//loop through 52 weeks and print out how much is made keep a running total
	for i,income := range incomes {

		go func ( i int , income Income) {
			defer wg.Done()
			for week := 1 ; week <= 52; week++{
				balance.Lock()
				temp :=bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("on week %d you earned $%d.00 from %s\n",week,income.Amount,income.Source)
			}

		}(i,income)


	}

	wg.Wait()

	fmt.Printf("final bank balace: $%d.00 ",bankBalance)
	fmt.Println()

}