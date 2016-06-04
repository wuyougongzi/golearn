package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//spawing-proc
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	checkError(err)

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	gerpCmd := exec.Command("grep", "hello")
	grepIn, _ := gerpCmd.StdinPipe()
	gerpOut, _ := gerpCmd.StdoutPipe()
	gerpCmd.Start()
	grepIn.Write([]byte("hello grep\n goodbye grep"))
	grepIn.Close()

	grepBytes, _ := ioutil.ReadAll(gerpOut)
	gerpCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	//execing-proc
	binary, lookErr := exec.LookPath("ls")
	checkError(lookErr)

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	checkError(execErr)
}
