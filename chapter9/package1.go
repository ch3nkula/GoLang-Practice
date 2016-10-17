/**
 * Creating and using packages in GoLang 
 */

package main

import "fmt"
import ud_math "go-book/ud_packages/math"

func main() {
	xs := []float64{1, 2, 3, 4}

	// Use the package
	avg := ud_math.Average(xs)

	sum := ud_math.Sum(1, 2, 3, 4)

	fmt.Println("Sum is: ", sum)

	fmt.Println("Average is: ", avg)
}
