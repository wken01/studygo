package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {

	p := Person{name: name}
	p.age = 42
	return &p
}

func setPerson(person Person) {
	person.age = 15
}

func setPerson2(person *Person) {
	person.age = 15
}

func main() {

	fmt.Println(Person{"Bob", 20})

	fmt.Println(Person{name: "Alice", age: 30})

	fmt.Println(Person{name: "Fred"})

	fmt.Println(&Person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	s.age = 51
	fmt.Println(s.age)

	fmt.Println("无指针：", s)

	sp := &s
	fmt.Println("有指针：", sp)

	sp.age = 51
	fmt.Println("有指针改值：", sp.age)

	setPerson(s)
	fmt.Println("函数中，无指针，修改值： ", s)

	setPerson2(sp)
	sp.age = 52
	fmt.Println("函数中，有指针，修改值： ", sp)
}
