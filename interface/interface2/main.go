package main

import (
	"fmt"
)

type notifier interface {
	notify()
}
type user struct {
	name  string
	email string
}

type admin struct {
	*user
	level string
}

func (u *user) notify() {
	fmt.Printf("send user email %s, %s\n", u.name, u.email)
}

func (a *admin) notify() {
	fmt.Printf("send admin email %s, %s\n", a.name, a.email)
}

func sendNotifer(n notifier) {
	n.notify()
}

func main() {
	ad := admin{
		user: &user{
			name:  "tom",
			email: "1212@qq.com",
		},
		level: "super",
	}
	//接口是一种类型,&ad 就是实现了notifier的实现类
	fmt.Printf("&ad type is:%T\n", &ad)
	sendNotifer(&ad)
	ad.user.notify()
	ad.notify()
}
