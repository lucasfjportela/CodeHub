package main

import (
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"net"

	"codehub-sd/messageFormat"
)

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func handleClientAuthentication(conn *net.TCPConn) {

	/*Test users*/

	test := [][]string{{"matheu", "123"}, {"bean", "456"}}

	auth := false

	mu := messageFormat.MessageFormat{}
	decoder := gob.NewDecoder(conn)
	decoder.Decode(&mu)

	payload := mu.Payload.([]string)

	decryptPassword, _ := base64Decode([]byte(payload[1]))

	for _, i := range test {
		if i[0] == payload[0] && i[1] == string(decryptPassword) {
			auth = true
			break
		}
	}

	encoder := gob.NewEncoder(conn)
	encoder.Encode(auth)

}

func main() {
	fmt.Println("Starting AuthServer...")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "192.168.0.105:1515")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	var tcpConn net.TCPConn

	defer tcpConn.Close()

	for {
		fmt.Println("Listening in AuthServer...")
		tcpConn, _ := listener.AcceptTCP()
		go handleClientAuthentication(tcpConn)
	}
}