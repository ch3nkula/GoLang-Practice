/* Testing the math package */

package math

import "testing"

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