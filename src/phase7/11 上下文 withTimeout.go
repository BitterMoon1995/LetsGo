package main

import (
	"context"
	"fmt"
	"time"
)

/**
演示withTimeout
*/

func timeoutGo(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到信号，timeoutGo退出,time=", time.Now().Second())
			/*
				到时且上层任务还未取消：context deadline exceeded
			*/
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("timeoutGo 跑着", time.Now().Second())
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// 创建一个子节点的context,3秒后自动超时
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Second*3)

	go timeoutGo(timeoutContext)

	fmt.Println("现在开始等待8秒,time=", time.Now().Second())
	time.Sleep(8 * time.Second)

	fmt.Println("等待8秒结束,准备调用cancel()函数，发现timeoutGo已经结束了，time=", time.Now().Second())
	cancel()
}
