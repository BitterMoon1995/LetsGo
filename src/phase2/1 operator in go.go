package main

import "fmt"

/*
golang中自增自减是单独的语句，不能作为赋值运算使用；
golang中没有前置的++或--
*/
func selfOperate() {
	i := 1
	i++
	//var j int = i++
	//--i
	fmt.Println(i)
}

func bitOperate() {
	a := 6
	fmt.Println(a << 2)
	fmt.Println(a >> 1)
}

func main() {
	bitOperate()
}
