package main

import (
	"fmt"
	"reflect"
)

type fairy struct {
	Name      string
	Age       int
	FaceScore float32
	isSlut    bool
}

func (f fairy) sleep() {
	fmt.Println("仙女睡觉")
}

func (f fairy) wear() {
	fmt.Println("仙女穿裙")
}

func editStructValue(any interface{}) *fairy {
	value := reflect.ValueOf(any).Elem()
	fmt.Println("kind：", value.Kind(), "  type：", value)
	fieldQt := value.NumField()
	for i := 0; i < fieldQt; i++ {
		fmt.Printf("字段%d 值%v\n", i, value.Field(i))
	}
	/*
		set字段时字段必须是公开（大写）的，否则
		panic: reflect: reflect.Value.SetString using value obtained using unexported field
	*/
	name := value.FieldByName("Name")
	name.SetString("李萱诗")
	age := value.FieldByName("Age")
	age.SetInt(48)

	//简单断个言操作个指针
	any2 := any.(*fairy)
	return any2
}

func main() {
	anJian := fairy{
		Name:      "安简",
		Age:       22,
		FaceScore: 8.5,
		isSlut:    true,
	}
	LXS := editStructValue(&anJian)
	fmt.Println(*LXS)
}
