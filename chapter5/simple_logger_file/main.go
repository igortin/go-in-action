package main

import (
	"log"
	"os"
)

var filename = "prog.log"

func main() {
	var file *os.File
	var err error

	file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, os.Args[0]+": ", log.LstdFlags|log.Lshortfile)
	logger.Println("No alarm")
	logger.Fatalln("Exception programm")
}
