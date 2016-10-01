package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")

	var number int32

	fmt.Scanf("%d", &number)

	switch number {
		case 0: fmt.Println("Zero")

		case 1: fmt.Println("One")

		default: fmt.Println("Number not 0 or 1 :)")
	}
}