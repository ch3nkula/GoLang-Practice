/**
 * User a the main package 
 */
package main

// Import the input/output library
import "fmt"

// main function
func main() {
	fmt.Print("Enter a number: ")
	var input float64

	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
