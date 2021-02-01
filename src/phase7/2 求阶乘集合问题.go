package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[uint64]uint64, 10)
	/*
		加入全局互斥锁来解决并发写的问题（低端）
	*/
	lock sync.Mutex

	channel = make(chan uint64, 120)
)

/*
计算某一个数的阶乘
*/
func calFactorial(n uint64) {
	var res uint64 = 1
	var i uint64
	for i = 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func calFactorial2(n uint64) {
	var res uint64 = 1
	var i uint64
	for i = 1; i <= n; i++ {
		res *= i
	}
	channel <- res
}

func mutexWay() {
	var i uint64
	for i = 1; i <= 50; i++ {
		go calFactorial(i)
	}
	time.Sleep(time.Second * 5)

	fmt.Println(myMap)
}

func channelWay() {
	var i uint64
	for i = 1; i <= 50; i++ {
		go calFactorial2(i)
	}
	/*
		就是不能close channel！这意味着协程代码是异步运行的！
	*/
	//close(channel)

	m := make(map[int]uint64)
	for i := 1; i <= 50; i++ {
		m[i] = <-channel
	}
	fmt.Println(m)
}

func main() {
	channelWay()
}
