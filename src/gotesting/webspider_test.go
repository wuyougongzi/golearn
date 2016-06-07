package gotesting

import (
	"fmt"
	"gowebspider"
	"testing"
)

func TestDownloader(t *testing.T) {
	respStr, _ := gowebspider.HttpDownload("http://www.baidu.com/")
	fmt.Println(respStr)
}
