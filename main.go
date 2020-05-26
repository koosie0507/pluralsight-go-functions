package main

import (
	"./simplemath"
	"fmt"
)

func main() {
	ans, _ := simplemath.Divide (1, 0)
	fmt.Printf("%f\n", ans)
	numbers := []float64{1, 3.14, 2, 6.22, 4.21}
	fmt.Printf("%f\n", simplemath.Sum(numbers...))
}

