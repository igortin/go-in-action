package main

import (
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdout, os.Stdin)
	time.Sleep(30 * time.Second)
}

func echo(r io.Reader, w io.Writer) {
	io.Copy(w, r)
}
