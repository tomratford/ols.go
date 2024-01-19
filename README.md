# OLS

This is my attempts at implementing ordinary least squares in Go.

The usage is:

```console
$ ./ols input.csv y ~ x1 + x2 + x3
```

For a csv `input.csv` styled as:

```csv
y,x1,x2,x3
23,1.23,5.483,3.29
40,1.54,4.343,3.20
...
```

## Goals

* [X] Support OLS of continuous exploratory variables versus a continuous response variable
* [ ] Support input in the classic R formula style
* [ ] Factor variable support
* [ ] Support interactions between values
* [ ] Calcuate stddev for each variable
* [ ] Create appropriate diagnostics (R^2, sigma, etc...) for a LM
* [ ] Support formats other than CSV
