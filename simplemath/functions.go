package simplemath

import (
	"errors"
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

func Divide(p1, p2 float64) (answer float64, err error) {
	if p2 == 0 {
		err = errors.New("cannot divide by zero")
	} else {
		answer = p1 / p2
	}
	return
}
