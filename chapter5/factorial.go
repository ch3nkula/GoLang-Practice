/**
 * Factorial function based on recursion in GoLang
 */

package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func main() {
	var factor uint

	factor = 5

	fmt.Println(factorial(factor))
}
