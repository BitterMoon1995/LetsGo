package main

import (
	"fmt"
	"reflect"
)

type Woman interface {
	WearDress()
}

type NativeWoman struct {
	Name string
}

func (n NativeWoman) WearDress() {
	fmt.Println("死妈母狗发浪")
}

func (n NativeWoman) HandInMoney() {
	fmt.Println("交房,交车,交工资")
}

type Trans struct {
	Name string
}

func (t Trans) WearDress() {
	fmt.Println("仙女穿裙")
}

/**
接口的多种不同的实现方式即为多态。

接口体现多态特征，主要表现在【多态参数】 和 【多态数组】
*/

func showGirls(woman Woman) {
	woman.WearDress()
}

func polymorphicSlice() {
	women := make([]Woman, 0)
	women = append(women, Trans{Name: "薇儿"})
	women = append(women, Trans{Name: "蕊蕊"})
	women = append(women, NativeWoman{Name: "死妈母狗"})

	fmt.Println(women)
}

/*
类型断言的写法是 a.(T) ，会尝试把a转换为T类型，如果不成功会报panic
*/
func typeAssert() {
	//e1.
	var a interface{}
	var girl Trans = Trans{Name: "美妖雨薇儿"}
	a = girl                       //可以
	fmt.Println(reflect.TypeOf(a)) //???

	c := Trans{Name: "Lucy"}
	//c = a//不可以，盘点十大理由：

	c = a.(Trans) //可以
	fmt.Println(c)

	//e2.
	var x interface{}
	var y = 1
	x = y
	var z int = x.(int)
	fmt.Println(reflect.TypeOf(z))
}

/*
给类型断言加上检测机制，接收并处理断言的结果，避免因为类型转换失败导致整个程序终止
*/
func assertWithDetection() {
	var a int = 10
	var b interface{}
	var d float64 = 2.88

	b = a
	c, ok := b.(int)
	fmt.Println(c, ok)

	/*
		程序不会因为转换失败而终止，但是如果是将转换的结果赋给已有值的变量，
		该变量的值会丢失
	*/
	d, ok = b.(float64)
	fmt.Println(d, ok)

	if e, ok := b.(float64); !ok {
		fmt.Println("转换失败！")
	} else {
		fmt.Println(e)
	}
}

/*
类型断言在多态数组中的应用
*/
func typeAssertInPolymorphicArray() {
	women := make([]Woman, 0)
	women = append(women, Trans{Name: "薇儿"})
	women = append(women, Trans{Name: "蕊蕊"})
	women = append(women, NativeWoman{Name: "三台死妈母狗"})
	women = append(women, NativeWoman{Name: "都江堰丑母驴"})

	for _, woman := range women {
		bitch, ok := woman.(NativeWoman)
		if ok {
			bitch.HandInMoney()
		} else {
			woman.WearDress()
		}
	}
}

/*
类型断言在类型判断中的应用
*/
func typeJudge(items ...interface{}) {
	for _, x := range items {
		switch x.(type) {
		case int:
			fmt.Println(x, " 为int型")
		case float64:
			fmt.Println(x, " 为float64型")
		case string:
			fmt.Println(x, " 为string型")
		case nil:
			fmt.Println(x, " 为nil型")
		default:
			fmt.Println(x, " 不知道是个什么东西")
		}
	}
}

func main() {
	var u interface{}
	typeJudge(1, 2.558, "nigger", u)
}
