package main

import (
	"fmt"
	"unsafe"
)

func checkSize2() {
	var f1 float32 = 3.12
	fmt.Println(unsafe.Sizeof(f1))
	var f2 float64 = 3.14
	fmt.Println(unsafe.Sizeof(f2))
	f3 := 3.1456854 //默认float64位
	fmt.Println(unsafe.Sizeof(f3))
	f4 := 3.14159265358979323846
	fmt.Println(f4)
}

func scientificNotation() {
	f := 3.14e2
	fmt.Println(f)
	f2 := 3.14e-2
	fmt.Println(f2)
}

func precisionLoss() {
	m1 := 8.2
	m2 := 4.8
	fmt.Println(m1 - m2)

	d := 1129.6
	fmt.Println(d * 100)

	//解决：第三方包decimal
}

func main() {
	precisionLoss()
}
