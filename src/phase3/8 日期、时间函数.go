package main

import (
	"fmt"
	"reflect"
	"time"
)

func now() {
	now := time.Now()
	fmt.Println(now, reflect.TypeOf(now))
	fmt.Println("当前年：", now.Year())
	fmt.Println("当前月：", now.Month())
	fmt.Println("当前日：", now.Day())
	fmt.Println("当前时：", now.Hour())
	fmt.Println("当前分：", now.Minute())
	fmt.Println("当前秒：", now.Second())
	fmt.Println("当前毫秒：", now.UnixNano()/1e6)
}

/*
2006年1月2日下午3点4分5秒，我
*/
func dateFormat() {
	now := time.Now()
	formattedStr := now.Format("2006/01/02 15:04:05")
	fmt.Println(formattedStr)
	formattedDay := now.Format("2006/01/02")
	fmt.Println(formattedDay)
	formattedTime := now.Format("15:04:05")
	fmt.Println(formattedTime)
}

func constantOfTime() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
	}
}

func timeStamp() {
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano() / 1e6)
}

func main() {
	timeStamp()
}
