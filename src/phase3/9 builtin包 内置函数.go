package main

import (
	"fmt"
	"reflect"
)

func newFunc() {
	p := new(int)
	fmt.Println("值：", p, "类型：", reflect.TypeOf(p), "地址：", &p, "指向：", *p)
}

func main() {
	newFunc()
}
