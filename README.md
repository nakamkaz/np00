# np00
a library of purpose for matrix calc

* Install

```
go get github.com/nakamkaz/np00
```

* Update 

```
go get -u github.com/nakamkaz/np00
```

* Usage 

```go
package main

import (
	"fmt"
  "github.com/nakamkaz/np00"
)

func main() {

	a := []float64{1010, 1000, 990}

	fmt.Println("Max ", np00.MaxFloatInSlice(a))
	fmt.Println("Sum of a: ", np00.Sum(a))
	fmt.Printf("SoftMax: %v\n", np00.SoftMax(a))
	fmt.Printf("Sum of SoftMax: %v\n", np00.Sum(np00.SoftMax(a)))

	n := np00.NParray{
		[]float64{1, 2, 3},
		[]float64{7, 8, 9},
	}
	m := np00.NParray{
		[]float64{0.1, 0.3},
		[]float64{11, 13},
		[]float64{0.7, 17},
	}
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(m.Shape())
	fmt.Println(np00.Dot(n, m))

}
```
```
Max  1010
Sum of a:  3000
SoftMax: [0.999954600070331 4.539786860886666e-05 2.061060046209062e-09]
Sum of SoftMax: 1
{
[ 1 2 3 ]
[ 7 8 9 ]
}
{
[ 0.1 0.3 ]
[ 11 13 ]
[ 0.7 17 ]
}
[3 2]
{
[ 24.200000000000003 77.3 ]
[ 95 259.1 ]
}
```

* Usage 2

https://gist.github.com/nakamkaz/8cd9ff25b7fec5c029a45f6bc4c58e3a
