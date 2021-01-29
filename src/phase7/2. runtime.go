package main

import (
	"fmt"
	"runtime"
)

/*
获取以及设置CPU核心数
go1.8以后，默认情况下所有CPU都会用上；1.8以前要手动设置
*/
func cpuCores() {
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
	runtime.GOMAXPROCS(numCPU - 1) //比如想留一个运行其他程序
}

func main() {
	cpuCores()
}
