package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"../codehub-sd/messageFormat"
)

func handleClientAuthentication(conn *net.TCPConn) {

	/*Test users*/

	test := [][]string{{"matheu", "123"}, {"bean", "456"}}

	auth := false

	mu := messageFormat.MessageFormat{}
	decoder := gob.NewDecoder(conn)
	decoder.Decode(&mu)

	payload := mu.Payload.([]string)

	for _, i := range test {
		if i[0] == payload[0] && i[1] == payload[1] {
			auth = true
			break
		}
	}

	encoder := gob.NewEncoder(conn)
	encoder.Encode(auth)

}

func main() {
	fmt.Println("Starting AuthServer...")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:1515")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	var tcpConn net.TCPConn

	defer tcpConn.Close()

	for {
		fmt.Println("Listening in AuthServer...")
		tcpConn, _ := listener.AcceptTCP()
		fmt.Println("Dále nessa autenticação")
		go handleClientAuthentication(tcpConn)
	}

	//m := messageFormat{ origin : "client", reqType : "auth", payload : {"Teteu", "123"}}

}
