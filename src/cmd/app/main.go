package main

import (
	"fmt"
	pshttp "http"
	"math"
	"net/http"
	"semver"
	"simplemath"
	"strings"
)

type MathExpression = string
type MathFunc = func(float64, float64) float64

const (
	Add MathExpression = "+"
	Sub MathExpression = "-"
	Mul MathExpression = "*"
	Div MathExpression = "/"
)

func mathExpression(expr MathExpression) MathFunc {
	switch expr {
	case Add:
		return simplemath.Add
	case Mul:
		return simplemath.Multiply
	case Sub:
		return simplemath.Subtract
	case Div:
		return func(a, b float64) float64 {
			answer, _ := simplemath.Divide(a, b)
			return answer
		}
	default:
		return func(a, b float64) float64 {
			return 0
		}
	}
}

func double(f1, f2 float64, f MathFunc) float64 {
	return 2 * f(f1, f2)
}

func squares() func() int64 {
	x := 0.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))
	}
}

func main() {
	ans, _ := simplemath.Divide(1, 0)
	fmt.Printf("%f\n", ans)
	numbers := []float64{1, 3.14, 2, 6.22, 4.21}
	fmt.Printf("%f\n", simplemath.Sum(numbers...))

	sv := semver.NewSemanticVersion(1, 0, 0)
	sv.IncMajor()
	fmt.Println(sv.String())

	var tripper http.RoundTripper = &pshttp.RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)

	a := func(name string) string {
		return fmt.Sprintf("anon func %s", name)
	}
	fmt.Println(a("1a"))

	f := mathExpression(Sub)
	fmt.Println(double(1, 2, f))

	p2 := squares()
	for n := 1; n <= 5; n++ {
		fmt.Print(p2(), " ")
	}
	fmt.Println()
}
