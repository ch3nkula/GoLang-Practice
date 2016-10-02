/* A program to half a number and check if it's half is 
is even or false and return true or false respectively */

package main

import "fmt"

func half(value int) bool {
	x := value / 2

	if x % 2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	
	var value int

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &value)

	status := half(value)

	fmt.Println(value / 2, status)
}