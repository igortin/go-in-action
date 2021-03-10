package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	w := newWords()
	for _, f := range os.Args[1:] {
		// Wait Group counter increase
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			// Wait Group counter decrease
			wg.Done()
		}(f)
	}

	// Wait Group wait until counter = 0
	wg.Wait()

	fmt.Println("Words that appear more than once:")
	w.Lock()
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
	w.Unlock()
}

type words struct {
	sync.Mutex
	found map[string]int
}

func newWords() *words {
	return &words{
		found: map[string]int{},
	}
}

func (w *words) append(word string, n int) {
	// Lock map - w
	w.Lock()
	// Unlock map - w
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())

		dict.append(word, 1)
	}

	return scanner.Err()
}
