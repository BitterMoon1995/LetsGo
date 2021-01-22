package main

import "fmt"

/*
常量定义
常量可以声明了不使用；常量的值不可改变；声明时必须初始化
*/
func constDeclare() {
	const niggerName = "Jack"
	//niggerName = "ma"
	const pi = 3.14
	//const pi2
}

/*
批量声明
如果批量声明时某一个常量省略了值，则该常量的值将和上一个常量相同
*/
func declareInBatches() {
	const (
		n1 = 100
		n2 = 200
		n3 = 300
	)
	const (
		k1 = 4396
		k2
		k3
		k4 = 2800
	)
	fmt.Println(k2, k3, k4)
}

/* 常量计数器iota
可以用 _ 跳过某个值;可以用来插队；*/
func useIota() {
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)

	const (
		n1 = iota
		_
		n2
		n3
		n4
	)
	fmt.Println(n1, n2, n3, n4)

	const (
		m1 = iota
		m2 = 100
		m3 = iota
		m4
	)
	fmt.Println(m1, m2, m3, m4)

	const (
		l1, l2 = iota + 1, iota + 2
		l3, l4
		l5, l6
	)
	fmt.Println(l1, l2, l3, l4, l5, l6)
}

func main() {
	useIota()
}
