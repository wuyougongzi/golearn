package gowebspider

import (
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
	//fmt.Println(string(body))
	if goutility.CheckErr(err) {
		return
	}

	return
}
