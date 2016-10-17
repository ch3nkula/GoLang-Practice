/**
 * Closure: Definitions of functions in other functions 
 */

package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}

	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println("Incremented i: ", increment())
	fmt.Println("Incremented i", increment())

	fmt.Println("Addition of 1 and 1: ", add(1, 1))
}
