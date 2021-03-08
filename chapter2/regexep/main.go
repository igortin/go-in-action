package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Regexp object that will be used for mathing string
	regexpObject, _ := regexp.Compile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Printf("Regexp Object: %v\n", regexpObject)

	// Test string for match
	fmt.Println(regexpObject.MatchString("eve[7]"))
	fmt.Println(regexpObject.MatchString("Job[48]"))
	fmt.Println(regexpObject.MatchString("snakey"))
}
