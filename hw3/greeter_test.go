package greeter_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	Grt "gitlab.ozon.dev/qa/classroom-4/act-device-api/hw3"
	"math"
	"testing"
)

func TestGreet(t *testing.T) {
	type testCase struct {
		name     string
		hour     int
		expected string
	}
	testsHour := []testCase{
		{"N", math.MinInt, "Hello N!"},
		{"N", -1, "Hello N!"},
		{"N", 0, "Good night N!"},
		{"N", 6, "Good morning N!"},
		{"N", 12, "Hello N!"},
		{"N", 18, "Good evening N!"},
		{"N", 22, "Good night N!"},
		{"N", 24, "Hello N!"},
		{"N", math.MaxInt, "Hello N!"},
	}
	testsTitleName := []testCase{
		{"NIK", 0, "Good night NIK!"},
		{"nik", 0, "Good night Nik!"},
		{"Nik", 0, "Good night Nik!"},
		{"", 0, "Good night !"},
		{"123", 0, "Good night 123!"},
	}
	testsSpaceName := []testCase{
		{"qwe ", 0, "Good night Qwe!"},
		{" qwe", 0, "Good night Qwe!"},
		{"   ", 0, "Good night !"},
		{"qwe", 0, "Good night Qwe!"},
	}
	for _, tc := range testsHour {
		testName := fmt.Sprintf("%s, %d, %s", tc.name, tc.hour, tc.expected)
		t.Run(testName, func(t *testing.T) {
			result := Grt.Greet(tc.name, tc.hour)
			assert.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
	for _, tc := range testsTitleName {
		testName := fmt.Sprintf("%s, %d, %s", tc.name, tc.hour, tc.expected)
		t.Run(testName, func(t *testing.T) {
			result := Grt.Greet(tc.name, tc.hour)
			assert.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
	for _, tc := range testsSpaceName {
		testName := fmt.Sprintf("%s, %d, %s", tc.name, tc.hour, tc.expected)
		t.Run(testName, func(t *testing.T) {
			result := Grt.Greet(tc.name, tc.hour)
			assert.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
func TestNewGreet(t *testing.T) {
	type testCase struct {
		name     string
		hour     int
		expected string
	}
	testsHour := []testCase{
		{"N", -1, "Hello N!"},
		{"N", 0, "Good night N!"},
		{"N", 6, "Good morning N!"},
		{"N", 12, "Hello N!"},
		{"N", 18, "Good evening N!"},
		{"N", 22, "Good night N!"},
		{"N", 24, "Hello N!"},
	}
	testsTitleName := []testCase{
		{"NIK", 0, "Good night NIK!"},
		{"nik", 0, "Good night Nik!"},
		{"Nik", 0, "Good night Nik!"},
		{"", 0, "Good night !"},
		{"123", 0, "Good night 123!"},
	}
	testsSpaceName := []testCase{
		{"qwe ", 0, "Good night Qwe!"},
		{" qwe", 0, "Good night Qwe!"},
		{"   ", 0, "Good night !"},
		{"qwe", 0, "Good night Qwe!"},
	}
	for _, tc := range testsHour {
		testName := fmt.Sprintf("%s, %d, %s", tc.name, tc.hour, tc.expected)
		t.Run(testName, func(t *testing.T) {
			result := Grt.NewGreet(tc.name, tc.hour)
			assert.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
	for _, tc := range testsTitleName {
		testName := fmt.Sprintf("%s, %d, %s", tc.name, tc.hour, tc.expected)
		t.Run(testName, func(t *testing.T) {
			result := Grt.NewGreet(tc.name, tc.hour)
			assert.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
	for _, tc := range testsSpaceName {
		testName := fmt.Sprintf("%s, %d, %s", tc.name, tc.hour, tc.expected)
		t.Run(testName, func(t *testing.T) {
			result := Grt.NewGreet(tc.name, tc.hour)
			assert.Equal(t, tc.expected, result, "expected: %v, result: %v", tc.expected, result)
		})
	}
}
