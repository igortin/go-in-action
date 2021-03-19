package main

import (
	"log"
	"net"
)

/*
Send events to log service
nc -lk  1025
*/

func main() {
	// Create connection
	conn, err := net.Dial("tcp", "localhost:1025")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	logger := log.New(conn, "No: ", log.LstdFlags|log.Lshortfile)
	logger.Println("No alarm")
	// Not call log.Fatal because it exit immediatly
	// Utilize Log.Panic, it call defer func and correctly close connections
	logger.Panicln("Exception programm")
}
