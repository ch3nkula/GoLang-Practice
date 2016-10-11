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

	/* Execute the Close() if error occurs */
	defer file.Close()

	/* Write to the file */
	file.WriteString("test")
}