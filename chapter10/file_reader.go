/* Reading content of a file and output on the terminal in GoLang */

package main 

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		/* Handle the error here */
		return
	}

	defer file.Close()

	/* Get the file size */
	stat, err := file.Stat()
	if err != nil {
		return
	}

	/* Read the file */
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println("The content of the file is: ", str)
}