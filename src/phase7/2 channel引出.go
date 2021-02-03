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
	var i uint64 = 1
	for ; i <= n; i++ {

		var res uint64 = 1
		var j uint64 = 1
		for ; j <= i; j++ {
			res *= j
		}
		channel <- res
	}
	close(channel)
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
	calFactorial2(50)
	var i uint64 = 1
	for {
		v, f := <-channel
		if !f {
			break
		} else {
			myMap[i] = v
			i++
		}
	}
	fmt.Println(myMap)
}

func main() {
	channelWay()
}
