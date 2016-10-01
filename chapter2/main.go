package main

import "fmt"

var x string = "Hello World!!!"

func main(){
	var helloWorld string = "Hello World!"

	var hw string

	hw = "Hello World!!"

	fmt.Println(helloWorld)

	fmt.Println(hw)

	printText()
}

func printText() {
	fmt.Println(x)
}