package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(c chan int) {
		fmt.Println("Goroutine inside: send 1 to ch and stop until Main goroutine recieved value from ch")

		c <- 1
		fmt.Println("Goroutine inside: Goroutine can proccess, because Main got value from ch")
	}(ch)
	for {
		select {

		case v := <-ch:
			fmt.Printf("Main inside: Get from ch %v\n", v)
		default:
			time.Sleep(1 * time.Second)
			fmt.Print("Main inside: wait.\n")
		}
	}
}

/* Output
Goroutine inside: send 1 to ch and stop until Main goroutine recieved value from ch
Main inside: wait.
Main inside: Get from ch 1
Goroutine inside: Goroutine can proccess, because Main got value from ch
*/
