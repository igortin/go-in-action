package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt)

	chLocker := make(chan bool, 1)

	for i := 0; i < 2; i++ {
		// 7 x Goroutines
		go func(id int, c chan bool) {
			fmt.Printf("id #%v: wants Lock\n", id)
			c <- true
			fmt.Printf("id #%v: succesffully has Lock\n", id)
			// Bussness logic
			time.Sleep(1 * time.Second)
			fmt.Printf("id #%v: succesffully released Lock\n", id)
			<-c
		}(i, chLocker)
	}
	<-done
}
