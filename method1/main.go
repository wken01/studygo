package main

import "fmt"

type Rect struct {
	width, height int
}

//那么如何在 Pointer 和 Value 接收器之间进行选择呢？

/*
	（1）如果接收者是 map、func 或 chan，不要使用指向它的指针。
	（2）尽量对所有方法使用相同的接收器类型。
	（3）如果接收者是一个切片并且该方法没有重新切片或重新分配切片，则不要使用指向它的指针。
	（4）如果方法需要改变接收者，接收者必须是一个指针。
	（5）如果接收者是包含sync.Mutex或类似同步字段的结构，则接收者必须是指针以避免复制。
	（6）如果接收器是大型结构或数组，则指针接收器效率更高。大有多大？假设它相当于将其所有元素作为参数传递给方法。如果感觉太大，那么对于接收器来说也太大了。
	（7）函数或方法是否可以同时或在从此方法调用时改变接收者？调用方法时，值类型会创建接收器的副本，因此外部更新不会应用于此接收器。如果更改必须在原始接收器中可见，则接收器必须是指针。
	（8）如果接收器是结构体、数组或切片，并且它的任何元素都是指向可能发生变化的东西的指针，则更喜欢指针接收器，因为它会使读者更清楚意图。
	（9）如果接收者是一个小数组或结构，它自然是一个值类型（例如，类似time.Time类型），没有可变字段和指针，或者只是一个简单的基本类型，如 int 或 string，则值接收器更好。
	（10）值接收器可以减少可以生成的垃圾量；如果将值传递给值方法，则可以使用堆栈上的副本而不是在堆上分配。（编译器试图巧妙地避免这种分配，但它并不总是成功。）不要在没有首先进行分析的情况下选择值接收器类型。
	（11）最后，当有疑问时，使用指针接收器。
*/
func (r *Rect) area() int {
	return r.width * r.height
}

func (r Rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := Rect{width: 10, height: 5}

	fmt.Printf("r=%#v\n", r)
	fmt.Printf("%T\n", r) //r的类型

	//用结构类型接收
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r //相当于对结构体new了，分配了内存空间

	fmt.Printf("rp=%#v\n", rp)
	fmt.Printf("%T\n", rp) //rp的类型
	//用结构指针来接收
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
