package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	port := YOUR_PORT_HERE
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		// Read from standard in and write to the connection
		// until "Done" is typed.
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text)

		if text == "Done\n" {
			break
		}
	}
}
