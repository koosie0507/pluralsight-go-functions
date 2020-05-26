package simplemath

import (
	"errors"
	"math"
)

func Sum(numbers ...float64) float64 {
	acc := float64(0)
	for _, v := range numbers {
		acc += v
	}
	return acc
}

func Add(p1, p2 float64) float64 {
	return p1 + p2
}

func Subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func Multiply(p1, p2 float64) float64  {
	return p1 * p2
}

func Divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by zero")
	}
	return p1 / p2, nil
}
