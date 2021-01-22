package main

import (
	"fmt"
	"reflect"
	"utils"
)

var F2 = func(a int, b int) int {
	return a * b
}

func main() {
	/*在定义匿名函数时就直接调用，这样的匿名函数只能调用一次*/
	s := func() string { return "nigga" }()
	fmt.Println(s, "类型：", reflect.TypeOf(s))

	/*定义匿名函数后不调用，使用一个变量来接收，
	那么这个变量就变成了该函数类型，可以像函数一样调用它*/
	sum := func(a int, b int) int {
		return a + b
	}
	fmt.Println(sum(3, 4), "类型：", reflect.TypeOf(sum))

	res := utils.F(88, 99)
	fmt.Println(res, reflect.TypeOf(utils.F))
}
