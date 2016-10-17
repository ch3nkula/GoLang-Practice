/**
 * Even Number Generator
 */

package main

import "fmt"

func EvenNumberGenerator() func() uint {
	i := uint(0)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// The main function
func main() {
	nextEven := EvenNumberGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
