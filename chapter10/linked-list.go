/* A singly linked list in GoLang */

package main

import (
	"fmt"
	"container/list"
)

/* Main function */
func main() {
	var x list.List

	/* Pushing Items to the Linked list */
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	/* Loop through the list */
	fmt.Print("The list is: ")
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(int), " ")
	}

	/* Print a new line after the list ends */
	fmt.Println()
}