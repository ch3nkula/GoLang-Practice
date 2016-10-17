/**
 * Recover a GoLang program from a panic. That is a
 * run-time error 
 */

package main

import "fmt"

func main() {

	defer func() {
		str := recover()

		fmt.Println(str)
	}()

	panic("PANIC")

}
