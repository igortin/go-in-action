package main

import (
	"log"
	"time"
)

// Close channel by receiver - main goroutine. Raise exception: send on closed channel. Reason: 21 line.
func main() {
	until := time.After(5 * time.Second)
	ch := make(chan []byte)
	go externalVal(ch)
	for {
		select {
		case msg := <-ch:
			log.Printf("Event: %s", msg)
		case <-until:
			close(ch)
			time.Sleep(10 * time.Second)
			return
		}
	}
}

func externalVal(ch chan []byte) {
	for {
		ch <- []byte("GOOD")
		time.Sleep(1 * time.Second)
	}
}
