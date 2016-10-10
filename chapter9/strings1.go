/* Working and Manipulating strings in GoLang */

package main

/* Importing a list of packages */
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
			/* True */
			strings.Contains("test", "es"),

			/* Count is 2 */
			strings.Count("test", "t"),

			/* True */
			strings.HasPrefix("tests", "te"),

			/* True */
			strings.HasSuffix("test", "st"),

			/* Index is 1 */
			strings.Index("test", "e"),

			/* "a - b" */
			strings.Join([]string{"a", "b"}, "-"),

			/* == "aaaaa" */
			strings.Repeat("a", 5),

			/* "bbaa" */
			strings.Replace("aaaa", "a", "b", 2),

			/* []string{"a", "b", "c", "d", "e"} */
			strings.Split("a-b-c-d-e", "-"),

			/* "test" */
			strings.ToLower("TEST"),

			/* "TEST" */
			strings.ToUpper("test"),
		)

	/* Converting a byte to a string and vise versa */
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})

	fmt.Println("This is the byte: ", arr)

	fmt.Println("This is the string: ", str)
}