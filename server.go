package main

import (
	"fmt"
	"net"
)

/*Func that execute the server operation*/
func handleServerConnection(conn net.Conn) {
	readed := make([]byte, 1024)
	datasize, _ := conn.Read(readed)
	data := readed[:datasize]
	fmt.Print(string(data))
}

func main() {

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:2222")
	test, _ := net.ListenTCP("tcp", tcpAddr)

	var tcpConn net.TCPConn

	defer tcpConn.Close()

	for {
		tcpConn, _ := test.Accept()
		go handleServerConnection(tcpConn)
	}

}
