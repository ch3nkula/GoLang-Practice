/* Comparing two files if they are different in GoLang using the */
/* crc32 has funciton. If they are the same, its highly likely that their */
/* hashes are the same (not 100% certain) else their hashes will be diff */

package main

/* Importing the packages */

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

/* Get hash function definition */

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		return 0, err
	}

	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

/* Main function */
/* As an example, just try it with two empty files */
/* Output of 2 empty files should give "0 0 true" */
func main() {

	/* Get the hash of test1.txt */
	h1, err := getHash("test1.txt")
	if err != nil {
		return
	}

	/* Get the hash of test2.txt file */
	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}

	/* prints if the files are the same :) */
	fmt.Println(h1, h2, h1 == h2)
}
