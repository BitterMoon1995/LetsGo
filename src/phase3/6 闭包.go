package main

import "fmt"

/**
闭包就是一个函数与其相关的引用环境组合的一个整体（实体）

指的就是：
	n := 10
	return func(a int) int {
		n += a
		return n
	}
函数和其引用的变量共同构成闭包

Java中的【类】就是一个闭包，
类比此处，n就是类的属性（字段），返回的func就是类的方法（操作）
*/

/*累加器*/
func Accumulator() func(int) int {
	n := 10
	return func(a int) int {
		n += a
		return n
	}
}

func main() {
	accumulator := Accumulator()
	fmt.Println(accumulator(1))
	fmt.Println(accumulator(1))
	fmt.Println(accumulator(1))
}
