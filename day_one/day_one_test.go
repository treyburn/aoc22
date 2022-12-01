package day_one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCalories(t *testing.T) {
	var testData = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	want := []int{6000, 4000, 11000, 24000, 10000}

	got := CountCalories(testData)

	assert.Equal(t, want, got)
}

func TestFindHighestCalorie(t *testing.T) {
	type testCase struct {
		name string
		data []int
		want int
	}

	var testData = []testCase{
		{"standard", []int{6000, 4000, 11000, 24000, 10000}, 24000},
		{"empty", []int{}, -1},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			tc := test
			got := FindHighestCalorie(tc.data)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSumTop3Calories(t *testing.T) {
	type testCase struct {
		name string
		data []int
		want int
	}

	var testData = []testCase{
		{"standard", []int{6000, 4000, 11000, 24000, 10000}, 45000},
		{"empty", []int{}, -1},
		{"two", []int{100, 200}, 300},
		{"one", []int{100}, 100},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			tc := test
			got := SumTop3Calories(tc.data)

			assert.Equal(t, tc.want, got)
		})
	}
}
