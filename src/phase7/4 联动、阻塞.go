package main

import (
	"fmt"
	"time"
)

func w50(oc chan int) { //operating channel
	for i := 0; i < 50; i++ {
		oc <- i
		fmt.Println("写入：", i)
	}
	close(oc)
}
func readX(oc chan int, exitC chan bool) {
	for {
		v, ok := <-oc
		if !ok {
			fmt.Println("读结束")
			break
			/*
				标志位OK为true时，读出的值才有效，
				因为最后一次读，oc已空且关闭，ok=false，
				但是特别注意会读出一个值v=0，这个值是无效的。
				最后一次读不算，否则有无效数据且次数对不上
			*/
		} else {
			time.Sleep(time.Second)
			fmt.Println("读出：", v, ok)
		}
	}
	exitC <- true
	close(exitC)
}

func wr50() {
	oc := make(chan int, 1)
	exitC := make(chan bool, 1)
	go w50(oc)
	go readX(oc, exitC)
	for {
		v, ok := <-exitC
		if v || !ok {
			fmt.Println("读、写协程运行完毕")
			break
		}
	}
}

func nigger() {
	/*
		不需要把channel的缓冲大小设置为需要处理的总数据大小，大多数情况这也是不可能的。
		只要能保证channel有接收方且在正常接收，就不会死锁，即使读、写速度不一致也无所谓;
		而如果只写不读，就会在信道的缓冲区填满后，发送端阻塞。
	*/
	c := make(chan int, 7)
	c <- 2
	<-c
	close(c)
	i, ok := <-c
	fmt.Println(i, ok)

}

/*
协程+通道结合猴
开启一个writeChannel，向一个intChan中写入50个整数；
再开一个readChannel从中读，要求主线程在二者都完成后才能退出。

关键点：
1.再开一个协程exit channel，用以存储读取成功后的标志位
2.利用好读协程的第二个标志位返回值在读已close，
*/
func main() {
	wr50()
}
