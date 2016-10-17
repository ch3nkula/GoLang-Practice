/* Sort Package in GoLang */

package main

import ("fmt" ; "sort")

/* Person structure */
type Person struct {
	Name string
	Age int
}

/* ByName has type an array of struct Person */
type ByName []Person

/* Length of the persons name */
func (this ByName) Len() int {
	return len(this)
}

/* Less than function */
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

/* Swap the two values compared */
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

/* Main function */
func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}

	/* Sort the array of structures */
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}