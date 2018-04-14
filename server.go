package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"../codehub-sd/messageFormat"
)

/*Func that execute the server operation*/
func handleServerConnection(conn *net.TCPConn) {

	/*var data map[string]string

	decoder := gob.NewDecoder(conn)

	decoder.Decode(&data)

	//readed := make([]byte, 1024)
	//datasize, _ := conn.Read(readed)
	//data := readed[:datasize]
	fmt.Println(data)
	*/
	msg := messageFormat.MessageFormat{}

	decoder := gob.NewDecoder(conn)
	decoder.Decode(&msg)

	if msg.Origin == "Client" {
		if msg.ReqType == "sto" {
			fmt.Println("Aqui guarda o arquivo")
		}

		if msg.ReqType == "get" {
			fmt.Println("Aqui dá o arquivo pra ele")
		}
	}

}

func main() {
	fmt.Println("Starting Server...")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:1111")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	var tcpConn net.TCPConn

	defer tcpConn.Close()

	for {
		fmt.Println("Listening...")
		tcpConn, _ := listener.AcceptTCP()
		fmt.Println("Dále")
		go handleServerConnection(tcpConn)
	}

}
