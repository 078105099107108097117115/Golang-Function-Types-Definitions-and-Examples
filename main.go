package main

import (
	"fmt"
	"math"
)

// Function types and values

type binOp func(a, b int) string

func binOpConsumer(fn binOp, a int, b int) {
	res := fmt.Sprintf("%v -> %s", fn, fn(a, b))
	fmt.Println(res)
}

//Operator is a function type taking two float64 parameters
type Operator func(x, y float64) float64

//Map function definition using funtion types
func Map(op Operator, someCollection []float64, n float64) []float64 {
	res := make([]float64, len(someCollection))
	for ix, val := range someCollection {
		res[ix] = op(val, n)
	}
	return res
}

func main() {
	var add binOp = func(a, b int) string {
		return fmt.Sprintf("%d + %d = %d\n", a, b, a+b)
	}
	var multiply binOp = func(a, b int) string {
		return fmt.Sprintf(" %d x %d = %d", a, b, a*b)
	}
	binOpConsumer(add, 12, 13)
	binOpConsumer(multiply, 12, 13)

	a := []float64{12, 190, 31, 3, 2, 1}
	op := math.Pow
	out := Map(op, a, 2)
	fmt.Println(out)

}
