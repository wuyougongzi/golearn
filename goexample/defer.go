package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("test.file")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("createing")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing...")

	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closeing...")

	f.Close()
}
