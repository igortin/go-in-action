package main

import (
	"encoding/json"
	"fmt"
)

var JSON = `{
	"firstName": "Jean",
	"lastName": "Bartik",
	"age": 86,
	"education": [
		{
				"institution": "Northwest Missouri State Teachers College",
				"degree": "Bachelor of Science in Mathematics"
		},
		{
				"institution": "University of Pennsylvania",
				"degree": "Masters in English"
		}
	],
	"spouse": "William Bartik",
	"children": [
		"Timothy John Bartik",
		"Jane Helen Bartik",
		"Mary Ruth Bartik"
	]
}`

/*
 json in interface{} can not be accessed
 require converting interface{} to map[string]interface
*/

func main() {
	// Initialize variable type interface{}
	var p interface{}

	// Convert json string to byte array
	b := []byte(JSON)

	// Parse b and strorte to instance p
	err := json.Unmarshal(b, &p)
	if err != nil {
		panic(err)
	}
	PrintJson(p)
}

func PrintJson(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Printf("string %v\n", vv)
	case float64:
		fmt.Printf("float64 %v\n", vv)
	case []interface{}:
		fmt.Printf("array %v\n", vv)
		for _, u := range vv {
			// Recusrion
			PrintJson(u)
		}
	case map[string]interface{}:
		fmt.Printf("map %v\n", vv)
		for key, u := range vv {
			fmt.Printf("Key %v\n", key)
			// Recusrion
			PrintJson(u)
		}
	}
}
