package main

import (
	"fmt"
	"time"
)

/*
channel是引用类型
依然是先声明后make，或者声明同时make
★可以声明带缓冲或者不带缓冲的信道
*/
func declareChannel() {
	/*
		不带缓冲区的信道
		发送和接收操作在另一端准备好之前都会阻塞。
		这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

		这个例子中，首先明白永远是主线程先走，所以写入的操作一定要在主线程里进行；
		然后会发现主线程如果上来就睡的话，协程会被夹住（阻塞），
		而后主线程接收到第一个数据后，准备接收第二个数据时协程又睡了，主线程又会被阻塞，
		所以最好不要使用无缓冲信道
	*/
	chan2 := make(chan int)
	go func() {
		chan2 <- 2
		time.Sleep(time.Second)
		chan2 <- 3
	}()
	time.Sleep(time.Second)
	i := <-chan2
	fmt.Println(i)
	i = <-chan2
	fmt.Println(i)

	var chan1 chan int
	/*
		声明带缓冲的信道：
		当信道的缓冲区填满后，向其发送数据时会阻塞。当缓冲区为空时，接受方会阻塞。
	*/
	chan1 = make(chan int, 3)
	fmt.Println(chan1)
}

func noCacheChannel() {
	c := make(chan int) //声明一个int类型的无缓冲通道
	go func() {
		fmt.Println("ready to send in g1")
		c <- 1
		fmt.Println("send 1 to chan")
		fmt.Println("goroutine start sleep 1 second")
		time.Sleep(time.Second)
		fmt.Println("goroutine end sleep")
		c <- 2
		fmt.Println("send 2 to chan")
	}()

	fmt.Println("main thread start sleep 1 second")
	time.Sleep(time.Second)
	fmt.Println("main thread end sleep")
	i := <-c
	fmt.Printf("receive %d\n", i)
	i = <-c
	fmt.Printf("receive %d\n", i)
	time.Sleep(time.Second)
}

func cacheChannel() {
	c := make(chan int, 2) //声明一个int类型的有缓冲通道
	go func() {
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Printf("send %d\n", i)
		}
		time.Sleep(5 * time.Second)
		for i := 4; i < 6; i++ {
			c <- i
			fmt.Printf("send %d\n", i)
		}
	}()

	for i := 0; i < 6; i++ {
		time.Sleep(time.Second)
		v, ok := <-c
		fmt.Println(ok, v)
	}
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
	channel := make(chan int, 3)
	channel <- 211
	channel <- 985
	channel <- 911

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	/*
		从channel中取数据时可以再接收一个标志位参数
		从有数据的通道取值该标志位为true（不管关没关闭）；
		从已关闭且无数据的通道取值，该标志位为false，且取出的值为默认值（int就是0），注意这个值要丢弃；
		从未关闭且无数据的通道取值，该标志位没梗，只会直接报错，
		所以写入结束一定要关闭
	*/
	close(channel)
	v, ok := <-channel
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("啊这22222222", v)
	}
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
	readFromChannel()
}
