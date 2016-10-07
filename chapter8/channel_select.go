/* Selecting a particular channel using the special select keyword in GoLang */

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "From 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "From 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <- c1:
					fmt.Println(msg1)
			case msg2 := <- c2:
					fmt.Println(msg2)
			}
		}
	}()

	var input string

	fmt.Scanln(&input)
}