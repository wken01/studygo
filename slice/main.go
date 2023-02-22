package main

/*
	int add(int a,int b) {
    	return a+b;
	}
*/
import "fmt"

func main() {
	a1 := [...]int{1, 3, 5, 7, 9, 11, 15, 17}
	s1 := a1[:]
	fmt.Println("a1[:]--->", s1)
	s1 = append(s1[0:1], s1[2:5]...)
	fmt.Println(s1)
	fmt.Println(a1)

	//fmt.Println(C.add(1+2))

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置

	fmt.Println(slice1)

	c := 3
	fmt.Println(c << 2)
}
