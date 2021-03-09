package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	// WaitGroup Counter
	var wg sync.WaitGroup

	for _, filename := range os.Args[1:] {
		fileSlice := strings.Split(filename, `/`)
		file := fileSlice[len(fileSlice)-1]

		// Exclude files .gz or .bz2 from compress
		l := len(file)

		if l > 3 && file[len(file)-3:] == ".gz" {
			continue
		}

		if l > 4 && file[len(file)-4:] == ".bz2" {
			continue
		}

		// Increment counter in WaitGroup
		wg.Add(1)
		go func(file string) {
			err := compresFile(file)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Successfully compressed %v\n", file)
			// Decrement from WaitGroup counter
			wg.Done()

		}(filename)
	}

	// Wait until WaitGroup counter is 0 value
	wg.Wait()
}

func compresFile(filensme string) error {
	// Bussiness logic
	return nil
}
