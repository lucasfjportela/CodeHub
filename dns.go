package main

import (
	"codehub-sd/messageFormat"
	"encoding/gob"
	"fmt"
	"net"
)

type dns struct {
	/*map["Serverx : "1994864:1554""]*/
	table map[string]string
}

func (d *dns) handleDNSConnection(conn *net.TCPConn) {

	encoder := gob.NewEncoder(conn)

	for _, ipp := range d.table {
		tcpAddrServer, _ := net.ResolveTCPAddr("tcp", ipp)
		conn, err := net.DialTCP("tcp", nil, tcpAddrServer)
		encoderServer := gob.NewEncoder(conn)

		if err != nil {
			continue
		}

		msg := messageFormat.MessageFormat{Origin: "DNS", ReqType: "ver"}

		encoderServer.Encode(msg)
		fmt.Println("Client requests server address")
		encoder.Encode(ipp)

		conn.Close()
		return
	}

}

func main() {
	//var dnsTable dns
	dnsTable := &dns{}
	dnsTable.table = make(map[string]string)
	dnsTable.table["Server1"] = "localhost:1111"
	dnsTable.table["Server2"] = "localhost:1112"
	dnsTable.table["Server3"] = "localhost:1113"

	fmt.Println("Starting DNS Server...")
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:2223")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	var tcpConn net.TCPConn
	defer tcpConn.Close()

	for {
		tcpConn, _ := listener.AcceptTCP()
		go dnsTable.handleDNSConnection(tcpConn)
	}
}
