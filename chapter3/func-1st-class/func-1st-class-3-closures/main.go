package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Функции первого класса. В Go функции можно присвоить переменным, передать функции другим функциям или даже написать функцию для возвращения функции.
Замыкание получение доступа к переменным вне области видимости
*/

// Создаем новый тип данных
type kelvin float64

// Создаем новый тип данных
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(300))
}

func realSensor() kelvin {
	return 0
}

/*
Функция которая вызывает другую функцию для фиксирования данных каждые 1 сек
На вход принимаеться параметр тип sensor - семантика: func() kelvin
*/

func measureTemperature(s sensor) {
	for i := 0; i < 2; i++ {
		value := s()
		fmt.Printf("%v K\n", value)
		time.Sleep(time.Second)
	}
}

/*
 Функция калибровка возращает функцию с исправленным расчетом отклонения, который передает измерительное устройство
*/
func calibrate(s sensor, ofset kelvin) sensor {
	// Return function with semantic - sensor
	return func() kelvin {
		// Замыкание: получем s и ofset значение из чужой области видимости
		origValue := s()

		fmt.Printf("origValue: %v, ofset: %v\n", origValue, ofset)

		return origValue + ofset
	}
}

func main() {
	/* Функция принимает параметр - функцию
	Возращает функцию в переменную newSensorFunc:
					func() kelvin {
						return s() + ofset
					}
	*/
	newSensorFunc := calibrate(fakeSensor, 10)

	// Вызов функции newSensorFunc() и получение исправленного значения
	// value := newSensorFunc()
	// fmt.Println(value)

	// Получаем калибровочный исправленные значения и записываем в консоль
	measureTemperature(newSensorFunc)
}
