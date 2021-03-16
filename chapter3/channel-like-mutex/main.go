package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Create buffered channel
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt)

	chLocker := make(chan bool, 1)

	for i := 0; i < 5; i++ {

		/*
			id 0 send value to chan chLocker - true and locked it (like Mutex)
			id 0 succefully start bussines logic
			id 0 succefully do step by step bussiness logic
			id 1 send value to chan chLocker but hang until id 0 will release lock (buffer fulled)
			id 0 succesfully completed bussiness logic
			id 0 release lock
			id 1 succesfully started bussiness logic
		*/
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
