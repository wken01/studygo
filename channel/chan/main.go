package main

import (
	"fmt"
	"time"
)

// 2个goroutine之间的协作
var strChan = make(chan string, 3)
func main() {
	syncChan1 := make(chan struct{},1)
	syncChan2 := make(chan struct{},2)
	//第一个goroutine
	go func() {//接收操作
		<-syncChan1
		fmt.Println("Received a sync sygnal and wait a second ...[receiver]")
		time.Sleep(time.Second)
		for{
			if elem, ok := <-strChan;ok {  //strChan接收值 ,chan变量在 <- 右边代表接收值
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("stopped .[receiver]")
		//（1）为了主goroutine不要过早结束 （2）struct{} 空结构体，不占用内存空间，且有相同的内存地址，建议用于传递信号的通道都以struct{}作为元素类型
		syncChan2 <- struct{}{}
	}()
	//第二个goroutine
	go func() {//发送操作
		for _,elem := range []string{"a","b","c","d"} { //遍历切片
			strChan <- elem //发送给strChan ,chan变量在 <- 左边代表发送给变量
			fmt.Println("Send", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Send a sync singal.[sender]")
			}
		}
		fmt.Println("Wait 2 seconds...[sender]")
		time.Sleep(time.Second*2)
		close(strChan)
		//（1）为了主goroutine不要过早结束（2）struct{} 空结构体，不占用内存空间，且有相同的内存地址，建议用于传递信号的通道都以struct{}作为元素类型
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}


/**-------结果
	Send a [sender]
	Send b [sender]
	Send c [sender]
	Send a sync singal.[sender]
	Received a sync sygnal and wait a second ...[receiver]
	Received: a [receiver]
	Received: b [receiver]
	Received: c [receiver]
	Received: d [receiver]
	Send d [sender]
	Wait 2 seconds...[sender]
	stopped .[receiver]
 */