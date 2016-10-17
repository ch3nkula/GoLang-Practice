/**
 * Variadic Functions 
 */

package main

import "fmt"

func add(args ...int) int {
	total := 0

	for _, v := range args {
		total += v
	}

	return total
}

func main() {
	slice := []int{4, 5, 6}

	fmt.Println(add(slice...))

	fmt.Println(add(1, 2, 3))
}
