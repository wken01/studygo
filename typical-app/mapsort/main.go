package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//对map 的key进行排序，然后根据按照key顺序进行答应value
//思路：先初始化map，然后存入一个slice，然后对slice中的值进行排序，然后再遍历slice，按key的顺序遍历map，获取map值
//备注：对map、slice综合应用
func main() {
	m := make(map[string]int, 100)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		m[key] = value
	}

	keys := make([]string, 0, 200)
	for k := range m {
		keys = append(keys, k)
	}

	//对string进行排序
	//如何降序呢
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v, m[v])
	}

}
