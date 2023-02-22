package main

import (
	"fmt"
	"strconv"
	s "strings"
)

var p = fmt.Println
var pf = fmt.Printf

func main() {

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

	//格式化
	var url = "Code=%s&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	p("string format: ", target_url)

	s := "hello,中国"
	p("len:       ", len(s)) //字符串的 ASCII 字符个数或字节长度

	r1 := []rune(s)
	pf("str to rune:%s, len(r1): %d \n", string(r1), len(r1))

	p("-----------------------number to str-----------------------")
	f, _ := strconv.ParseFloat("1.234", 64)
	p(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	p(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	p(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	p(u)

	k, _ := strconv.Atoi("135")
	p(k)

	_, e := strconv.Atoi("wat")
	p(e)
}
