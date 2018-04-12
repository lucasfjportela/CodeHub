package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

/*Func that execute the server operation*/
func handleServerConnection(conn net.Conn) {

	var data map[string]string

	decoder := gob.NewDecoder(conn)

	decoder.Decode(&data)

	//readed := make([]byte, 1024)
	//datasize, _ := conn.Read(readed)
	//data := readed[:datasize]
	fmt.Println(data)
}

func main() {
	fmt.Println("Starting Server...")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:1111")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	var tcpConn net.TCPConn

	defer tcpConn.Close()

	for {
		fmt.Println("Listening...")
		tcpConn, _ := listener.Accept()
		fmt.Println("Dále")
		go handleServerConnection(tcpConn)
	}

}
