/* Maps are associative arrays, an unordered collection of key-value
   pairs. For example, a hash table, dictionary etc... */

package main

import "fmt"

func main() {
	//var x map[string]int

	x := make(map[string]int)

	x["key"] = 10

	fmt.Println(x)
}
