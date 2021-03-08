package main

import (
	"flag"
	"fmt"
)

// variables
var boolvar bool
var word string

func init() {
	flag.BoolVar(&boolvar, "boolvar", false, "flag-cli -yes")
	flag.BoolVar(&boolvar, "b", false, "flag-cli -yes")
	flag.StringVar(&word, "word", "", "flag-cli -word hi")
	flag.StringVar(&word, "w", "", "flag-cli -word hi")

}

func main() {
	flag.Parse()

	fmt.Printf("%s %s: %v\n", "flag", "boolvar", boolvar)

	if boolvar {
		if len(word) > 0 {
			fmt.Printf("%s: %v\n", "word", word)
		} else {
			fmt.Printf("%s: %v\n", "word", nil)
		}
	}
}
