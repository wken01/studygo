package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
我们使用 Go 标准库 encoding/json 的 Unmarshal 函数，可以很容易将 json 数据解码到 struct，从而方便我们读取返回数据。
但是，需要读者朋友们注意的是，假如三方接口返回数据的字段类型随机变化（比如示例中的 Id 字段，可能是整型或字符串随机返回），
我们使用 Unmarshal 函数解码时，就有可能会返回错误
解决办法：以把返回数据解码到 map[string]interface{} 类型的变量中
*/
func main() {
	// 三方返回普通 json 字符串
	jsonRes := `{
	 "Id": 1001,
	 "Name": "frank"
	}`
	data2 := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonRes), &data2)
	if err != nil {
		log.Printf("json Unmarshal err:%v\n", err)
		return
	}
	fmt.Printf("data2=%+v", data2)
}
