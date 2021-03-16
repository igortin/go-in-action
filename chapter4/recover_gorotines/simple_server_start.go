package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

/*
Problem: in case goroutine crash than main programm crash
Solution: To goroutine code append block recover to trap crash and programm should be ok
*/

func main() {
	listen()
}

func listen() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Caught a panic and recovered\n")
		}
	}()

	listener, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println("Failed to open port on 1026")
		return
	}
	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}

		// Start goroutine and it created separated stack of functions
		go handle(conn)
	}
}

/*
Each goroutine create separated stack of functions.
Crash can not be transfered from one stack to another stack.
Thats why recover block in main will not fix a problem.
Thats why recover block appended to block of subprogram.
*/
func handle(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Caught a panic and recovered\n")
		}
		conn.Close()
	}()

	reader := bufio.NewReader(conn)

	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket.")
		conn.Close()
	}
	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	defer func() {
		conn.Close()
	}()
	conn.Write(data)
	panic(errors.New("Failure in response!"))
}
