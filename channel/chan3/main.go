package main

import "fmt"

//缓冲型的channel
func main() {
	intChan := make(chan int, 10)
	for i :=0; i<10;i++ {
		intChan <- i
	}
	close(intChan)
	syncChan := make(chan struct{}, 1)
	go func() {
		Loop: //和break Loop 对应，成对出现
			for{
				select {
					case e, ok := <- intChan:
						if !ok {
							fmt.Println("End")
							break Loop
						}
						fmt.Printf("Received: %v\n",e)
				}
			}
			syncChan <- struct{}{}
	}()
	<- syncChan
}
