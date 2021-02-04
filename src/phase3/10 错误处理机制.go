package main

import (
	"errors"
	"fmt"
)

/*
Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效
defer + recover
捕获错误，终止程序，但可以成功返回
*/
func deferAndRecover() int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获到错误")
			fmt.Println(err)
		}
	}()

	a := 1
	b := 0
	c := a / b
	fmt.Println("a/b后续代码")
	return c
}

/*
errors.New自定义错误
panic捕获错误，输出错误，终止程序，不返回
*/
func diyError(fileName string) (err error) {
	if fileName == "nigga.avi" {
		return nil
	} else {
		return errors.New("文件名错误")
	}
}

func diyErrorT() {
	err := diyError("nigga.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("调用后代码")
}

func main() {
	diyErrorT()
	fmt.Println("主函数后续代码")
}
