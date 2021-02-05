package main

import (
	"context"
	"fmt"
	"time"
)

/**
似乎就是官方版done channel
(WithCancel returns a copy of parent with a new Done channel. The returned
context's Done channel is closed when the returned cancel function is called
or when the parent context's Done channel is closed, whichever happens first.

Canceling this context releases resources associated with it, so code should
call cancel as soon as the operations running in this Context complete.)

简单模拟：主线程下有两个协程01 02，01是context带值的02是不带值的，01下还有一个子协程011，
全部都是with cancel

总书记批示：子任务如何知道自己的父任务是谁？
*/
func goroutine01(ctx context.Context) {
	/*
		The cancel function should be called, not discarded, to avoid a context leak
	*/
	ctx2, cancel2 := context.WithCancel(ctx)
	go goroutine011(ctx2)
	for {
		select {
		/*
			Done is provided for use in select statements:
		*/
		case <-ctx.Done():
			/*
				(If Done is not yet closed, Err returns nil.
				If Done is closed, Err returns a non-nil error explaining why)

				上层任务被取消：context canceled
			*/
			fmt.Println(ctx.Err())
			fmt.Println("01接收到主线程信号，01准备退出")
			cancel2()
			fmt.Println("01已关闭011")
			return
		default:
			fmt.Println("goroutine01 running")
			time.Sleep(time.Second)
		}
	}
}

func goroutine011(ctx context.Context) {
	fmt.Println("011上下文获取到value：", ctx.Value("nigger"))
	for {
		select {
		case <-ctx.Done():
			fmt.Println("011接收到上层任务信号，退出")
			return
		default:
			fmt.Println("goroutine011 running")
			time.Sleep(time.Second)
		}
	}
}

func goroutine02(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("02接收到主线程信号，退出")
			return
		default:
			fmt.Println("goroutine02 running")
			time.Sleep(time.Second)
		}
	}
}

func test011() {
	/*

	 */
	context1, cancel1 := context.WithCancel(context.Background())
	context2, cancel2 := context.WithCancel(context.Background())
	/*
		Use context Values only for request-scoped data that transits processes and
		APIs, not for passing optional parameters to functions.
	*/
	context1wv := context.WithValue(context1, "nigger", "nigger")

	go goroutine01(context1wv)
	go goroutine02(context2)
	time.Sleep(3 * time.Second)
	cancel1()
	fmt.Println("主线程关闭context1wv")
	time.Sleep(3 * time.Second)
	cancel2()
	fmt.Println("主线程关闭context2")
	time.Sleep(2 * time.Second)
	fmt.Println("主线程关闭")
}
func main() {
	test011()
}
