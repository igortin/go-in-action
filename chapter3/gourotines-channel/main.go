package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Create channel with timeout
	done := time.After(30 * time.Second)

	// Create channel for between goroutines
	echo := make(chan []byte)
	// Start goroutine to get stdin
	go readStdin(echo)

	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Succefully completed")
			// close channel
			close(echo)
			os.Exit(0)
		}
	}
}

// Bussines logic
func readStdin(out chan<- []byte) {
	for {
		// create slice of byte
		data := make([]byte, 1024)

		// from stdin write to slice
		l, _ := os.Stdin.Read(data)

		if l > 0 {
			// send []byte to channel
			out <- data
		}
	}
}
