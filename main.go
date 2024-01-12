package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"ols/matrix"
	"os"
	"strconv"
)

func usage() {
	fmt.Print("usage: ols <input.csv> <response> ~ [exploratory]\n")
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	records, err := ReadFromCSV(os.Args[1])

	// Check if the error is that the file isn't real
	if os.IsNotExist(err) {
		usage()
		log.Fatalf("error: file '%s' does not exist\n", os.Args[1])
	} else if err != nil {
		usage()
		log.Fatal(err)
	}

	mod, err := OLS(records, "y", []string{"x1", "x2", "x3", "x4"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(mod.coef)
}

type model struct {
	dep string
	ind []string
	fitted, coef matrix.Matrix
}

func OLS(records [][]string, D string, Es []string) (model, error) {
	// For assume:
	//  - the first row is names
	names := records[0]
	Xs_ind := make(map[int]int, len(names))
	var y_ind int
	counter := 1
	for i, name := range names {
		for _, E := range Es {
			if name == E {
				Xs_ind[i] = counter
				counter += 1
			}
		}
		if name == D {
			y_ind = i
		}
	}

	n := len(records) - 1
	m := len(records[0]) - 1
	y := matrix.Zero(n, 1)
	X := matrix.Zero(n, m)
	for i, row := range records[1:] {
		X.Set(i, 0, 1) // intercept
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
				y.Set(i, 0, val)
			} else if key, ok := Xs_ind[j]; ok {
				X.Set(i, key, val)
			} else {
				continue
			}
		}
	}

	xT := matrix.Transpose(X)
	xTx, err := matrix.Multiply(xT, X)
	if err != nil {
		return model{}, err
	}
	xTx_inv, err := matrix.Inverse(xTx)
	if err != nil {
		return model{}, err
	}
	xTx_inv_xT, err := matrix.Multiply(xTx_inv, xT)
	if err != nil {
		return model{}, err
	}

	hat, err := matrix.Multiply(X, xTx_inv_xT)
	if err != nil {
		return model{}, err
	}

	fitted, err := matrix.Multiply(hat, y)
	if err != nil {
		return model{}, err
	}

	coef, err := matrix.Multiply(xTx_inv_xT, y)
	if err != nil {
		return model{}, err
	}

	mod := model{
		dep: D,
		ind: Es,
		fitted: fitted,
		coef:   coef,
	}

	return mod, nil
}

func ReadFromCSV(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return [][]string{}, err
	}

	r := csv.NewReader(file)
	return r.ReadAll()
}
