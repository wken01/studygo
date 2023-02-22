package main

import "fmt"

func main() {

	//Channels are the pipes that connect concurrent goroutines.
	//You can send values into channels from one goroutine and receive those values into another goroutine.
	messages := make(chan string) //默认是无缓冲的，非缓冲 channel，channel 发送和接收动作是同时发生的，如果没有人接收，发送就一定会堵塞

	go func() { messages <- "ping" }()

	msg := <-messages
	//By default sends and receives block until both the sender and receiver are ready.
	//This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
	fmt.Println(msg)
}
