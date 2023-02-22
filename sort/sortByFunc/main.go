package main

import (
	"fmt"
	"sort"
)

type byLength []string

// 三个接口实现
// func Sort(data Interface)
/*
	type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}*/
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
