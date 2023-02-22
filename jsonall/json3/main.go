package main

import (
	"log"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

/*
普通 json 使用该方式处理确实可行，但是如果嵌套 json，也可以这么处理，但是读取嵌套 json 的子字段就不那么方便了。

我们再构造一个三方接口返回数据是嵌套 json 的变量。
*/
// func main() {
// 	// 三方返回嵌套 json 字符串
// 	jsonRes := `{
// 	"Id": 1001,
// 	"Name": "frank",
// 	"Details": {
// 		"Gender": "man",
// 		"Age": 18,
// 		"Phone": "13800138000",
// 		"address": "Beijing"
// 		}
// 	}`
// 	data := new(User)
// 	err := json.Unmarshal([]byte(jsonRes), &data)
// 	if err != nil {
// 	 log.Printf("json Unmarshal err:%v\n", err)
// 	 return
// 	}
// 	fmt.Printf("data=%+v", data)
//    }
   
//    type User struct {
// 	Id      int
// 	Name    string
// 	Details Details
//    }
   
//    type Details struct {
// 	Gender  string
// 	Age     int
// 	Phone   string
// 	Address string
//    }

   //但是，如果返回数据中的 Age 字段是字符串类型，我们使用 Unmarshal 函数解码时，就会返回以下错误：

   /*
   虽然，我们可以使用普通 json 中的处理方式，将返回数据解码到 map[string]interface{} 类型的变量中。
   但是，如果我们想要读取内嵌 json 中的子字段，就读取不到了。
   怎么解决这个问题呢？我们可以借助三方库 mapstructure，
   使用该三方库的 Decode 函数替代 Go 标准库 encoding/json 的 Unmarshal 函数。
   */

   func main() {
	// 三方返回嵌套 json 字符串
	jsonRes := `{
	"Id": 1001,
	"Name": "frank",
	"Details": {
		"Gender": "man",
		"Age": "18",
		"Phone": "13800138000",
		"address": "Beijing"
		}
	}`
	tmpData := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonRes), &tmpData)
	if err != nil {
	 log.Printf("json Unmarshal err:%v\n", err)
	 return
	}
	data2 := new(User)
	err = mapstructure.Decode(tmpData, data2)
	if err != nil {
	 log.Printf("decode err:%v\n", err)
	 return
	}
	fmt.Printf("data2=%+v\n", data2)
	fmt.Printf("age=%v\n", data2.Details.Age)
   }
   
   type User struct {
	Id      int
	Name    string
	Details Details
   }
   
   type Details struct {
	Gender  string
	Age     interface{}
	Phone   string
	Address string
   }