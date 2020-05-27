package main

import (
	"errors"
	"fmt"
	pshttp "http"
	"io"
	"math"
	"net/http"
	"readers"
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

func readSomething(reader io.ReadCloser) error {
	value, err := reader.Read([]byte("test"))
	defer func() {_ = reader.Close()}()
	if err == io.EOF {
		fmt.Println("reached end of file")
		return err
	}
	if err != nil {
		fmt.Printf("an error occured: %s\n", err.Error())
		return err
	}
	fmt.Println(value)
	return nil
}

func readAll(reader io.ReadCloser) (err error) {
	defer func() {
		_ = reader.Close()
		if p := recover(); p != nil {
			fmt.Println(p)
			err = errors.New("it's just a mild case of anxiety, nothing to worry about")
		}
	}()
	for {
		value, readErr := reader.Read([]byte("test"))
		if readErr == io.EOF {
			break
		} else if readErr != nil {
			err = readErr
			return
		}

		fmt.Println(value)
	}
	return nil
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

	var funcs []func() int64
	for i:=1; i<=5; i++ {
		closedI := i
		funcs = append(funcs, func() int64 {
			return int64(math.Pow(float64(closedI), 2))
		})
	}

	for _, sq := range funcs {
		fmt.Print(sq(), " ")
	}
	fmt.Println()
	fmt.Println()

	readSomething(readers.NewBadReader("some error"))
	err := readAll(readers.NewSimpleReader(6))
	if err != nil {
		fmt.Printf("Something bad happened: %s", err)
	}
}
