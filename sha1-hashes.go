package main

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"fmt"
)

// sha1 used exmple
func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	///////////////////////////////////////
	fmt.Println()
	fmt.Println("///////////////////////////////////////////////////")
	fmt.Println()

	data := "abcdefg123!?$*"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, err := b64.StdEncoding.DecodeString(sEnc)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(sDec))
	}

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
