package main

import (
	"fmt"
	"gowebspider"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("start...")

	spider := gowebspider.NewWebSpider()
	spider.Fetch("http://image.baidu.com/channel/index", "testpath")

	fmt.Printf("time cost %v\n", time.Now().Sub(start))
}
