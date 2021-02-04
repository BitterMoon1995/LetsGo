package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
★★并发控制
理念：
就是控制流（主线程）牢牢地控制逻辑流（协程）
期望：
主线程不能阻塞、更不能搁这儿计时，0但是又能给协程恰好的运行时间；能开多个协程；协程还能开子协程
实现方式：
1.全局变量：放弃，多个协程还要加同步锁干寄吧呢
2.done channel
3.channel select（6有演示）：优雅，但多协程、子协程不太方便
4.waitgroup
5.context
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
	//创建一个可取消子context,context.Background():
	//返回一个空的Context，这个空的Context一般用于整个Context树的根节点。
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			//使用select调用<-ctx.Done()判断是否要结束
			case <-ctx.Done():
				fmt.Println("goroutine exit")
				return
			default:
				fmt.Println("goroutine running.")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	//取消context
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("main fun exit")
}
