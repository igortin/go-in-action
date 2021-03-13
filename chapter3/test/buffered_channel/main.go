package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)

	go func(c chan int) {
		fmt.Println("Goroutine inside: send 1 to ch")

		c <- 1
		fmt.Println("Goroutine inside: Goroutine not blocked")
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
Goroutine inside: send 1 to ch
Goroutine inside: Goroutine not blocked
Main inside: wait.
Main inside: Get from ch 1
*/
