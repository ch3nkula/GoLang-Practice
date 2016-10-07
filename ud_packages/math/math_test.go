/* Testing the math package */

package math

import "testing"

/* Test pair structure */
type testpair struct {
	values []float64
	average float64
}

type testpairsum struct {
	values []int
	sum int
}

var tests = []testpair {
	{ []float64{1, 2}, 1.5 },
	{ []float64{1, 1, 1, 1, 1, 1}, 1 },
	{ []float64{-1, 1}, 0 },
}

var testsum = []testpairsum {
	{ []int{1, 2, 3, 4, 5}, 15 },
	{ []int{1, 2, 3}, 6 },
}

/* Testing the Average function */
func TestAverage(t *testing.T) {
	var v float64

	v = Average([]float64{1, 2})

	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

/* Testing the Sum function */
func TestSum(t *testing.T) {
	var s int

	s = Sum(1, 2, 3, 4)

	if s != 10 {
		t.Error("Expected 10, got ", s)
	}
}

/* Testing average pairs */
func TestAveragePair(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
					"For", pair.values,
					"expected", pair.average,
					"got", v,
				)
		}
	}
}

/* Testing sum pairs */
func TestSumPair(t *testing.T) {
	for _, pair := range testsum {
		s := Sum(pair.values...)
		if s != pair.sum {
			t.Error(
					"For", pair.values,
					"expected", pair.sum,
					"got", s,
				)
		}
	}
}