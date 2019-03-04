package helpers

import (
	"fmt"
	"net"
)

//this is my comment
func MakeConnection(ipaddress string) (net.UDPConn, error) {
	address, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:9760", ipaddress))
	CheckError(err)
	Conn, err := net.DialUDP("udp", nil, address)
	CheckError(err)
	// if err != nil {
	// 	return nil, err
	// }
	if err == nil {
		fmt.Println("I think it dialed correctly")
	}
	return *Conn, nil
}

//Check the error
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
