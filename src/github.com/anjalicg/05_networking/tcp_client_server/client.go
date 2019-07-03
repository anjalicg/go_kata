// https://systembash.com/a-simple-go-tcp-server-and-tcp-client/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin) // to read from console
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n") // sending text to connection object
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n') // reading from connection
		fmt.Print("Message from server: " + message)
	}

}
