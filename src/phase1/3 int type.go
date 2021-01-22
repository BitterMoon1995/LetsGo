package main

import (
	"fmt"
	"unsafe"
)

func checkSize() {
	var (
		i1 int8
		i2 int16
		i3 int
		i4 int64
	)
	fmt.Println(unsafe.Sizeof(i1), unsafe.Sizeof(i2), unsafe.Sizeof(i3), unsafe.Sizeof(i4))
	//i1 = 127
	//i2 = 32767
	i3 = 9223372036854775807
}

func intConvert() {
	var a1 int32 = 12
	var a2 int64 = 21
	//fmt.Println(a1 + a2)
	fmt.Println(int64(a1) + a2)
	fmt.Println(a1 + int32(a2))

	var a3 int16 = 128
	//高位向低位转换，注意转换后范围是否足够
	fmt.Println(int8(a3))
	var a4 int16 = 127
	fmt.Println(int8(a4))
}

func main() {
	intConvert()
}
