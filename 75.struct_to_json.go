package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	Age  int    `json:"name"`
	Name string `json:"name"`
}

func main() {
	person := Person{Age: 20, Name: "zlw"}
	pv := reflect.ValueOf(person)
	pt := reflect.TypeOf(person)
	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")

	for i := 0; i < pt.NumField(); i++ {
		// 获取json tag
		jsonTag := pt.Field(i).Tag.Get("json")
		jsonBuilder.WriteString("\"" + jsonTag + "\"")
		jsonBuilder.WriteString(":")
		// 获取字段的值
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", pv.Field(i)))
		if i < pt.NumField()-1 {
			jsonBuilder.WriteString(",")
		}
	}

	jsonBuilder.WriteString("}")
	fmt.Println(jsonBuilder.String())
}
