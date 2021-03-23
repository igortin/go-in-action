package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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

	// Casting p and create instance m
	m := p.(map[string]interface{})

	// Casting to string
	// firstname := m["firstName"].(string)
	// fmt.Println(reflect.TypeOf(firstname)) // string
	// fmt.Println(firstname)

	// Casting interface{} to []interfaces{}
	educations := m["education"].([]interface{})

	// Take index 0
	education := educations[0] // []interface[0]

	// Casting interface[0] to map[string]interface{}
	edu := education.(map[string]interface{})

	// Casting interface{} to string
	s := edu["degree"].(string)
	fmt.Println(reflect.TypeOf(s)) // string
	fmt.Println(s)
}

// Output
// [map[degree:Bachelor of Science in Mathematics institution:Northwest Missouri State Teachers College] map[degree:Masters in English institution:University of Pennsylvania]]
// Bachelor of Science in Mathematics
