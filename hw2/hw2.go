package main

import (
	"fmt"
)

func main() {
	s := []float64{2, 3, 5, 7, 11, 13}
	fmt.Print(Average(s))
}

func Average(s []float64) float64 {
	var sum float64 = 0
	for _, n := range s {
		sum = sum + n
	}
	return sum / float64(len(s))
}
