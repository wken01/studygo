package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	p1 := &point{
		x: 1,
		y: 2,
	}

	fmt.Println(p1)
	//值
	fmt.Printf("struct1: %v\n", p)
	//带属性的值
	fmt.Printf("struct2: %+v\n", p)

	fmt.Printf("struct3: %#v\n", p)
	//To print the type of a value, use %T.
	fmt.Printf("type: %T\n", p)
	//Formatting booleans is straight-forward
	fmt.Printf("bool: %t\n", true)
	//There are many options for formatting integers. Use %d for standard, base-10 formatting.
	fmt.Printf("int: %d\n", 123)
	//This prints a binary representation.
	fmt.Printf("bin: %b\n", 14)
	//This prints the character corresponding to the given integer. 可参考ascii码表
	fmt.Printf("char: %c\n", 33)
	//%x provides hex encoding. 16进制
	fmt.Printf("hex: %x\n", 456)

	fmt.Printf("float1: %f\n", 78.9)

	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"")

	//To double-quote strings as in Go source, use %q.
	fmt.Printf("str2: %q\n", "\"string\"")

	fmt.Printf("str3: %x\n", "hex this")

	fmt.Printf("pointer: %p\n", &p)

	//指定宽度
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	//指定宽度。靠右
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	//指定宽度。靠左
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	//格式化
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
