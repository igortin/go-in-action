package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

func newHandler() *handler {
	return &handler{}
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	// name is not exist

	if name == "" {
		name = "no existing"
	}

	fmt.Fprintln(w, "name is", name)
}

func main() {
	handler := newHandler()

	ch := make(chan os.Signal)
	// Send signals to channels
	signal.Notify(ch, os.Interrupt, os.Kill)

	// Goroutine for monitoring os signals
	go func(ch <-chan os.Signal) {
		<-ch
		// Gracefull stop server ListenAndServe after all connection served
		manners.Close()
	}(ch)

	port := os.Getenv("http_port")
	fmt.Printf("Srever HTTP startrting on %s", port)

	manners.ListenAndServe(":8080", handler)
}
