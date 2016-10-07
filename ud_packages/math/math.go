/* This is the math package that will be used in the chapter 9 folder */

package math

/* Finds the average of a series of numbers */
func Average(xs []float64) float64 {
	total := float64(0)

	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}

/* Finds the sum of a series of numbers */
func Sum(x ...int) int {
	total := int(0)

	for _, sum := range x {
		total += sum
	}

	return total
}