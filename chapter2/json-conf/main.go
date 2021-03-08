package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//
type config struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	// Create Parser for RO
	decoder := json.NewDecoder(file)
	// Intialize instance empty conf
	conf := &config{}
	// Fill instance conf
	err := decoder.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("conf.Path: %s\n", conf.Path)
}
