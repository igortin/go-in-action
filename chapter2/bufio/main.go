package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("./test")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	buffer := bufio.NewReader(file)

	for {
		line, _, err := buffer.ReadLine()

		if err != nil {
			if err == io.EOF {
				return
			}
			log.Println(err)
		}

		s := string(line)

		re, _ := regexp.Compile(`[1-7][1-7][1-7]+`)

		isTrue := re.MatchString(s)

		if isTrue {
			fmt.Println(s)
		}

	}

}
