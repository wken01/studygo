package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var chinese = "我是中国人， I am Chinese"
	fmt.Println("chinese length", len(chinese))
	fmt.Println("chinese word length", len([]rune(chinese)))
	fmt.Println("chinese word length", utf8.RuneCountInString(chinese))

	//一个汉字是3个字符
	const s = "hello,世界"

	fmt.Println("Len:", len(s))
	fmt.Println("s:", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	//解码第一个字符，返回字符和该字符的字节长度
	c, i := utf8.DecodeRuneInString("Hello")
	fmt.Printf("ASCII 值： %d, %c，字节数：%d\n", c, c, i)

	fmt.Println()
	fmt.Println("====================runes========================")

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	fmt.Println("Rune s:", utf8.FullRuneInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d, width: %d\n", runeValue, i, width)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
