package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("usage: ols <input.csv> <response> ~ [exploratory]\n")
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	records, err := ReadFromCSV(os.Args[1])

	// Check if the error is that the file isn't real
	if os.IsNotExist(err) {
		usage()
		fmt.Printf("error: file '%s' does not exist\n", os.Args[1])
	} else if err != nil {
		usage()
		fmt.Println(err)
	}

	fmt.Println(records)
}

func ReadFromCSV(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return [][]string{}, err
	}

	r := csv.NewReader(file)
	return r.ReadAll()
}

/*
   Essentially I now need a matrix library.
   Because however I swing it I'm gonna be inverting some kinda matrix
*/
