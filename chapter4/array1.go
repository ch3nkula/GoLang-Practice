package main

import "fmt"

func main() {
	var x [5]int

	/* Another kind of for loop */
	for _, value := range x {
		fmt.Print(value)
	}

	x[4] = 10

	fmt.Println("Array Values are: ", x)
}