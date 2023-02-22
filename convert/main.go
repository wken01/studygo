package main

import (
	"fmt"
	"strconv"
)

//各种基本类型之间的转换
func main() {
	//(1)整形和字符串转换
	var i int = 10
	s := strconv.Itoa(i)
	fmt.Printf("整形转字符串: %s \n", s)

	//(2)字符串转整形
	i, err := strconv.Atoi("10")
	if err != nil {
		fmt.Printf("字符串转整形： %d, %v \n", i, err)
	} else {
		fmt.Printf("字符串转整形： %d \n", i)
	}

	// 进制之间的转换
	//（1）十进制（整型）转二进制、八进制、十六进制（字符型）
	// strconv.FormatInt(i int64，base int），接收两个参数
	// i：需要进行转换的数据，类型为int64, base : 需要转换的进制 ，类型为int
	fmt.Printf("10进制 to 16进制: %v \n", strconv.FormatInt(int64(i), 16))
	fmt.Printf("10进制 to 8进制: %v \n", strconv.FormatInt(int64(i), 8))

	//(2) 二进制、八进制、十六进制（字符型） 转 十进制（整型）
	/*
					strconv.ParseInt(s string, base int, i int):
			        接收三个参数：
		                s：需要转换为十进制即整型（int64）的数据，类型需为字符串
		                base: 将s当作什么进制转换为十进制，如2、8、16，类型需为int
		                i：转换的过程中对结果进行约束，类型需为int
			        返回值：
		                result：成功即为转换后的十进制整型数据（int64），失败即为0
		                err : 成功为空，失败为提示信息
	*/
	j, err1 := strconv.ParseInt("11001", 2, 64)
	if err1 != nil {
		fmt.Printf("2进制 to 10进制: %v, %v \n", j, err1)
	} else {
		fmt.Printf("2进制 to 10进制: %v \n", j)
	}

}
