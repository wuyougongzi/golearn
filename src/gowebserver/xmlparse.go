package main

import (
	"encoding/xml"
	"fmt"
	"goutility"
	"io/ioutil"
	"os"
)

type Resurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version, attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

/*
servers.xml eg.

<?xml version="1.0" encoding="utf-8"?>
<servers version="1">
    <server>
        <serverName>Shanghai_VPN</serverName>
        <serverIP>127.0.0.1</serverIP>
    </server>
    <server>
        <serverName>Beijing_VPN</serverName>
        <serverIP>127.0.0.2</serverIP>
    </server>
</servers>
*/

func readXml() {
	file, err := os.Open("servers.xml")
	if goutility.CheckErr(err) {
		return
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if goutility.CheckErr(err) {
		return
	}

	v := Resurlyservers{}
	err = xml.Unmarshal(data, &v)
	if goutility.CheckErr(err) {
		return
	}

	fmt.Println(v)
}

func writeXml() {
	v := &Resurlyservers{Version: "1"}
	serverName := xml.Name{" ", " "} //使用一个空的xml.name用以保存通用server结构
	v.Svs = append(v.Svs, server{serverName, "shanghai_vpn", "127.0.0.1"})
	v.Svs = append(v.Svs, server{serverName, "beijing_vpn", "198.0.0.1"})
	output, err := xml.MarshalIndent(v, " ", " ")
	if goutility.CheckErr(err) {
		return
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

func main() {
	readXml()
	writeXml()
}
