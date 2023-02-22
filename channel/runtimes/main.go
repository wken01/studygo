package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// Our two function calls are running asynchronously in separate goroutines now. Wait for them to finish (for a more robust approach, use a WaitGroup).
func main() {

	f("direct")

	go f("goroutine")

	//匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
