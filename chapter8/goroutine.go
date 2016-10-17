/* Concurrency in GoLang using Goroutines */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	/* The Goroutine */
	for i := 0; i < 10; i++ {
		go f(0)
	}

	var input string
	fmt.Print("Enter a number: ")
	fmt.Scanln(&input)
}
