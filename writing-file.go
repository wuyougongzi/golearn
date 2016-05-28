package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(("dat.data"), d1, 0644)
	check(err)

	f, err := os.Create("dat.data1")
	check(err)

	defer f.Close()

	d2 := []byte{115, 11, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("write %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("write %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered \n")
	fmt.Printf("write %d bytes\n", n4)

	w.Flush()
}
