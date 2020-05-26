package main

import (
	"fmt"
	pshttp "http"
	"net/http"
	"semver"
	"simplemath"
	"strings"
)

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
}

