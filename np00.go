package np00

import (
	"fmt"
	"math"
)

type Array []float64
type NParray []Array

func MaxFloatInSlice(fls []float64) (m float64) {

	m = fls[len(fls)-1]
	for _, e := range fls {
		if m <= e {
			m = e
		}
	}
	return m
}

func Sum(fls []float64) (s float64) {
	return SumOfSlice(fls)
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

func MakeNParray(row, col int) NParray {
	npa := make([]Array, row)
	for z := range npa {
		Array := make([]float64, col)
		npa[z] = Array
	}
	return npa
}

func (n NParray) Shape() [2]int {
	row := len(n)
	col := len(n[row-1])
	return [2]int{row, col}
}

func (m NParray) colsToArray(col int) (fa []float64) {
	fa = make([]float64, m.Shape()[0])
	for r := range m {
		fa[r] = m[r][col]
	}
	return
}

func (a Array) add(b Array) (f float64) {

	if len(a) == len(b) {
		for idx := range a {
			f += a[idx] * b[idx]
		}
	}
	return f
}

func (n NParray) String() string {

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

func add(n NParray, m NParray) NParray {

	npa := make([]Array, n.Shape()[0])
	for z := range npa {
		Array := make([]float64, n.Shape()[1])
		npa[z] = Array
	}

	if n.Shape()[0] == m.Shape()[0] && n.Shape()[1] == m.Shape()[1] {
		for r := range n {
			for k := range n[r] {
				npa[r][k] = n[r][k] + m[r][k]
			}
		}
	} else {
		panic("Shape check error")
	}
	return npa
}

func (n NParray) multi(f float64) NParray {

	npa := make([]Array, n.Shape()[0])
	for z := range npa {
		Array := make([]float64, n.Shape()[1])
		npa[z] = Array
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

func Dot(n NParray, m NParray) NParray {

	if n.Shape()[1] != m.Shape()[0] {
		panic("error row x col")
	}

	npa := make([]Array, n.Shape()[0])
	for z := range npa {
		Array := make([]float64, m.Shape()[1])
		npa[z] = Array
	}
	for ncol := range n {
		for mrow := range m[0] {
			npa[ncol][mrow] = n[ncol].add(m.colsToArray(mrow))
		}
	}
	return npa
}
