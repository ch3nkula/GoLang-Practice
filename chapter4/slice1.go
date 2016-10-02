package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}

	/* Appending to a slice */
	slice2 := append(slice1, 4, 5)

	slice3 := make([]int, 2)

	/* Copying slices */
	copy(slice3, slice1)

	fmt.Println(slice1, slice2)

	fmt.Println(slice1, slice3)

	fmt.Println(slice1[0:2])
}