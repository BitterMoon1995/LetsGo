package main

import (
	"fmt"
	"time"
)

/*
select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。
每个 case 必须是一个通信操作，要么是发送要么是接收。
select 随机执行一个可运行的 case，而如果没有 case 可运行，那么：
有 default 子句，则执行该语句；
没有 default 子句，select 将阻塞，直到某个case语句可以执行。

跳出select方式：return；	switch label；	goto（麻了）。
*/
func go01(intChan chan int, stopCh chan bool) {
	var i int
	i = 4396
	for j := 0; j < 50; j++ {
		intChan <- i
		time.Sleep(time.Millisecond * 20)
	}
	stopCh <- true
}

func testSelect() {
	intChan := make(chan int)
	stopChan := make(chan bool)

	go go01(intChan, stopChan)

	for {
		select {
		case i := <-intChan:
			fmt.Println("Receive", i)
		case _ = <-stopChan:
			fmt.Println("结束")
			goto end
		}
	}
end:
}

func main() {
	testSelect()
}
