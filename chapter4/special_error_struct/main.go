package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(`./test.txt`)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	pe := GetChar(file)
	if pe.Message != "" {
		log.Println(pe.Message)
		log.Println(pe.Char)
	}
}

func GetChar(file *os.File) ParserErr {
	r := bufio.NewReader(file)
	i := 0
	e := ParserErr{
		Message: "",
		Char:    i,
	}

	for {
		b, err := r.ReadByte()
		if err != nil {
			e.Message = err.Error()
			e.Char = i
			return e
		}
		i++
		fmt.Println(string(b))
	}
}

type ParserErr struct {
	Message string
	Char    int
}

func (p *ParserErr) Error() string {
	msg := fmt.Sprintf("Message %s - at char index %v", p.Message, p.Char)
	return msg
}
