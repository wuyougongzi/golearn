package gowebspider

//package main

import (
	//"fmt"
	"gostl"
	"goutility"
	"io/ioutil"
	"net/http"
)

func HttpDownload(url string) (respStr string, err error) {
	resp, err := http.Get(url)
	if goutility.CheckErr(err) {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	respStr = string(body)
	if goutility.CheckErr(err) {
		return
	}

	return
}

//use the string pattern match find url
func findUrl(descStr string, patternStr string, queue *gostl.Queue) {

}

func ParseRespon(respStr string) *gostl.Queue {
	urlQueue := new(gostl.Queue)

	//urlQueue.Push("test")

	return urlQueue
}

/*
func main() {
	respStr, err := HttpDownload("https://www.baidu.com")
	goutility.CheckErr(err)
	urlQ := ParseRespon(respStr)
	//fmt.Println(urlQ.Front())
}
*/
