package main

import (
	"fmt"
	//"utils"
)

//var a = globalVariableInit()
//
//func globalVariableInit() int {
//	fmt.Println("全局变量先初始化")
//	return 10
//}
//
func init() {
	fmt.Println("init() 函数随后执行")
}

func main() {
	fmt.Println("main() 函数最后执行")
	//fmt.Println(utils.Name, utils.Age)
}
