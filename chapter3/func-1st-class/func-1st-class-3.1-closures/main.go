package main

import (
	"fmt"
)

type kelvin float64

func main() {
	var k kelvin = 294.0
	// Closures: get acces to varibale k
	sensor := func() kelvin {
		return k
	}

	fmt.Println(sensor()) // Выводит: 294

	k++
	fmt.Println(sensor()) // Выводит: 295

}
