package main

import "fmt"

func theForRange() {
	str := "你好nigga"
	for key, value := range str {
		fmt.Println("索引：", key, "值：", value)
	}
}

/*
多个满足条件时，可用逗号分割
golang中不需要写break,需要穿透时，分支语句最后写上fall through
*/
func theSwitchCase(s string) {
	switch s {
	case "nigga", "nigger", "尼哥":
		fmt.Println("尼哥黑鬼")
	case "whiteMan":
		fmt.Println("白人帅哥")
	case "whiteWoman":
		fmt.Println("睾贵白女")
	default:
		fmt.Println("疑似辱了")
	}
}

/*
case使用表达式,这是switch后不要跟判断变量
用fall through进行分支穿透，注意只能穿透一层
*/
func expressionCase(age int) {
	switch {
	case age < 14:
		fmt.Println("长长的路上我想")
		fallthrough
	case age >= 14 && age < 25:
		fmt.Println("如果沉默太沉重")
	default:
		fmt.Println("重国亩狗の支配")
	}
}

func main() {
	expressionCase(13)
}
