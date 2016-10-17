package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}

	// convert the text read to string
	str := string(bs)
	fmt.Println("The content is: ", str)
}
