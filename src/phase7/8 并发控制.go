package main

import (
	"fmt"
	"sync"
	"time"
)

/**
★★并发控制
理念：
就是控制流（主线程）能够牢牢地控制逻辑流（协程）及其子逻辑流（子协程）
期望：
主线程有多个子任务，子任务还可以有多个子任务；
子任务既要感知主线程的取消信号，也需要感知上层任务的取消信号。
上层任务取消后，所有的下层任务都会被取消；
中间某一层的任务取消后，只会将当前任务的下层任务取消，而不会影响上层的任务以及同级任务。
实现方式：
1.全局变量：放弃，多个协程还要加同步锁干寄吧呢
2.done channel
3.channel select（6有演示）：优雅，但多协程、子协程不太方便
4.waitgroup
5.context
*/

/*
让子协程监听这个done channel，一旦主协程关闭done channel，那么子协程就可以推出了，这样就实现了主协程通知子协程的需求。
*/
func byDoneChannel() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("goroutine exit")
				return
			default:
				fmt.Println("goroutine running")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	stop <- true
	time.Sleep(2 * time.Second)
	fmt.Println("main fun exit")
}

func byWaitgroup() {
	//定义一个WaitGroup
	var wg sync.WaitGroup
	//计数器设置为2
	wg.Add(2)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("goroutineA finish")
		//计数器减1
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("goroutineB finish")
		//计数器减1
		wg.Done()
	}()
	//会阻塞代码的运行，直到计数器地值减为0。
	wg.Wait()
	fmt.Println("main fun exit")
}

func main() {
	byDoneChannel()
}
