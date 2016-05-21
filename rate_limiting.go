/*
rate limit process
https://gobyexample.com/rate-limiting

两种rate-limit的方式
1 直接限制每间隔200ms 处理一次请求
2 前面集中的处理3个，是将定时器缓存三个，后面每间隔间隔200ms再处理一个
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)

	//simulate 5 request
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	//after closed request can not receive message
	close(requests)

	//use tick to limit request process
	limits := time.Tick(time.Microsecond * 200)

	//prcess a request per 200ms
	for req := range requests {
		<-limits
		fmt.Println("request", req, time.Now())
	}

	//the other way
	burstyLimiter := make(chan time.Time, 3)

	//add 3 ticked in a short time
	for i := 1; i <= 2; i++ {
		burstyLimiter <- time.Now()
	}

	//burstLimiter add a tick per 200ms
	go func() {
		for t := range time.Tick(time.Microsecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("bursyRequests: ", req, time.Now())
	}
}
