package main

import (
	"fmt"
	"testing"
)

func main() {
	s := []float64{2, 3, 5, 7, 11, 13}
	fmt.Print(Average(s))
}

func Average(s []float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(s); i++ {
		sum = sum + s[i]
	}
	return sum / float64(len(s))
}

type testpair struct {
	values   []float64
	expected float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		result := Average(pair.values)
		if result != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"result", result,
			)
		}
	}
}
