package np00

import (
	"fmt"
	"math"
)

type array []float64
type nparray []array

func MaxFloatInSlice(fls []float64) (m float64) {

	m = fls[len(fls)-1]
	for _, e := range fls {
		if m <= e {
			m = e
		}
	}
	return m
}

func SumOfSlice(fls []float64) (s float64) {
	for _, e := range fls {
		s += e
	}
	return s
}

func sumExpC(fls []float64) (s float64) {
	c := MaxFloatInSlice(fls)
	for _, e := range fls {
		s += math.Exp(e - c)
	}
	return s
}

func SoftMax(fls []float64) (sm []float64) {
	c := MaxFloatInSlice(fls)
	sum_exp_c := sumExpC(fls)
	sm = make([]float64, len(fls))

	for i, v := range fls {
		sm[i] = math.Exp(v-c) / sum_exp_c
	}
	return sm
}

func Makenparray(row, col int) nparray {
	npa := make([]array, row)
	for z := range npa {
		array := make([]float64, col)
		npa[z] = array
	}
	return npa
}

func (n nparray) Shape() [2]int {
	row := len(n)
	col := len(n[row-1])
	return [2]int{row, col}
}

func (m nparray) colsToarray(col int) (fa []float64) {
	fa = make([]float64, m.shape()[0])
	for r := range m {
		fa[r] = m[r][col]
	}
	return
}

func (a array) add(b array) (f float64) {

	if len(a) == len(b) {
		for idx := range a {
			f += a[idx] * b[idx]
		}
	}
	return f
}

func (n nparray) String() string {

	var str string = ""
	str += fmt.Sprintf("{\n")

	for r := range n {
		str += fmt.Sprintf("[ ")

		for k := range n[r] {
			str += fmt.Sprintf("%v ", n[r][k])
		}

		str += fmt.Sprintf("]\n")
	}
	str += fmt.Sprintf("}")
	return str
}

func add(n nparray, m nparray) nparray {

	npa := make([]array, n.shape()[0])
	for z := range npa {
		array := make([]float64, n.shape()[1])
		npa[z] = array
	}

	if n.shape()[0] == m.shape()[0] && n.shape()[1] == m.shape()[1] {
		for r := range n {
			for k := range n[r] {
				npa[r][k] = n[r][k] + m[r][k]
			}
		}
	} else {
		panic("shape check error")
	}
	return npa
}

func (n nparray) multi(f float64) nparray {

	npa := make([]array, n.shape()[0])
	for z := range npa {
		array := make([]float64, n.shape()[1])
		npa[z] = array
	}

	for r := range n {
		for k := range n[r] {
			npa[r][k] = n[r][k] * f
		}
	}
	return npa
}

/*
 [a b]  .  (c  = a*c + b*d
            d)
img: https://s3.amazonaws.com/nkvd/pub/matrixDot.png

*/

func dot(n nparray, m nparray) nparray {

	if n.shape()[1] != m.shape()[0] {
		panic("error row x col")
	}

	npa := make([]array, n.shape()[0])
	for z := range npa {
		array := make([]float64, m.shape()[1])
		npa[z] = array
	}
	for ncol := range n {
		for mrow := range m[0] {
			npa[ncol][mrow] = n[ncol].add(m.colsToarray(mrow))
		}
	}
	return npa
}
