package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Main programm: Outside goroutine")

	go func() {
		fmt.Println("Goroutine inside #1")
	}()

	fmt.Println("Main programm: Outside goroutine again")

	// Ask to GO scheduler to check Queue goroitines
	runtime.Gosched()
}
