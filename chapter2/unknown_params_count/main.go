package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	//args := os.Args[1:]

	// newStr, err := Concat(args)

	// Unkwnow parameters length
	newStr, err := Concat("a", "b", "c", "v", "w")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newStr)
}

//Concat paramters converts to []string
func Concat(args ...string) (string, error) {
	// join []string items to string
	if len(args) == 0 {
		return "", errors.New("no parameters")
	}
	s := strings.Join(args, " ")
	return s, nil
}
