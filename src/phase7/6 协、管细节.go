package main

import (
	"fmt"
	"time"
)

/*
channel可以声明为只读或只写
应用：可在用在方法形参上，防止误操作
*/
func rwOnlyChannel() {
	var rc chan<- int
	rc = make(chan int, 3)
	rc <- 20
	//Invalid operation: <-rc (receive from send-only type chan<- int)
	//<-rc

	var wc <-chan int
	wc = make(chan int, 3)
	fmt.Println(wc)
	//wc<-5
}

/*
goroutine错误处理
*/
func recoverInGoroutine(bc chan bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("协程错误：", err)
		}
	}()

	var imap map[int]string
	imap[0] = "nigger"
	/*
		unreachable
	*/
	fmt.Println("协程结束")
}

func test2() {
	bc := make(chan bool, 10)
	go recoverInGoroutine(bc)
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
	}
	fmt.Println("主线程结束")
}

func main() {

}
