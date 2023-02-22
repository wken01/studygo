package main

import (
	"fmt"
	"os"
)

func main() {

	//Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
	//panic("a problem")

	//假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
	/*
			比如panic函数内有：
		    defer   函数1
		    defer   函数2
		    defer   函数3
			那么执行顺序就是：
		    函数3
		    函数2
		    函数1
	*/

	f1()
	f2()
}

func f1() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("文件不存在1")
		}
	}()
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

func f2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("文件不存在2")
		}
	}()
	_, err := os.Create("/tmp/file2")
	if err != nil {
		panic(err)
	}
}
