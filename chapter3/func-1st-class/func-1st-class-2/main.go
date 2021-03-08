package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Функции первого класса. В Go функции можно присвоить переменным, передать функции другим функциям или даже написать функцию для возвращения функции.
*/

type kelvin float64

// test Sensor
func fakeSensor() kelvin {
	return kelvin(rand.Intn(300))
}

func realSensor() kelvin {
	return 0
}

/*
Функция которая вызывает другую функцию для фиксирования данных каждые 1 сек
В функциию можно передать realSensor или fakeSensor
*/
func measureTemperature(sensor func() kelvin) {
	for i := 0; i < 2; i++ {
		value := sensor()
		fmt.Printf("%v K\n", value)
		time.Sleep(time.Second)
	}
}

func main() {
	// Функция принимает параметр - функцию
	measureTemperature(fakeSensor)
}
