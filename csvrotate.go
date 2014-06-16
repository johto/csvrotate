package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s input.csv\n", os.Args[0])
		os.Exit(1)
	}
	fh, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read file %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}
	defer fh.Close()

	r := csv.NewReader(fh)
	header, err := r.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read CSV header: %s\n", err)
		os.Exit(1)
	}
	maxLength := 0
	for _, datum := range header {
		if len(datum) > maxLength {
			maxLength = len(datum)
		}
	}
	padding := make([]string, len(header))
	for i, datum := range header {
		padding[i] = strings.Repeat(" ", maxLength - len(datum))
	}
	maxPadding := strings.Repeat(" ", maxLength)

	recordNo := 1
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "could not read from CSV file: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("-- RECORD %d --\n", recordNo)
		for i, _ := range header {
			first := true
			for _, line := range strings.Split(record[i], "\n") {
				if first {
					fmt.Printf("%s%s | %s\n", header[i], padding[i], line)
					first = false
				} else {
					fmt.Printf("%s | %s\n", maxPadding, line)
				}
			}
		}
		fmt.Printf("-- END OF RECORD %d --\n\n", recordNo)
		recordNo++
	}
}
