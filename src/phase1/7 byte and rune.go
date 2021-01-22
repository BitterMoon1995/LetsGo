package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func byteBelongToInt() {
	var b = 'b'
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Printf("类型：%T\n", b)
	fmt.Println(reflect.ValueOf(b))
	fmt.Printf("原样输出：%c", b)
}

/*
字符串中一个字母占用一个字节（uint8 ASCII），一个汉字占用三个字节（rune UTF-8）
不能用Sizeof来判断，因为string的底层是用 来实现的
*/
func bytesOccupy() {
	str := "你好go"
	fmt.Println(len(str))
	fmt.Println(unsafe.Sizeof(str))
}

func byteInKanji() {
	b := '汉'
	fmt.Printf("值(十进制)：%v 原样输出：%c 类型：%T\n", b, b, b)
}

/*
字符分为byte类型(uint8)与rune类型(int32)
*/
func byteInString() {
	str := "nigga"

	for i := 0; i < len(str); i++ {
		fmt.Printf("值(十进制)：%v 原样输出：%c 类型：%T\n", str[i], str[i], str[i])
	}
}

func runeInString() {
	str := "你好golang"

	for _, r := range str {
		fmt.Printf("值(十进制)：%v 原样输出：%c 类型：%T\n", r, r, r)
	}
}

func editString1() {
	var str string = "big"
	bytes := []byte(str)
	bytes[0] = 'p'
	fmt.Println(string(bytes))
}

func main() {
	editString1()
}
