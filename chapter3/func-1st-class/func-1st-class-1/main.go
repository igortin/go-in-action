package main

import (
	"fmt"
	"math/rand"
)

/*
Функции первого класса. В Go функции можно присвоить переменным, передать функции другим функциям
или даже написать функции для возвращения функций.
*/

// Example: Create programm to get Kelvin temperature
type kelvin float64

// test Sensor
func fakeSensor() kelvin {
	return kelvin(rand.Intn(300))
}

func realSensor() kelvin {
	return 0
}

func main() {
	// Объявление переменой типа: func() kelvin - семантика
	var sensor func() kelvin
	//  Присваивает функции переменным
	sensor = fakeSensor
	// Получение значений
	fmt.Println(sensor())
	sensor = realSensor
	fmt.Println(sensor())
}
