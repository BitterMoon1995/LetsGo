package main

import "fmt"

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
	ran := 8000

	srcChannel := make(chan int, 8000)
	go initSrc(ran, srcChannel)

	dstChannel := make(chan int, 1000)
	exitChannel := make(chan bool, 1)

	for i := 0; i < 4; i++ {
		go calPrime(srcChannel, dstChannel, exitChannel)
	}
	/*
		★★
	*/
	for i := 0; i < 4; i++ {
		<-exitChannel
	}
	close(exitChannel)
	close(dstChannel)

	for v := range dstChannel {
		fmt.Println(v)
	}
}
