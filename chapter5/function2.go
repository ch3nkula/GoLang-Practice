/* Returning multiple values from a function */

package main

import "fmt"

func multiple_return() (int, int) {
	return 5, 6
}

func array_return() ([]int) {
	return []int{1, 2}
}

func main() {
	x, y := multiple_return()

	fmt.Println(x, y)

	arr := array_return()

	fmt.Println(arr)
}