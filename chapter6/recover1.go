/* Recover a GoLang program from a panic. That is a 
   run-time error */

package main

import "fmt"

func main() {

	panic("PANIC")

	str := recover()

	fmt.Println(str)

}