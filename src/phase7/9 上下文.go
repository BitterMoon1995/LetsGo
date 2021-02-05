package main

import (
	"context"
	"fmt"
	"time"
)

/**
使用守则：
不要把Context放在结构体中，要以参数的方式进行传递
以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位
给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
*/

/*
单个
*/
func learn01() {
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

/*
多个
*/
func learn02() {
	//创建一个可取消子context,context.Background():返回一个空的Context，这个空的Context一般用于整个Context树的根节点。
	ctx, cancel := context.WithCancel(context.Background())
	ctxTwo, cancelTwo := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			//使用select调用<-ctx.Done()判断是否要结束
			case <-ctx.Done():
				fmt.Println("goroutineA exit")
				return
			default:
				fmt.Println("goroutineA running.")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			//使用select调用<-ctx.Done()判断是否要结束
			case <-ctx.Done():
				fmt.Println("goroutineB exit")
				return
			default:
				fmt.Println("goroutineB running.")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	go func(ctxTwo context.Context) {
		for {
			select {
			//使用select调用<-ctx.Done()判断是否要结束
			case <-ctxTwo.Done():
				fmt.Println("goroutineC exit")
				return
			default:
				fmt.Println("goroutineC running.")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctxTwo)

	time.Sleep(4 * time.Second)
	fmt.Println("main fun exit")
	//取消context
	cancel()
	cancelTwo()
	time.Sleep(5 * time.Second)
}

func main() {

}
