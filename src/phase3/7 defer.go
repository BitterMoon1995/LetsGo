package main

import "fmt"

/**
在函数中，经常需要创建资源，为了在函数执行完毕后及时释放资源，Go的设计者提供延时机制

当执行到defer修饰的语句时，不执行，而是将其压入独立的栈
当整个函数执行完毕后，再从defer栈中弹出语句，依次执行
*/
func sum(a int, b int) int {
	defer fmt.Println("a=", a)
	defer fmt.Println("b=", b)
	fmt.Println("sum=", a+b)
	return a + b
}

/*
在defer将语句放入到栈时，也会将相关的值拷贝同时入栈
*/
func selfIncr(a int) {
	defer fmt.Println(a)
	a++
}

func main() {
	a := 1
	selfIncr(a)
}
