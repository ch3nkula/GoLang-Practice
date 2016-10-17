package main

// Importing the required packages
import (
	"fmt"
	"os"
	"path/filepath"
)

// Main function
func main() {
	// Walk the directory tree using Walk() Golang method
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
