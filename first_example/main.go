package main

import (
	"fmt"
	"sync"
	
)


func printSomething ( s string,wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}


func  main()  {

		var wg sync.WaitGroup
//go tells the compiler to execute it in its own  go routine

	words := [] string {
		"alpha",
		"beta",
		"gamma",
		"phi",
		"zeta",
		"theta",
		"delta",
		"epsilon",

	}
	wg.Add(len(words))
	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s",i,x),&wg)
	}
	wg.Wait()
	//you use wait groups
	wg.Add(1)

	printSomething("this the second thing",&wg)

}
