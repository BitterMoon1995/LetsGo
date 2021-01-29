package main

import "fmt"

var (
	myMap = make(map[int]int, 10)
)

/*
编写一个函数启动协程来计算各个数的阶乘，并放入到一个map中
@author:god周周神原创
*/
func calFactorial(n int) map[int]int {
	for i := 1; i <= n; i++ {
		res := 1
		for j := 1; j <= i; j++ {
			res *= j
		}
		myMap[i] = res
	}
	return myMap
}
func main() {
	m := calFactorial(5)
	fmt.Println(m)
}
