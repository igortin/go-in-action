package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/*
		Отложенный перехват, восстановление не дает программе сразу свалиться и утечь ресурсам. Происходит устранение последствий
		в отлолженной функции с замыканием, закрыватся файл или канал - это и есть восстановление.
		Текст ошибки перехватывается и передается в main

	*/
	file, err := OpenFileSVC("data.svc")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
}

/*
Определение пременных в реализованно при определении функции
*/
func OpenFileSVC(filename string) (file *os.File, err error) {
	// Можно определить для использования в отложенном замыкании если нет variables в определнении функции
	// var file *os.File
	// var err error

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Trapped: file %v was closed correctly\n", file.Name())
			file.Close()
			err = r.(error)
		}
	}()

	// Try to open file
	file, err = os.Open(filename)
	if err != nil {
		return nil, err
	}

	// Raise crash
	removeLines(file)

	// Example unreachable line in case removeLines problems
	fmt.Println("unreachable line in code")
	return file, nil
}

func removeLines(f *os.File) {
	panic(errors.New("no lines"))
}
