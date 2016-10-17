/* Socket programming in GoLang. TCP Server */

package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

/* Server function */

func server() {
	/* Listen on port 9999 */
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		/* Accept a connection */
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		/* Handle the connection */
		go handleServerConnection(c)
	}
}

/* Handle connection function */

func handleServerConnection(c net.Conn) {
	/* Receive the message */
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

/* Client function */

func client() {
	/* Connect to the Server */
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Send a message */
	msg := "Hello World from Client Computer"
	fmt.Println("Sending", msg)

	err = gob.NewEncoder(c).Encode(msg)

	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

/* Main function */

func main() {
	/* Run the goroutines :) */
	go server()
	go client()

	var input string
	fmt.Scanln(&input)

}
