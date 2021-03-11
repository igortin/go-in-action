package main

import (
	"log"
	"time"
)

// Close channel by sender - goroutine. Main goroutine form ch catch nil. Output: Event: _
func main() {
	timeout := time.After(1 * time.Second)

	ch := make(chan []byte)
	// Goroutine
	go externalVal(ch)
	for {
		select {
		case msg := <-ch:
			log.Printf("Event: %s", msg)
		case <-timeout:
			log.Println("Time Out.")
			return
		default:
			log.Println("Wait.")
		}
	}
}

func externalVal(ch chan []byte) {
	ch <- []byte("GOOD")
	close(ch)
	log.Println()
}
