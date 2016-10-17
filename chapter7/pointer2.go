/* Pointers in GoLang using the new built-in function */

package main

import "fmt"

/* function to assign 1 to where "ptr" is pointing to */
func one(ptr *int) {
	*ptr = 1
}

/* Main function */
func main() {
	ptr := new(int)

	one(ptr)

	fmt.Println("ptr points to memory address with value: ", *ptr) //ptr is 1
}
