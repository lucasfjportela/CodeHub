package main

import (
	"codehub-sd/messageFormat"
	"encoding/gob"
	"fmt"
	"net"
)

/*
type Req struct {
	Login      string
	Password   string
	Authorized bool
}

*/

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

func handleClientAuthConnection(conn *net.TCPConn, msgUser messageFormat.MessageFormat) bool {

	encoderServer := gob.NewEncoder(conn)
	encoderServer.Encode(msgUser)

	var authResponse bool

	decoder := gob.NewDecoder(conn)

	decoder.Decode(&authResponse)

	//readed := make([]byte, 1024)
	//datasize, _ := conn.Read(readed)
	//data := readed[:datasize]
	return authResponse
}

func handleClientServerConnection() {

}

func main() {

	//var list ListUsers
	//pl := {"bean", "123"}.([]string)
	tcpAddrAUTH, _ := net.ResolveTCPAddr("tcp", "localhost:1115")
	tcpAddrDNS, _ := net.ResolveTCPAddr("tcp", "localhost:2223")
	c := make(chan string)

	conn, _ := net.DialTCP("tcp", nil, tcpAddrAUTH)

	msgUser := messageFormat.MessageFormat{Origin: "CLIENT", ReqType: "auth", Payload: []string{"bean", "123"}}

	ver := handleClientAuthConnection(conn, msgUser)

	/*Verify User auth*/
	if ver == false {
		fmt.Println("Lixo, não pode conectar")
	} else {

		conn2, _ := net.DialTCP("tcp", nil, tcpAddrDNS)

		go handleClientDNSConnection(conn2, c)
		serverAddr := <-c
		fmt.Println(serverAddr)
	}

	/*encoder := gob.NewEncoder(conn)
	err := encoder.Encode(a)

	if err != nil {
		panic(err)
	}
	*/
	//}

	//conn.Write([]byte(m))

}
