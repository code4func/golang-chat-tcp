package main

import (
	"net"
	"log"
	"fmt"
	"os"
	"bufio"
)

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')

		fmt.Print(msg)
	}
}

func main() {
    connection, err := net.Dial("tcp", "localhost:3000")
    if err != nil {
    	log.Fatal(err)
    }

    fmt.Print("your name:")
    nameReader := bufio.NewReader(os.Stdin)
    nameInput, _ := nameReader.ReadString('\n')

    nameInput = nameInput[:len(nameInput) - 1]
    
    fmt.Println("********** MESSAGES **********")

    go onMessage(connection)

    for {
    	 msgReader := bufio.NewReader(os.Stdin)
    	 msg, err := msgReader.ReadString('\n')
    	 if err != nil {
    	 	break
    	 }

    	 msg = fmt.Sprintf("%s: %s\n", nameInput, 
    	 					msg[:len(msg) - 1])

    	 connection.Write([]byte(msg))					
    }	

    connection.Close()
}












