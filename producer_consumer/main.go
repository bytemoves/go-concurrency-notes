package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NUmberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func ( p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch 
	return <- ch
}
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NUmberOfPizzas{
		delay := rand.Intn(5) + 1

		fmt.Printf("Recieved order number #%d!\n",pizzaNumber)
		rnd := rand.Intn(12) + 1

		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("making pizza #%d. It will take %d seconds...\n",pizzaNumber,delay)
		//delay

		time.Sleep(time.Duration(delay) *time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("** We ran out of ingridient for pizza #%d!",pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("** The cook quit while making pizza #%d!",pizzaNumber)

		} else {
			success = true
			msg = fmt.Sprintf("pizza order #%d is ready",pizzaNumber)
		}
		p := PizzaOrder {
			pizzaNumber: pizzaNumber,
			message: msg,
			success: success,
		}

		return &p
		

	} 

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}
func pizzeria ( pizzaMaker *Producer) {
	//keep track of pizza we are makingi
	var i = 0

	// run until we recieve a quit noti

	//try to make pizzas 

	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber 
			select {
				// tried to make a pizza( se sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:

				//close channel
				close(pizzaMaker.data)
				close(quitChan)

				return

			}
		}
	}

}

func main() {
	//seed rnadom generator
	rand.Seed(time.Now().UnixNano())

	//print out a message
	color.Cyan("The pizzeria is open for business")
	color.Cyan("--------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

/// run producer in background
	go pizzeria(pizzaJob)
//create and run a consumer 
	for i := range  pizzaJob.data {
		if i.pizzaNumber <= NUmberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("order #%d is out for delivery!",i.pizzaNumber)
		
			} else {
				color.Red(i.message)
				color.Red("the customer is really mad!")
			}

		} else {
			color.Cyan("Done making pizzas....")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("***error closing channel!",err)
			}
		}
	}

	//ending message
	color.Cyan("------------------------")
	color.Cyan("Done for the day")
	color.Cyan("We made %d pizzas but failed to make %d , with a%d attempts in total.", pizzasMade,pizzasFailed,total)

	switch {
	case pizzasFailed > 9:
		color.Red("it was an awful day")
	case pizzasFailed >= 6 :
		color.Red("it was not a very good day")

	case pizzasFailed >=4:
		color.Yellow("it was an okay day")


	case pizzasFailed >=2:
		color.Yellow("it was an pretty good day")

	default:
		color.Green("it was a great day")


	}
}
