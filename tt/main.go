package main

import (
	"fmt"
)

func modifyArr(arr [2]int) {
	arr[0] = 100
}

func setUser(user *User, _user User) {
	user.Name = _user.Name
}

func main() {
	var i int = 10
	fmt.Printf("%d \n", i)
	fmt.Printf("%b \n", i)

	var b int = 077
	fmt.Printf("%o \n", b)

	var c int = 0xff
	fmt.Printf("%x \n", c)

	//数组
	var arr = [...]int{1, 2}
	fmt.Println(arr, len(arr))

	arr[0] = 3
	fmt.Println(arr, len(arr))
	fmt.Printf("%T \n", arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], ",")
	}

	for index, v := range arr {
		fmt.Println(index, v)
	}

	modifyArr(arr)
	fmt.Println("after modified:", arr)

	//切片
	a := []int{1, 2}
	fmt.Printf("%T \n", a)
	t := append(a, 3)
	fmt.Println(t)

	a1 := make([]int, 2, 5)
	a1[0] = 1
	a1[1] = 2
	fmt.Println("a1 slice:", a1)

	//map
	m := make(map[string]int, 8)
	m["a1"] = 0
	m["a2"] = 1
	fmt.Println("map get value: ", m["a2"])

	//map 的初始化
	userInfo := map[string]string{
		"username": "yxm",
		"age":      "100",
	}

	fmt.Println(userInfo)

	v, ok := userInfo["username"]
	if ok {
		fmt.Println("cunzai", v)
	} else {
		fmt.Println("no cunzai")
	}

	//map 的遍历
	for k, v := range userInfo {
		fmt.Printf("key: %v, value: %v \n", k, v)
	}

	fmt.Println("-----------------range point")
	type User struct {
		Name string
		Age  int
	}

	type User1 struct {
		Name string
		Age  int
	}

	type User2 struct {
		Name string
		Age  int
	}
	userList := []User{
		User{Name: "aa", Age: 1},
		User{Name: "bb", Age: 1},
	}

	var newUser []*User
	//按照正常的理解，应该第一次输出aa，第二次输出bb，但实际上两次都输出了bb，这是因为 for range 的时候，
	//变量u实际上只初始化了一次（每次遍历的时候u都会被重新赋值，但是地址不变），导致每次append的时候，添加的都是同一个内存地址，所以最终指向的都是最后一个值bb。
	//u的地址一样，而值不一样
	// for _, u := range userList {
	// 	newUser = append(newUser, &u)
	// }

	//修改1
	for i := 0; i < len(userList); i++ {
		u := userList[i]
		newUser = append(newUser, &u)
	}

	//修改2
	var newUser1 []*User1
	for _, u := range userList {
		user1 := User1{Name: u.Name}
		newUser1 = append(newUser1, &user1)
	}

	//修改3

	// 第一次：bb
	// 第二次：bb
	for _, nu := range newUser {
		fmt.Println()
		fmt.Printf("%+v", nu.Name)
	}

	for _, nu := range newUser1 {
		fmt.Println()
		fmt.Printf("%+v", nu.Name)
	}
	fmt.Println()
}
