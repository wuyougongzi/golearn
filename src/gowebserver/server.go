package main

import (
	"fmt"
	"goutility"
	//"io/ioutil"
	"net"
	//"strconv"
	"strings"
	"time"
)

func server() {
	service := ":1200" //port

	fmt.Println("create client ip addr")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if goutility.CheckErr(err) {
		return
	}

	fmt.Println("listen client request")
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if goutility.CheckErr(err) {
		return
	}

	for {
		conn, err := listener.Accept()
		if goutility.CheckErr(err) {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128) //set maxium request length to 128B to prevent flood attack
	defer conn.Close()

	fmt.Println("read client request")
	//for {
	read_len, err := conn.Read(request)
	fmt.Println(read_len)
	if goutility.CheckErr(err) {
		return
	}

	if read_len == 0 {
		return
	} /*else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
		dateTime := strconv.FormatInt(time.Now().Unix(), 10)
		conn.Write([]byte(dateTime))
	} else {
		respond := string("from server")
		dayTime := time.Now().String()
		respond = strings.Join([]string{respond, " ", dayTime}, "")
		conn.Write([]byte(respond))
	}
	*/
	fmt.Println(string(request[:read_len]))
	request = make([]byte, 128)

	//	}
	//clientrequest, err := ioutil.ReadAll(conn)
	//	if goutility.CheckErr(err) {
	//		return
	//	}
	//	fmt.Println(clientrequest)

	respond := string("from server")
	dayTime := time.Now().String()
	respond = strings.Join([]string{respond, " ", dayTime}, "")
	conn.Write([]byte(respond))
}

func main() {
	/*
		respond := string("from server")
		dateTime := time.Now().String()
		respond = strings.Join([]string{respond, " ", dateTime}, "")
		fmt.Println(respond)
	*/
	server()
}
