package main

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// SJSON 和 GJSON
const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Get(json, "name.last")
	println(value.String())

	value1, _ := sjson.Set(json, "name.last", "Anderson")
	println(value1)
}
