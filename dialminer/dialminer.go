/*
	Simple package to make a tcp connection to a miner
	service.
*/

package dialminer

import (
	"fmt"
	"net"
	"os"
)

func DialMiner(host string, port string, command string) ([]byte, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Resolve error!", err.Error())
		fmt.Println("Host:", host, "Port:", port)
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("TCP Dial error!", err.Error())
		return nil, err
	}
	_, err = conn.Write([]byte(command))
	if err != nil {
		fmt.Println("Write error!", err.Error())
		return nil, err
	}
	reply := make([]byte, 4096)
	_, err = conn.Read(reply)
	if err != nil {
		fmt.Println("Read error!", err.Error())
		return nil, err
	}
	return reply, nil
}
