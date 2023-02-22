package main

import (
	"fmt"
)

type args struct {
	nums []int
}

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func add(nums ...int) int {
	total := 0
	//for _, num := range nums {
	for i, num := range nums {
		fmt.Println("index:", i)
		total += num
	}
	return total
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	total := add(1, 2)
	fmt.Println(total)

}
