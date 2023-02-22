package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	//2种方式都可以
	fmt.Println("describe:", co.base.describe())
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	//Embedding structs with methods may be used to bestow interface implementations onto other structs.
	//Here we see that a container now implements the describer interface because it embeds base.
	//内部类型提升到外部了
	var d describer = co
	fmt.Println("describer:", d.describe())
}
