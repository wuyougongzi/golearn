package main

import (
	"fmt"
	"goutility"
	"io/ioutil"
	"net"
	"os"
)

func client() {
	service := os.Args[1]
	fmt.Println("create addr")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if goutility.CheckErr(err) {
		return
	}

	fmt.Println("connect to remote server")
	tpcConn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if goutility.CheckErr(err) {
		return
	}

	_, err = tpcConn.Write([]byte("HEAD / HTTP/1.0"))

	fmt.Println("read server return")
	result, err := ioutil.ReadAll(tpcConn)
	if goutility.CheckErr(err) {
		return
	}
	fmt.Println(string(result))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	client()
	//parse ip
	/*	name := os.Args[1]
		addr := net.ParseIP(name)
		if addr == nil {
			fmt.Println("Invalid addr ip")
		} else {
			fmt.Println("the addr", addr)
		}
		os.Exit(0)
	*/

}
