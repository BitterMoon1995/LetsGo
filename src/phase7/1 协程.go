package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
并发(concurrency)：指在同一时刻只能有一条指令执行，但多个进程指令被快速的轮换执行，
使得在宏观上具有多个进程同时执行的效果，但在微观上并不是同时执行的，
只是把时间分成若干段，使多个进程快速交替的执行。

并行(parallel)：指在同一时刻，有多条指令在多个处理器上同时执行。

所以把多个微服务部署在单机上是有意义的！微服务数等于处理器核心数就挺好。


go协程特点：
1.有独立的栈空间
2.共享程序堆空间
3.调度由用户控制
4.go没有传统线程，协程类似线程，但是是轻量级的线程
★5.go关键字修饰的函数调用（协程代码）是异步进行的，非同步！最好做好回调或者利用缓冲容器
6.协程函数不要有返回值，加了go关键字接收不了

go主进程/线程：物理意义上的内核线程
*/
func test() {
	for i := 0; i < 100; i++ {
		fmt.Println("协程杀尼哥", i)
		time.Sleep(time.Second)
	}
}

/*
获取以及设置CPU核心数
go1.8以后，默认情况下所有CPU都会用上；1.8以前要手动设置
*/
func cpuCores() {
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
	runtime.GOMAXPROCS(numCPU - 1) //比如想留一个运行其他程序
}

/*
以主线程为主，主线程任务执行完，即使协程还没有执行完，也一并全部退出
*/
func main() {
	/*
		主线程(go进程)执行到这一行时，开启协程并执行test()
		lets go！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！
	*/
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("主线程杀尼哥", i)
		time.Sleep(time.Second)
	}
}
