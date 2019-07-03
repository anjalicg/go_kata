// https://systembash.com/a-simple-go-tcp-server-and-tcp-client/

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listen, _ := net.Listen("tcp", ":8080")
	conn, _ := listen.Accept()
	fmt.Println(conn)
	defer listen.Close()
	fmt.Printf("%T\n", conn)
	for {
		// will listen for message to process ending in newline (\n)
		fmt.Println("Conn=", conn)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
