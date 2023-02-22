package main

import "fmt"

func main() {

	//有缓冲的channel
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	//因为这个通道是缓冲的，所以我们可以将这些值发送到通道中，而无需相应的并发接收。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
