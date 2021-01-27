package main

import "fmt"

/**
方法就是一类只能被指定的数据类型调用的、带特殊的 接收者 参数的函数
方法调用时，会将调用方法的变量，也当做参数传给方法
*/

type Nigga struct {
	Name string
}

func (nigga Nigga) getNiggaName() {
	fmt.Println(nigga.Name)
}

/*
非自定义数据类型，也可以有自己的方法
*/
type Integer int

func (i Integer) value() {
	fmt.Println(i)
}
func (i *Integer) plus1() {
	*i = *i + 1
}

/*
如果一个数据类型实现了 String() 这个方法，那么对于该类型，
fmt.Println会默认调用该方法进行输出 （就是toString）
*/
type SchoolBabe struct {
	Name      string
	Score     int
	Height    float32
	Signature string
}

func (babe SchoolBabe) String() string {
	s := fmt.Sprintf("名字：%v 分（脸）：%v 身高：%v 个性签名：%v", babe.Name, babe.Score, babe.Height, babe.Signature)
	return s
}

/*
结构体变量调用方法时，编译器做了优化，即使方法申明的接收者（调用方）是该类型的指针类型，
仍然可以用值类型的变量调用，反之亦然
*/
func (babe *SchoolBabe) coquetry() {
	fmt.Println(babe.Name, " 在臭美")
}
func (babe SchoolBabe) lascivious() {
	fmt.Println(babe.Name, " 在发骚")
}

func main() {
	weier := new(SchoolBabe)
	weier.Name = "薇儿"
	weier.Height = 1.68
	weier.Score = 9
	weier.Signature = "长腿校花骚薇儿"
	weier.lascivious()

	ruirui := SchoolBabe{
		Name:      "蕊蕊",
		Score:     9,
		Height:    1.65,
		Signature: "骚美甜蜜嗲精",
	}
	ruirui.coquetry()
}
