package main

import (
	"fmt"
)

/**
golang用结构体来作为“类”这一概念的承载，
用结构体的字段和方法来抽象现实事物的属性和行为

结构体的本质是自定义的数据类型，代表一类事物，
结构体变量（实例）是具体事物。
字段可以是基本类型、数组、结构体、引用类型，
如果是引用类型，在未初始化值的时候，值为nil，即还没有分配空间。
*/
type Cat struct {
	Name  string
	Breed string
	Age   int
}

/*
下两种声明方式得到的是结构体变量本身，
如果进行彼此间赋值，遵循值类型的原则，进行值拷贝
*/
func declareObj1() {
	var cat Cat
	/*结构体是值类型，声明即初始化并分配空间*/
	fmt.Println(cat)

	cat.Name = "小乖喵"
	cat.Age = 1
	cat.Breed = "布偶"

	fmt.Printf("值：%v\n", cat)

	cat2 := Cat{
		Name:  "Diona",
		Breed: "小坏喵",
		Age:   11,
	}
	fmt.Printf("类型：%T\n", cat2)

	/*
		进行的是值拷贝，cat cat2彼此独立，互不影响
	*/
	cat2 = cat
	cat2.Breed = "田园橘猫"
	fmt.Println(cat)
}

/*
下两种声明方法事实上返回的是结构体指针，
正确的属性赋值方式是(*cat1).Name = "..."等等
但是go做了简化，可以直接指针.属性，编译时会自动猴
*/
func declareObj2() {
	cat1 := new(Cat)
	cat1.Age = 2
	cat1.Breed = "田园三花"
	cat1.Name = "坏喵喵"

	fmt.Printf("类型：%T\n", cat1)

	cat2 := &Cat{}
	cat2.Age = 3
	cat2.Breed = "田园橘猫"
	cat2.Name = "大猫猪"
	fmt.Printf("值：%v\n", cat2)
}

func trapOfStructDeclare() {
	cat := Cat{
		Name:  "大橘Mango",
		Breed: "田园肥橘",
		Age:   3,
	}
	cat1 := new(Cat)
	cat1 = &cat

	cat2 := &Cat{}
	/*
		此时指针cat1 cat2指向同一个结构体，
		修改其中一个的值，势必会引起另一个的改变，因为是对同一个结构体的修改。
		当然指针cat1 cat2本身的地址是不一样的，但是值（所指向结构体）是一样的
	*/
	cat2 = cat1
	cat2.Name = "大坏喵喵"

	fmt.Println(cat1)
	fmt.Printf("结构体地址：%p\n", &cat)
	fmt.Printf("二猫地址：%p %p\n", &cat1, &cat2)
	fmt.Printf("二猫值：%p %p", cat1, cat2)

}

func main() {
	trapOfStructDeclare()
}
