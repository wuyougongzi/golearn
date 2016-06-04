package main

import (
	"fmt"
	"time"
)

//no buffer chan use example
func chanShow() {
	//make chan
	//no buffer chan
	//receive or send info must have to goroutine or it will be deadlock
	message := make(chan string)
	go func() { message <- "ping" }()
	msg := <-message
	fmt.Println(msg)
}

//chan buffered example
func bufferChanShow() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*
通过代码现在猜测，go 的channel 作为参数时（就算是在外面创建的有缓冲区大小的channel），
传入时默认无缓冲区，会根据函数内部的写入缓冲区（即调用 channel <- data）
次数和外面调用的次数适应缓冲区的大小，且外面调用读（调用data <- channel）
的次数要不大于里面写的次数。不然会出现deadlock，所以该例子中func main 去掉 <- done 时，
调用worker看着像没有生效，是因为此时done是一个无缓冲区channel，
外面没有了读，worker里面的写就会被阻塞。
*/
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
	//done <- false
	//<-done
}

func channelDirection1(receiver chan<- string, msg string) {
	receiver <- msg
}

func channelDirection2(sender <-chan string, receiver chan<- string) {
	//sender <- <-receiver
	receiver <- <-sender
}

func main() {
	//	chanShow()
	//	bufferChanShow()

	//done is a chan type with a buffer,chan data type is bool

	/*	done := make(chan bool, 1)
		fmt.Println("work start:")
		go worker(done)
		fmt.Println("work end:")
		//must have next line code
		//because
		fmt.Println("done: ", <-done)
		//fmt.Println("done: ", <-done)
		//fmt.Println("done: ", <-done)
	*/

	pc := make(chan string, 1)
	qc := make(chan string, 1)
	channelDirection1(pc, "msg start")
	channelDirection2(pc, qc)

	//fmt.Println("pc: ", <-pc)
	fmt.Println("qc: ", <-qc)

	//t := 1
	//fmt.Println(t)
}
