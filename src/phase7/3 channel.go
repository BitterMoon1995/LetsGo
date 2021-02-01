package main

import (
	"fmt"
)

/*
channel是引用类型
1.先声明后make
*/
func declareMap1() {
	var chan1 chan int
	chan1 = make(chan int, 3)
	fmt.Println(chan1)

}

func writeToChannel() {
	channel := make(chan int)
	num := 4396
	channel <- 1
	channel <- num
	channel <- 3
	//str:="死妈尼哥儿"
	//channel<-str
}

/*
channel的容量是不会动态增长的
*/
func lenAndCapOfMap() {
	channel := make(chan int)
	fmt.Println(len(channel), cap(channel))

	channel2 := make(chan int, 3)
	channel2 <- 1
	channel2 <- 2
	channel2 <- 3
	/*
		fatal error: all goroutines are asleep - deadlock!
	*/
	//channel2<-4
	fmt.Println(len(channel2), cap(channel2))
}

/*
管道的容量是固定不变的，它的价值在于边读边取，取了又放，流动起来
*/
func readFromChannel() {
	channel := make(chan int, 10)
	channel <- 211
	channel <- 985
	channel <- 911

	<-channel
	channel <- 721
	fmt.Println(<-channel)

	/*
		fatal error: all goroutines are asleep - deadlock!
		在没有使用协程的情况下，超量取又要麻
	*/
	<-channel
	<-channel
	//<-channel
}

type SweetGirl struct {
	name      string
	bestDress string
}

/*
interface{}类型的channel
*/
func interfaceType() {
	channel := make(chan interface{}, 10)
	g := SweetGirl{
		name:      "美妖",
		bestDress: "雨薇",
	}
	channel <- g
	channel <- "死妈尼哥儿"
	g1 := <-channel
	/*
		运行后发现可以推导出类型，但编译时认为是空接口类型，.name过不了编译
	*/
	fmt.Printf("%T %v", g1, g1)
	//fmt.Println(g1.name)
	g2 := g1.(SweetGirl)
	fmt.Println(g2.name)
}

/*
关闭channel，不能再写入，但可以读
*/
func closeChannel() {
	c := make(chan int, 10)
	c <- 985
	c <- 211
	close(c)
	/*
		panic: send on closed channel
	*/
	//c<-911
	i := <-c
	fmt.Println(i)
}

/*
遍历管道，不能使用for i，因为管道的len是动态改变的。必须使用for range（遍历的是值，channel是个队列没有索引）
遍历管道之前必须先关闭，否则deadlock
*/
func traverseChannel() {
	c := make(chan int, 20)
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	traverseChannel()
}
