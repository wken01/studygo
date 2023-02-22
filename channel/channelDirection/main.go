package main

import "fmt"

//当使用通道作为函数参数时，您可以指定通道是否仅用于发送或接收值。这种特异性提高了程序的类型安全性。

//发送值
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//接收值
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
