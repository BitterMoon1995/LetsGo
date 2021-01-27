package main

import "fmt"

/*
1.并不是只有结构体可以实现接口，只要是自定义类型都可以
*/
type integer int

func (i integer) insert() {
	fmt.Println("64位整数插入USB接口")
}

func (i integer) pullout() {
	fmt.Println("64位整数拔出USB接口")
}

/*
2.一个自定义类型可以实现多个接口，
但是所有这些接口不能有同名方法，否则报method redeclare，
同名但签名不一样也不行，毕竟GO没有方法重载。
*/
type HTML interface {
	//insert()
	HDInsert()
}
type VGA interface {
	VGAInsert()
}

/*
3.一个接口可以继承一个或多个其他接口，如果要再实现这个接口，就要实现所有它继承的接口
*/
type GraphicsCard interface {
	HTML
	VGA
}

/*
4.空接口没有任何方法，所有类型都实现了空接口，也有一个固定写法为interface{}
*/
type T interface {
}

type RTX3080 struct {
}

func (R RTX3080) HDInsert() {
	panic("implement me")
}

func (R RTX3080) VGAInsert() {
	panic("implement me")
}

func main() {
	//var card GraphicsCard = RTX3080{}
	//var t T = RTX3080{}
	//var i interface{} = RTX3080{}
	//var int1 int = 2
	//var i2 interface{} = int1
}
