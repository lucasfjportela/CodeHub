package main

import (
	"../codehub-sd/messageFormat"
	"encoding/gob"
	"fmt"
	"net"
)



func handleClientAuthentication(conn *net.TCPConn){
	mu := messageFormat.MessageFormat{}

	decoder := gob.NewDecoder(conn)
	decoder.Decode(&mu)

	payload := mu.Payload.([]string)

	fmt.Println( payload )
	/*if  {
		fmt.Println("AUTORIZADO	")
	}*/

}

func main (){
	fmt.Println("Starting AuthServer...")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:1115")
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
	