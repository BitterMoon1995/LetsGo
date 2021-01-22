package main

import "fmt"

/*
判断条件不用括号包裹
可以在判断条件中声明变量。注意：该变量为块级作用域，即仅在声明其的if语句中可以访问，而非所处函数（反观）
*/
func ifInGo() {
	flag := true
	if flag {
		fmt.Println("nigga")
	}

	if a := 34; a > 32 {
		fmt.Println("nigga~~~~~")
	}

	//fmt.Println(a)
}
func forInGo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//fmt.Println(i)
}
func whereIsMyWhile() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func infiniteLoop() {
	i := 0
	for {
		if i > 10 {
			break
		} else {
			i++
			fmt.Println(i)
		}
	}
}

/*
经典永流传之循环控制
*/
func breakAndContinue() {
	for i := 0; i < 20; i++ {
		fmt.Println(i)
		if i == 7 {
			fmt.Println("到7了")
			continue
		}
		if i == 18 {
			fmt.Println("可莉可莉")
			break
		}
	}
}

func main() {
	breakAndContinue()
}
