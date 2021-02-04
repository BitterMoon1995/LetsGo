package main

import (
	"fmt"
	"time"
)

func initSrc(n int, src chan int) {
	for i := 0; i < n; i++ {
		src <- i
	}
	close(src)
}

func calPrime(src chan int, dst chan int, exit chan bool) {
	for {
		v, f := <-src
		if !f {
			break
		} else {
			flag := true
			for i := 2; i < v-1; i++ {
				if v%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				dst <- v
			}
		}
	}
	/*
		这个协程完成了，此时不要关闭任何信道因为其他协程还要用，给exitC写入一个标志位即可
	*/
	exit <- true
	fmt.Println("一个协程退出")
}

func main() {
	ran := 48000
	/*
		CPU密集型，并行数=CPU核心数即可
	*/
	parallelNum := 12
	srcChannel := make(chan int, 8000)
	dstChannel := make(chan int, 1000)
	exitChannel := make(chan bool, 1)
	start := time.Now().UnixNano()

	go initSrc(ran, srcChannel)

	for i := 0; i < parallelNum; i++ {
		go calPrime(srcChannel, dstChannel, exitChannel)
	}
	/*
		★★★
		主协程任务退出逻辑，
		盘点十大exit逻辑放在协程的理由：
		因为这个循环就是很直接地 <-exitChannel 四次，
		而<-exitChannel不安全，exitChannel为空会死锁，
		特别是本程序中等待 exit <- true 是很漫长的，
		所以这个逻辑执行的优先级一定不能高，不能放在主函数里面，必须是协程代码
	*/
	go func() {
		for i := 0; i < parallelNum; i++ {
			<-exitChannel
		}
		close(exitChannel)
		close(dstChannel)
		end := time.Now().UnixNano()
		fmt.Println("耗费毫秒：", (end-start)/1000)
	}()

	/*
		主线程延时逻辑，没有的话协程没有运行时间
	*/
	for {
		_, f := <-dstChannel
		if !f {
			break
		} else {
			//fmt.Println(v)
		}
	}
	fmt.Println("主线程结束")
}
