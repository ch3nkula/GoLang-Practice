/* Implementation of the SHA-1 crypto hash function in GoLang */

package main 

/* Importing the crypto sha1 package */

import (
	"fmt"
	"crypto/sha1"
)

/* Main function */

func main() {
	
	h := sha1.New()
	h.Write([]byte("test"))

	bs := h.Sum([]byte{})
	fmt.Println(bs)

}