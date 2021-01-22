package main

import (
	"fmt"
	"reflect"
)

func addrOfVerify() {
	var i int = 6324
	/*泪目
	语句含义：
	1.point是一个指针变量
	2.point是指向int的指针,*int
	3.point本身的值是i的地址
	*/
	var point *int = &i
	fmt.Println(point)
	/* *获取指针变量指向的值 */
	fmt.Println(*point)
	fmt.Println("卷起来咯：", &point)
}

func editByPointer() {
	a := 10
	b := 20
	pa := &a
	*pa = 99
	pa = &b
	*pa = 88
	fmt.Println(a, b)
}

/*
值类型都有对应的指针
值类型包括：基本数据类型（int float bool string）、数组、结构体
引用类型包括：指针、slice切片、map、channel、interface等
*/
func arrayIsValueType() {
	arr := [...]int{}
	pointer := &arr
	fmt.Println(reflect.TypeOf(pointer))
}

type Student struct {
	name   string
	age    int
	isGirl bool
}

func structIsValueType() {
	var stu Student
	stu.name = "薇儿"
	stu.age = 14
	stu.isGirl = false
	var pointer *Student = &stu
	fmt.Println(*pointer)
}

/*
在传参的时候，值类型默认都是值传递，引用类型默认是引用传递
但两者传递的都是变量的【副本】，前者传递的是值的拷贝，后者是地址的拷贝
一般来说，地址的拷贝效率高，因为数据量总是很小，
而值拷贝考虑到有长数组、结构体、长字符串，拷贝效率可能很低
*/
func valueTransmit(a int) {
	fmt.Println(&a)
}

func main() {
	a := 10
	fmt.Println(&a)
	valueTransmit(a)
}
