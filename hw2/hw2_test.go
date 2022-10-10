package main_test

import (
	"github.com/stretchr/testify/assert"
	M "gitlab.ozon.dev/qa/classroom-4/act-device-api/hw2"
	"testing"
)

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
		assert.Equal(t, M.Average(pair.values), pair.expected)
	}
}
