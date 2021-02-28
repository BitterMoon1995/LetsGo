package main

import (
	"fmt"
	"reflect"
)

func reflectBasic(i interface{}) {
	/*
		返回的是reflect.Type类型，不是真正的数据类型
	*/
	reflectType := reflect.TypeOf(i)
	fmt.Println(reflectType)
	/*
		基本类型kind等于type
	*/
	fmt.Println("kind：", reflectType.Kind(), "type：", reflectType)

	/*
		同样返回的是reflect.Value，因为不是int，所以不能进行整型的运算
	*/
	reflectValue := reflect.ValueOf(i)
	fmt.Printf("%v %T\n", reflectValue, reflectValue)
	var x int64 = 88
	//y:=x+rValue

	/*
		reflect.Value可以通过Int()转成Int64
	*/
	realValue := reflectValue.Int()
	realValue += x
	/*
		panic: reflect: call of reflect.Value.Bool on int Value
	*/
	//reflectValue.Bool()

	/*
		reflect.Value也可以通过Interface()转成Interface，
		再可以通过类型断言转成int
	*/
	toInterface := reflectValue.Interface()
	toInt := toInterface.(int)
	fmt.Println(toInt)
}

type girl struct {
	name     string
	age      int
	isBeauty bool
}

func reflectStruct(i interface{}) {
	reflectType := reflect.TypeOf(i)
	/*
		结构体，kind就是struct，type是包名.结构体名
	*/
	fmt.Println("kind：", reflectType.Kind(), "  type：", reflectType)
	reflectValue := reflect.ValueOf(i)
	fmt.Printf("%v %T\n", reflectValue, reflectValue)

	/*
		转回Girl类型：先转成interface再断言
	*/
	vInterface := reflectValue.Interface()
	/*
		vInterface.name undefined (type interface {} is interface with no methods)
		反射的本质是在运行时确认类型，在编译阶段vInterface还是interface {}
	*/
	//fmt.Println(vInterface.name)
	vGirl, ok := vInterface.(girl)
	if ok {
		fmt.Printf("%v %T\n", vGirl, vGirl)
		fmt.Println(vGirl.name)
	}
}

/*
简单研究下.Elem()
*/
func theElemFunc(inter interface{}) {
	elem := reflect.ValueOf(inter).Elem()
	fmt.Printf("%p\n", &elem)
}

/*
通过反射修改变量的值，一定要传入变量的指针，否则报using unaddressable value。
再调用Elem()获取到pointer points to，再进行SetXXX
*/
func editBasicByReflect(inter interface{}) {
	value := reflect.ValueOf(inter)
	fmt.Println("kind：", value.Kind(), "  type：", reflect.TypeOf(value))
	/*
		panic: reflect: reflect.Value.SetInt using unaddressable value
	*/
	//value.SetInt(8848)
	value.Elem().SetInt(8848)
}

func main() {
	num := 998
	p := &num
	fmt.Println(p)
	theElemFunc(p)
}
