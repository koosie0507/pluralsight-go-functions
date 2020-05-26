package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	ans, _ := divide (1, 0)
	fmt.Printf("%f\n", ans)
	numbers := []float64{1, 3.14, 2, 6.22, 4.21}
	fmt.Printf("%f\n", sum(numbers...))
}

func sum(numbers ...float64) float64 {
	acc := float64(0)
	for _, v := range numbers {
		acc += v
	}
	return acc
}

func add(p1, p2 float64) float64 {
	return p1 + p2
}

func subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func multiply(p1, p2 float64) float64  {
	return p1 * p2
}

func divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by zero")
	}
	return p1 / p2, nil
}