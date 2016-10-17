/* Implementation of the CRC32 non-cryptographic hash funciton in GoLang */

package main

/* Import the packages */

import (
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()

	/* Writing bytes to the Writer interface of the */
	/* crc32 hash object */
	h.Write([]byte("test"))

	/* Sum32 returns a uint32 */
	v := h.Sum32()
	fmt.Println(v)
}
