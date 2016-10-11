/* Creating a file in GoLang */

package main

import (
	"os"
)

/* Main function */
func main() {
	file, err := os.Create("testing.txt")

	if err != nil {
		/* Handle the Error */
		return
	}

	defer file.Close()

	file.WriteString("test")
}