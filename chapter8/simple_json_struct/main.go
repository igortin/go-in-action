package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json: "name"`
	Age     int    `json: "age"`
	Counter int    `json: "counter"`
}

// Json doc
var JSON = `{
	"name" : "Beauty Angel",
	"age": 777,
	"counter": 7
}`

func main() {
	var p Person
	b := []byte(JSON)

	err := json.Unmarshal(b, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}
