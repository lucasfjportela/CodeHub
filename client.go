package main

import (
	"../codehub-sd/messageFormat"
	"encoding/gob"
	"fmt"
	"net"
)

type Req struct {
	Login    string
	Password string
	Authorized bool
}

//type ListUsers [] messageFormat.MessageFormat

func handleClientDNSConnection(conn *net.TCPConn, c chan string) {

	var dnsResponse string

	decoder := gob.NewDecoder(conn)

	decoder.Decode(&dnsResponse)

	//readed := make([]byte, 1024)
	//datasize, _ := conn.Read(readed)
	//data := readed[:datasize]
	c <- dnsResponse
}

func handleClientAuthConnection(conn *net.TCPConn, a chan string) {
	var authResponse string

	decoder := gob.NewDecoder(conn)

	decoder.Decode(&authResponse)

	//readed := make([]byte, 1024)
	//datasize, _ := conn.Read(readed)
	//data := readed[:datasize]
	a <- authResponse
}

func handleClientServerConnection() {

}

func main() {

	//var list ListUsers
	//pl := {"bean", "123"}.([]string)
	tcpAddrAUTH, _ := net.ResolveTCPAddr("tcp", "localhost:1115")

	a := make(chan string)
		
	conn, _ := net.DialTCP("tcp", nil, tcpAddrAUTH)
	
	encoderServer := gob.NewEncoder(conn)
	msgUser := messageFormat.MessageFormat{Origin: "CLIENT", ReqType: "auth", Payload: []string{"bean", "123"}}
	encoderServer.Encode(msgUser)

	go handleClientAuthConnection(conn, a)
	ac := <-a
	fmt.Println(ac)

	tcpAddrDNS, _ := net.ResolveTCPAddr("tcp", "localhost:2223")

	c := make(chan string)
		
	conn2, _ := net.DialTCP("tcp", nil, tcpAddrDNS)
		
	go handleClientDNSConnection(conn2, c)
	s := <-c
	fmt.Println(s)
	conn.Close()
	conn2.Close()

	/*encoder := gob.NewEncoder(conn)
	err := encoder.Encode(a)

	if err != nil {
		panic(err)
	}
	*/
	//}

	//conn.Write([]byte(m))

}
