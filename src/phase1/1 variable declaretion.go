package main

import (
	"fmt"
)

/*
在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的。例如，Pizza 就是个已导出名，Pi 也同样，它导出自 math 包。
pizza 和 pi 并未以大写字母开头，所以它们是未导出的。
在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。

三种变量声明写法
*/
func ThreeKinds() {
	/*盘点十大变量初始化方式：先声明再赋值、声明时赋值、声明时自动推导*/
	var username string
	username = "big nigger"
	fmt.Println(username)

	var username2 string = "black nigger"
	fmt.Println(username2)

	var username3 = "foolish nigger"
	var age = 22
	var isGirl = false
	fmt.Print(username3)
	fmt.Print(age)
	fmt.Println(isGirl)
}

/*全局变量*/
var niggerName = "lama"

/*批量声明*/
func DeclareInBatches() {
	var (
		username string
		age      int
		isGirl   bool
	)
	username = "Mary"
	age = 20
	isGirl = true

	fmt.Println(username, age, isGirl)

	var (
		a = "black"
		b = "nigger"
		c = "slave"
	)

	fmt.Println(a, b, c)
}

/*短变量声明*/
func localDeclare() {
	username := "明月光"
	var username2 = "为何又照地堂"
	fmt.Println(username, username2)

	//短变量批量声明
	name, age, isGirl := "Lucy", 22, true
	fmt.Println(name, age, isGirl)
}

/*
匿名变量
调用函数时，短横线代替不需要的返回值
*/
func getUserInfo() (string, int) {
	return "西八老马", 20
}

func main() {
	var username, age = getUserInfo()
	var username0, _ = getUserInfo()
	var _, age0 = getUserInfo()
	fmt.Println(username, age)
	fmt.Println(username0, age0)
}
