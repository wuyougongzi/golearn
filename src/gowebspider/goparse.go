package gowebspider

import (
	//"fmt"
	"gostl"
	//"goutility"
)

func Parse(respStr string) *gostl.Queue {
	urlQueue := new(gostl.Queue)
	urlQueue.Push("test")
	return urlQueue
}

/*
func main() {
	respStr, err := HttpDownload("https://www.baidu.com")
	goutility.CheckErr(err)
	//urlQ := new(gostl.Queue)
	urlQ := Parse(respStr)
	fmt.Println(urlQ.Front())

}
*/
