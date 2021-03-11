package main

import (
	"log"
	"time"
)

// Reciever (main goroutine) send signal to Sender (externalVal goroutine) about stop process.
// Pattern: Наблюдатель
func main() {
	// Control channel for close `ch` channel on sender side
	control := make(chan bool)

	// Timer channel
	timer := time.After(5 * time.Second)

	// Bussiness data channel for interaction between goroutines
	ch := make(chan []byte)

	// Start goroutine
	go externalVal(ch, control)

	/*
		Get values for channels
	*/
	for {
		select {
		case msg := <-ch:
			// Get bussine value
			log.Printf("Event: %s", msg)
		case <-timer:
			// Get time.Time object from timer channel
			// Send stop signal to control channel for sender
			control <- true
			log.Println("Timer stop programm")
			return
		}
	}
}

// func in goroutine
func externalVal(ch chan []byte, control chan bool) {
	for {
		select {
		case <-control:
			// Get stop signal from control channel and close ch channel
			close(ch)
			log.Println("Channel ch: closed")
			return
		default:
			// Send []byte to `ch` channel
			ch <- []byte("GOOD")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
