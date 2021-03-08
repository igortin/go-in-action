package main

import (
	"log"
	"os"
	"strings"
)

func main() {
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

		err := compresFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Compressed successfully: %s\n", filename+".gz")
	}
}

func compresFile(filensme string) error {
	// Bussiness logic
	return nil
}
