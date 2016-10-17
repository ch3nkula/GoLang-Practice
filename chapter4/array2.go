package main

import "fmt"

func main() {
	x := [5]float64{98, 93, 77, 82, 83}

	for _, value := range x {
		fmt.Print(value, " ")
	}

	fmt.Println()
}
