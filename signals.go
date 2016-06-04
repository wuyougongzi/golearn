//have to be in bash termal neither can not run as except

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("goroutines in")
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true

	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
