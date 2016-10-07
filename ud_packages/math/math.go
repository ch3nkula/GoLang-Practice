/* This is the math package that will be used in the chapter 9 folder */

package math

func Average(xs []float64) float64 {
	total := float64(0)

	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}