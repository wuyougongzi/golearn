//go 加密和解密示例
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	//"os"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func base64FuncTest() {
	//encode
	hello := "hello world !!! heimeng!!!"
	debyte := base64Encode([]byte(hello))
	fmt.Println(debyte)

	//decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		panic("decode error")
	}

	if hello != string(enbyte) {
		fmt.Println("encode and decode is not equal")
	}

	fmt.Println(enbyte)
}

var commonIV = []byte{0x00, 0x01, 0x02, 0x03,
	0x04, 0x05, 0x06, 0x07,
	0x08, 0x09, 0x0a, 0x0b,
	0x0c, 0x0d, 0x0e, 0x0f}

func main() {
	//base64 test
	//base64FuncTest()

	plantText := []byte("hello world !!! heimeng!!!")

	/*if len(os.Args) > 1 {
		plantText = []byte(os.Args[1])
	}
	*/
	keyText := "123456789abcdefg"

	fmt.Println(keyText)

	block, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		panic("aes error")
	}

	//加密
	cfb := cipher.NewCFBEncrypter(block, commonIV)
	ciphertext := make([]byte, len(plantText))
	cfb.XORKeyStream(ciphertext, plantText)
	fmt.Println("%s=>%x\n", plantText, ciphertext)

	//解密
	cfbdec := cipher.NewCFBDecrypter(block, commonIV)
	plantTextCopy := make([]byte, len(plantText))
	cfbdec.XORKeyStream(plantTextCopy, ciphertext)
	fmt.Println("%x=>%s\n", ciphertext, plantTextCopy)
}
