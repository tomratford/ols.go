package main

import (
	"encoding/csv"
	"fmt"
	"ols/matrix"
	"os"
	"strconv"
	"strings"
)

func usage() {
	fmt.Printf("usage: ols <input.csv> <response> ~ [exploratory]\n")
}

type model struct {
	fitted, coef matrix.Matrix
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
		os.Exit(1)
	} else if err != nil {
		usage()
		fmt.Println(err)
		os.Exit(1)
	}

	// For now assume:
	//  - the first row is names
	//      all names beginning with X are X values, single y column
	names := records[0]
	Xs_ind := make(map[int]int, len(names))
	var y_ind int
	counter := 1
	for i, name := range names {
		if strings.HasPrefix(strings.ToLower(name), "x") {
			Xs_ind[i] = counter
			counter += 1
		}
		if strings.ToLower(name) == "y" {
			y_ind = i
		}
	}

	n := len(records) - 1
	m := len(records[0]) - 1
	y := matrix.Zero(n,1)
	X := matrix.Zero(n,m)
	for i, row := range records[1:] {
		X.Set(i,0,1) // intercept
		for j, v := range row {
			if v == "" {
				break
			}
			val, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Printf("error: parsing error, '%s' as a result of entry '%s'", err, v)
				os.Exit(1)
			}
			if j == y_ind {
				y.Set(i,0,val)
			} else if key, ok := Xs_ind[j]; ok {
				X.Set(i,key,val)
			} else {
				continue
			}
		}
	}

	xT := matrix.Transpose(X)
	xTx, err := matrix.Multiply(xT, X)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	xTx_inv, err := matrix.Inverse(xTx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	xTx_inv_xT, err := matrix.Multiply(xTx_inv, xT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hat, err := matrix.Multiply(X, xTx_inv_xT)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fitted, err := matrix.Multiply(hat, y)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	coef, err := matrix.Multiply(xTx_inv_xT, y)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mod := model{
		fitted: fitted,
		coef: coef,
	}

	fmt.Println(mod.coef)
}

func ReadFromCSV(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return [][]string{}, err
	}

	r := csv.NewReader(file)
	return r.ReadAll()
}
