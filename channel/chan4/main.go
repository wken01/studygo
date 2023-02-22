package main

import (
	"fmt"
	"time"
)
//发送操作的循环由go函数包裹，所以它与
func main() {
	sendingInterval := time.Second * 4
	receptionInterval := time.Second * 2
	intChan := make(chan int, 0)
	go func() {
		var ts0,ts1 int64
		for i := 1;i<= 5;i++ {
			intChan <- i
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Send:", i)
			} else {
				fmt.Printf("Send: %d[interval: %ds]\n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sendingInterval)
		}
		close(intChan)
	}()
	syncChan := make(chan struct{}, 1)
	go func() {
		var ts0,ts1 int64
		Loop:
			for {
				select {
				case v, ok := <-intChan:
					if !ok {
						break Loop
					}
					ts1 = time.Now().Unix()
					if ts0 == 0 {
						fmt.Println("received:", v)
					} else {
						fmt.Printf("received: %d[interval: %ds]\n",  v, ts1-ts0)
					}
				}
				ts0 = time.Now().Unix()
				time.Sleep(receptionInterval)
			}
			fmt.Println("End.")
			syncChan <- struct{}{}
	}()
    <- syncChan
}
