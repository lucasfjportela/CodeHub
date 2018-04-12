package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type req struct {
	login    string
	password string
}

func handleClientDNSConnection(conn *net.TCPConn, c chan string) {

	var dnsResponse string

	decoder := gob.NewDecoder(conn)

	decoder.Decode(&dnsResponse)

	//readed := make([]byte, 1024)
	//datasize, _ := conn.Read(readed)
	//data := readed[:datasize]
	c <- dnsResponse
}

func handleAuthConnection() {

}

func handleClientServerConnection() {

}

func main() {

	tcpAddrDNS, _ := net.ResolveTCPAddr("tcp", "localhost:2223")

	//for {
	c := make(chan string)
	conn, _ := net.DialTCP("tcp", nil, tcpAddrDNS)
	conn.Close()

	go handleClientDNSConnection(conn, c)
	a := <-c

	fmt.Println(a)

	/*encoder := gob.NewEncoder(conn)
	err := encoder.Encode(a)

	if err != nil {
		panic(err)
	}
	*/
	//}

	//conn.Write([]byte(m))

}
